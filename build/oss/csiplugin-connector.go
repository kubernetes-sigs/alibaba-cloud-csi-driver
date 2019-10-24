package main

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	// MBSIZE metrics
	MBSIZE = 1024 * 1024
	// LogFilename name of log file
	LogFilename = "/var/log/alicloud/csi_connector.log"
	// PIDFilename name of pid file
	PIDFilename = "/var/log/alicloud/connector.pid"
	// WorkPath workspace
	WorkPath = "./"
	// SocketPath socket path
	SocketPath = "/etc/csi-tool/connector.sock"
)

func main() {
	// log file is: /var/log/alicloud/csi_connector.log
	cntxt := &daemon.Context{
		PidFileName: PIDFilename,
		PidFilePerm: 0644,
		LogFileName: LogFilename,
		LogFilePerm: 0640,
		WorkDir:     WorkPath,
		Umask:       027,
		Args:        []string{"alibabacloud.csiplugin.connector"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalf("Unable to run connector: %s", err.Error())
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	log.Print("OSS Connector Daemon Is Starting...")

	runOssProxy()
}

func runOssProxy() {

	if IsFileExisting(SocketPath) {
		os.Remove(SocketPath)
	} else {
		pathDir := filepath.Dir(SocketPath)
		if !IsFileExisting(pathDir) {
			os.MkdirAll(pathDir, os.ModePerm)
		}
	}
	log.Printf("Socket path is ready: %s", SocketPath)
	ln, err := net.Listen("unix", SocketPath)
	if err != nil {
		log.Fatalf("Server Listen error: %s", err.Error())
	}
	log.Print("Daemon Started ...")
	defer ln.Close()

	// Handler to process the command
	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Printf("Server Accept error: %s", err.Error())
			continue
		}
		go echoServer(fd)
	}
	log.Print("Daemon Stoped ...")
}

func echoServer(c net.Conn) {
	buf := make([]byte, 2048)
	nr, err := c.Read(buf)
	if err != nil {
		log.Print("Server Read error: ", err.Error())
		return
	}

	cmd := string(buf[0:nr])
	log.Printf("Server Receive OSS command: %s", cmd)

	// run command
	if out, err := run(cmd); err != nil {
		reply := "Fail: " + cmd + ", error: " + err.Error()
		_, err = c.Write([]byte(reply))
		log.Print("Server Fail to run cmd:", reply)
	} else {
		out = "Success:" + out
		_, err = c.Write([]byte(out))
		log.Printf("Success: %s", out)
	}
}

func run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: " + cmd + ", with out: " + string(out) + ", with error: " + err.Error())
	}
	return string(out), nil
}

// IsFileExisting checks file exist in volume driver or not
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
