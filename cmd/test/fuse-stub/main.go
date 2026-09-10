//go:build linux

// fuse-stub is a minimal FUSE daemon for integration testing. It implements
// only the opcodes needed to make a mount point appear alive (INIT, GETATTR,
// STATFS) and returns ENOSYS for everything else.
//
// Usage:
//
//	fuse-stub <mountpoint>                  # opens /dev/fuse, calls mount(2)
//	fuse-stub --fd=3 <mountpoint>           # uses an inherited fd (fd-passing mode)
//	fuse-stub --fd=3 --state-dir=<dir> <mountpoint>  # recovery-aware mode
//
// In recovery-aware mode the stub writes a "fuse_fd" marker file on first run.
// On subsequent runs (marker exists) it skips mount(2) and reattaches to the
// existing FUSE connection via the inherited fd — matching ossfs2's behaviour.
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// FUSE kernel protocol constants (from <linux/fuse.h>).
const (
	fuseKernelVersion      = 7
	fuseKernelMinorVersion = 39 // match a recent kernel

	fuseInitOp    = 26
	fuseGetattrOp = 3
	fuseStatfsOp  = 17
	fuseDestroyOp = 38

	fuseMinReadBuf = 1<<20 + 4096 // kernel default max_write (1MB) + FUSE header
)

// On-wire header for every FUSE request/response.
type fuseInHeader struct {
	Len     uint32
	Opcode  uint32
	Unique  uint64
	NodeID  uint64
	UID     uint32
	GID     uint32
	PID     uint32
	Padding uint32
}

type fuseOutHeader struct {
	Len    uint32
	Error  int32
	Unique uint64
}

type fuseInitIn struct {
	Major        uint32
	Minor        uint32
	MaxReadahead uint32
	Flags        uint32
}

type fuseInitOut struct {
	Major                uint32
	Minor                uint32
	MaxReadahead         uint32
	Flags                uint32
	MaxBackground        uint16
	CongestionThreshold  uint16
	MaxWrite             uint32
	TimeGran             uint32
	MaxPages             uint16
	MapAlignment         uint16
	Flags2               uint32
	MaxStackDepth        uint32
	Unused               [6]uint32
}

type fuseAttrOut struct {
	AttrValid     uint64
	AttrValidNsec uint32
	Dummy         uint32
	Attr          fuseAttr
}

type fuseAttr struct {
	Ino       uint64
	Size      uint64
	Blocks    uint64
	Atime     uint64
	Mtime     uint64
	Ctime     uint64
	AtimeNsec uint32
	MtimeNsec uint32
	CtimeNsec uint32
	Mode      uint32
	Nlink     uint32
	UID       uint32
	GID       uint32
	Rdev      uint32
	Blksize   uint32
	Flags     uint32
}

type fuseStatfsOut struct {
	St fuseKstatfs
}

type fuseKstatfs struct {
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Bsize   uint32
	Namelen uint32
	Frsize  uint32
	Padding uint32
	Spare   [6]uint32
}

func main() {
	fdFlag := flag.Int("fd", 0, "inherited FUSE fd (0 = open /dev/fuse ourselves)")
	stateDir := flag.String("state-dir", "", "runtime state directory for recovery detection")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: fuse-stub [--fd=N] [--state-dir=DIR] <mountpoint>")
		os.Exit(1)
	}
	mountpoint := flag.Arg(0)

	fuseFd := *fdFlag
	isRecovery := false

	if *stateDir != "" {
		marker := filepath.Join(*stateDir, "fuse_fd")
		if _, err := os.Stat(marker); err == nil {
			isRecovery = true
		} else {
			if err := os.WriteFile(marker, []byte(fmt.Sprintf("%d", fuseFd)), 0o644); err != nil {
				die("write marker: %v", err)
			}
		}
	}

	if fuseFd <= 0 {
		fd, err := unix.Open("/dev/fuse", unix.O_RDWR, 0)
		if err != nil {
			die("open /dev/fuse: %v", err)
		}
		fuseFd = fd
	}

	if !isRecovery && *fdFlag <= 0 {
		// Only call mount(2) when we opened /dev/fuse ourselves.
		// In fd-passing mode the caller already mounted; we just serve.
		opts := fmt.Sprintf("fd=%d,rootmode=40000,user_id=%d,group_id=%d,allow_other",
			fuseFd, os.Getuid(), os.Getgid())
		if err := unix.Mount("fuse-stub", mountpoint, "fuse.fuse-stub", unix.MS_NOSUID|unix.MS_NODEV, opts); err != nil {
			die("mount: %v", err)
		}
	}

	// Unblock after mount so the parent can detect readiness via stat().
	fmt.Fprintf(os.Stderr, "fuse-stub: serving on %s (fd=%d, recovery=%v)\n", mountpoint, fuseFd, isRecovery)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Fprintln(os.Stderr, "fuse-stub: signal received, exiting")
		os.Exit(0)
	}()

	serve(fuseFd)
}

