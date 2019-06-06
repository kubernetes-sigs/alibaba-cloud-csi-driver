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
	MB_SIZE      = 1024 * 1024
	LOG_FILENAME = "/var/log/alicloud/csi_connector.log"
	PID_FILENAME = "/var/log/alicloud/connector.pid"
	WORK_PATH    = "./"
	SOCKET_PATH  = "/etc/csi-tool/connector.sock"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: PID_FILENAME,
		PidFilePerm: 0644,
		LogFileName: LOG_FILENAME,
		LogFilePerm: 0640,
		WorkDir:     WORK_PATH,
		Umask:       027,
		Args:        []string{"alibabacloud.csiplugin.connector"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run connector: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	log.Print("Daemon is starting...")

	runOssProxy()
}

func runOssProxy() {

	if IsFileExisting(SOCKET_PATH) {
		os.Remove(SOCKET_PATH)
	} else {
		pathDir := filepath.Dir(SOCKET_PATH)
		if ! IsFileExisting(pathDir) {
			os.MkdirAll(pathDir, os.ModePerm)
		}
	}
	log.Print("Socket path is ready: ", SOCKET_PATH)
	ln, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		log.Print("Server Listen error: ", err.Error())
	}
	log.Print("Daemon Started ...")
	defer ln.Close()

	// Handler to process the command
	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Print("Server Accept error: ", err.Error())
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
	log.Print("Server Receive cmd:", cmd)

	// run command
	if out, err := run(cmd); err != nil {
		reply := "Fail:" + cmd + ", error: " + err.Error()
		_, err = c.Write([]byte(reply))
		log.Print("Server Fail to run cmd:", reply)
	} else {
		out = "Success:" + out
		_, err = c.Write([]byte(out))
		log.Print("Success:", out)
	}
}

func run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: " + cmd + ", with out: " + string(out) + ", with error: " + err.Error())
	}
	return string(out), nil
}

// check file exist in volume driver;
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
