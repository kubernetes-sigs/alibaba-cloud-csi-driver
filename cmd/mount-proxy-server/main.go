package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/alinas"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/ossfs"
	"k8s.io/klog/v2"
)

var (
	socketPath    string
	handleTimeout time.Duration
)

func main() {
	flag.StringVar(&socketPath, "socket", "/var/run/csi/mounter.sock", "socket path")
	flag.DurationVar(&handleTimeout, "timeout", time.Second*30, "timeout for connection")
	klog.InitFlags(nil)
	flag.Parse()

	_ = os.Remove(socketPath)
	server.Init()

	listener, err := listen(socketPath)
	if err != nil {
		klog.ErrorS(err, "Failed to listen", "socket", socketPath)
		os.Exit(1)
	}
	defer os.Remove(socketPath)

	klog.InfoS("Listening on socket", "socket", socketPath)

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
	server.Terminate()
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