func serve(fd int) {
	buf := make([]byte, fuseMinReadBuf)
	for {
		n, err := unix.Read(fd, buf)
		if err != nil {
			if err == syscall.EINTR {
				continue
			}
			if err == syscall.ENODEV {
				// Connection was aborted (unmounted). Clean exit.
				return
			}
			die("read: %v", err)
		}
		if n < int(unsafe.Sizeof(fuseInHeader{})) {
			continue
		}

		var hdr fuseInHeader
		hdr.Len = binary.LittleEndian.Uint32(buf[0:4])
		hdr.Opcode = binary.LittleEndian.Uint32(buf[4:8])
		hdr.Unique = binary.LittleEndian.Uint64(buf[8:16])
		hdr.NodeID = binary.LittleEndian.Uint64(buf[16:24])

		switch hdr.Opcode {
		case fuseInitOp:
			handleInit(fd, &hdr, buf[unsafe.Sizeof(fuseInHeader{}):n])
		case fuseGetattrOp:
			handleGetattr(fd, &hdr)
		case fuseStatfsOp:
			handleStatfs(fd, &hdr)
		case fuseDestroyOp:
			return
		default:
			replyError(fd, hdr.Unique, syscall.ENOSYS)
		}
	}
}

func handleInit(fd int, hdr *fuseInHeader, body []byte) {
	var in fuseInitIn
	if len(body) >= int(unsafe.Sizeof(in)) {
		in.Major = binary.LittleEndian.Uint32(body[0:4])
		in.Minor = binary.LittleEndian.Uint32(body[4:8])
		in.MaxReadahead = binary.LittleEndian.Uint32(body[8:12])
		in.Flags = binary.LittleEndian.Uint32(body[12:16])
	}
	_ = in // logged for debugging if needed

	var out fuseInitOut
	out.Major = fuseKernelVersion
	out.Minor = fuseKernelMinorVersion
	out.MaxReadahead = 131072
	out.MaxWrite = 131072
	out.MaxBackground = 12
	out.CongestionThreshold = 9

	reply(fd, hdr.Unique, &out)
}

func handleGetattr(fd int, hdr *fuseInHeader) {
	var out fuseAttrOut
	out.AttrValid = 1
	out.Attr.Ino = hdr.NodeID
	if hdr.NodeID == 1 {
		out.Attr.Mode = unix.S_IFDIR | 0o755
		out.Attr.Nlink = 2
	} else {
		out.Attr.Mode = unix.S_IFREG | 0o644
		out.Attr.Nlink = 1
		out.Attr.Size = 0
	}
	out.Attr.UID = uint32(os.Getuid())
	out.Attr.GID = uint32(os.Getgid())
	reply(fd, hdr.Unique, &out)
}

func handleStatfs(fd int, hdr *fuseInHeader) {
	var out fuseStatfsOut
	out.St.Blocks = 1000000
	out.St.Bfree = 500000
	out.St.Bavail = 500000
	out.St.Files = 100000
	out.St.Ffree = 50000
	out.St.Bsize = 4096
	out.St.Namelen = 255
	out.St.Frsize = 4096
	reply(fd, hdr.Unique, &out)
}

func reply(fd int, unique uint64, payload any) {
	var payloadBytes []byte
	switch v := payload.(type) {
	case *fuseInitOut:
		payloadBytes = (*[unsafe.Sizeof(fuseInitOut{})]byte)(unsafe.Pointer(v))[:]
	case *fuseAttrOut:
		payloadBytes = (*[unsafe.Sizeof(fuseAttrOut{})]byte)(unsafe.Pointer(v))[:]
	case *fuseStatfsOut:
		payloadBytes = (*[unsafe.Sizeof(fuseStatfsOut{})]byte)(unsafe.Pointer(v))[:]
	default:
		panic("unknown payload type")
	}

	outSize := int(unsafe.Sizeof(fuseOutHeader{})) + len(payloadBytes)
	var outHdr fuseOutHeader
	outHdr.Len = uint32(outSize)
	outHdr.Error = 0
	outHdr.Unique = unique

	hdrBytes := (*[unsafe.Sizeof(fuseOutHeader{})]byte)(unsafe.Pointer(&outHdr))[:]
	msg := make([]byte, 0, outSize)
	msg = append(msg, hdrBytes...)
	msg = append(msg, payloadBytes...)

	for len(msg) > 0 {
		n, err := unix.Write(fd, msg)
		if err != nil {
			if err == syscall.EINTR {
				continue
			}
			return // ENODEV or other terminal error
		}
		msg = msg[n:]
	}
}

func replyError(fd int, unique uint64, errno syscall.Errno) {
	var outHdr fuseOutHeader
	outHdr.Len = uint32(unsafe.Sizeof(fuseOutHeader{}))
	outHdr.Error = -int32(errno)
	outHdr.Unique = unique

	hdrBytes := (*[unsafe.Sizeof(fuseOutHeader{})]byte)(unsafe.Pointer(&outHdr))[:]
	_, _ = unix.Write(fd, hdrBytes)
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "fuse-stub: "+format+"\n", args...)
	os.Exit(1)
}
