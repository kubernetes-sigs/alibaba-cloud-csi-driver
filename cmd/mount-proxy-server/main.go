package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/alinas"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/ossfs"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/ossfs2"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

var (
	socketPath     string
	drivers        []string
	handleTimeout  time.Duration
	enableNftables bool
)

func main() {
	flag.StringVar(&socketPath, "socket", "/var/run/csi/mounter.sock", "socket path")
	flag.StringSliceVar(&drivers, "driver", nil, "drivers to enable (e.g. 'ossfs,alinas')")
	flag.DurationVar(&handleTimeout, "timeout", time.Second*30, "timeout for connection")
	flag.BoolVar(&enableNftables, "enable-nftables", false, "enable nftables rules to restrict mount proxy port access (default: false)")
	utils.AddKlogFlags(flag.CommandLine)
	utils.AddGoFlags(flag.CommandLine)
	flag.Parse()

	_ = os.Remove(socketPath)
	server.Init(drivers)

	listener, err := listen(socketPath)
	if err != nil {
		klog.ErrorS(err, "Failed to listen", "socket", socketPath)
		os.Exit(1)
	}
	defer os.Remove(socketPath)

	klog.InfoS("Listening on socket", "socket", socketPath)

	if enableNftables {
		klog.InfoS("Init nftables", "enabled", true)
		if err := setupNftables(); err != nil {
			klog.ErrorS(err, "Failed to setup nftables")
			os.Exit(1)
		}
	} else {
		klog.InfoS("Skipping nftables setup", "enabled", false)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		klog.InfoS("exiting", "signal", sig)
		// close socket
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			klog.ErrorS(err, "Failed to accept connection")
			continue
		}
		go handle(conn)
	}
	server.Terminate(drivers)
}

func handle(conn net.Conn) {
	defer conn.Close()
	err := server.Handle(conn, handleTimeout)
	if err != nil {
		klog.ErrorS(err, "Failed to handle")
	}
}

func listen(socketPath string) (net.Listener, error) {
	if len(socketPath) < 100 {
		return net.Listen("unix", socketPath)
	}

	// Need to change the current working directory to the temp volume base path,
	// because the socket absolute path is longer than 108(104 on OSX) characters,
	// which will cause "bind: invalid argument" errors.
	exPwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get the current directory: %w", err)
	}

	socketDir := filepath.Dir(socketPath)

	if err = os.Chdir(socketDir); err != nil {
		return nil, fmt.Errorf("failed to change directory to %q: %w", socketDir, err)
	}

	defer func() {
		if err := os.Chdir(exPwd); err != nil {
			klog.ErrorS(err, "Failed to change directory back")
		}
	}()

	return net.Listen("unix", filepath.Base(socketPath))
}

const (
	nftablesTableName  = "csi_mount_proxy"
	nftablesChainName  = "early_filter"
	nftablesTargetPort = 12049
)

// setupNftables configures nftables to block local access to the mount proxy port
// from processes outside the target cgroup. This prevents unauthorized access.
func setupNftables() error {
	commands := []string{
		fmt.Sprintf("add table inet %s", nftablesTableName),
		fmt.Sprintf("add chain inet %s %s { type filter hook output priority -150 ; policy accept ; }", nftablesTableName, nftablesChainName),
		fmt.Sprintf("add rule inet %s %s oif lo tcp dport %d meta cgroup != 0 drop", nftablesTableName, nftablesChainName, nftablesTargetPort),
	}

	for _, cmd := range commands {
		c := exec.Command("nft", strings.Split(cmd, " ")...)
		if output, err := c.CombinedOutput(); err != nil {
			// Ignore table/chain already exists errors
			if !strings.Contains(string(output), "File exists") {
				return fmt.Errorf("nft error: %w, output: %s", err, string(output))
			}
		}
	}
	return nil
}
