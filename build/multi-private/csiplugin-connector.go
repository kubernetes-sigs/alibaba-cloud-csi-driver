package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	// OSSSocketPath socket path
	OSSSocketPath = "/run/csi-tool/connector/connector.sock"
	// DiskSocketPath socket path
	DiskSocketPath = "/run/csi-tool/connector/diskconnector.sock"
	// ShellPath is the fsfreeze shell path
	ShellPath = "/etc/csi-tool/fsfreeze.sh"
	// GetPathDevice get the device of specific path
	GetPathDevice = "df --output=source %s"
)

func main() {
	log.Print("OSS Connector Daemon Is Starting...")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		EnsureSocketPath(OSSSocketPath)
		log.Printf("Socket path is ready: %s", OSSSocketPath)
		ln, err := net.Listen("unix", OSSSocketPath)
		if err != nil {
			log.Fatalf("Server Listen error: %s", err.Error())
		}
		log.Print("Daemon Started ...")
		defer ln.Close()

		go watchDogCheck()
		// Handler to process the command
		for {
			fd, err := ln.Accept()
			if err != nil {
				log.Printf("Server Accept error: %s", err.Error())
				continue
			}
			go echoServer(fd)
		}
	}()
	go func() {
		defer wg.Done()
		EnsureSocketPath(DiskSocketPath)
		log.Printf("Socket path is ready: %s", DiskSocketPath)
		ln, err := net.Listen("unix", DiskSocketPath)
		if err != nil {
			log.Fatalf("runDiskProxy: server Listen error: %v", err.Error())
		}
		log.Print("Disk proxy daemon started ....")
		defer ln.Close()
		go watchDogCheck()

		// Handler to process the command
		for {
			fd, err := ln.Accept()
			if err != nil {
				log.Printf("Disk Server Accept error: %s", err.Error())
				continue
			}
			go freezeFilesystemServer(fd)
		}
	}()
	wg.Wait()
}

func watchDogCheck() {
	// watchdog of UNIX Domain Socket
	var socketsPath []string
	if os.Getenv("WATCHDOG_SOCKETS_PATH") != "" {
		socketsPath = strings.Split(os.Getenv("WATCHDOG_SOCKETS_PATH"), ",")
	}
	socketNotAliveCount := make(map[string]int)
	if len(socketsPath) == 0 {
		return
	}
	for {
		deadSockets := 0
		for _, path := range socketsPath {
			if err := isUnixDomainSocketLive(path); err != nil {
				log.Printf("socket %s is not alive: %v", path, err)
				socketNotAliveCount[path]++
			} else {
				socketNotAliveCount[path] = 0
			}
			if socketNotAliveCount[path] >= 6 {
				deadSockets++
			}
		}
		if deadSockets >= len(socketsPath) {
			log.Printf("watchdog find too many dead sockets, csiplugin-connector will exit(0)")
			os.Exit(0)
		}
		time.Sleep(time.Second * 10)
	}
}

func freezeFilesystemServer(c net.Conn) {
	buf := make([]byte, 2048)
	nr, err := c.Read(buf)
	if err != nil {
		log.Printf("freezeFilesystemServer:: server read error: %v", err.Error())
		return
	}
	command := string(buf[0:nr])
	log.Printf("freezeFilesystemServer:: server receive freeze parms: %v", command)
	err = checkFilesystemConsistentCommand(command)
	if err != nil {
		out := "Fail:" + err.Error()
		log.Printf("freezeFilesystemServer:: check disk command error: %v", out)
		if _, err := c.Write([]byte(out)); err != nil {
			log.Printf("freezeFilesystemServer:: check disk command write error: %v", out)
		}
		return
	}
	log.Printf("freezeFilesystemServer:: command: %v", command)
	// run command
	if out, err := run(command); err != nil {
		reply := "Fail: " + command + ", error: " + err.Error()
		_, err = c.Write([]byte(reply))
		log.Print("diskServer Fail to run cmd:", reply)
	} else {
		out = "Success:" + out
		_, err = c.Write([]byte(out))
		log.Printf("Success: %s", out)
	}
}

func checkFilesystemConsistentCommand(paramStr string) error {
	params := strings.Split(paramStr, " ")
	for index, param := range params {
		if index == 0 {
			if !strings.EqualFold(param, "/etc/csi-tool/freezefs.sh") {
				return fmt.Errorf("checkFilesystemConsistentParams:: scripts name: %v invalid", param)
			}
		} else {
			if !strings.HasPrefix(param, "--path") && !strings.HasPrefix(param, "--timeout") && !strings.HasPrefix(param, "--type") && !strings.HasPrefix(param, "&") {
				return fmt.Errorf("checkFilesystemConsistentParams:: paramStr: %v invalid", param)
			}
		}
		if index == 2 {
			globalPath := strings.Split(param, "=")[1]
			log.Printf("checkFilesystemConsistentCommand:: globalPath: %v", globalPath)
			if !isIsolateDevice(globalPath) {
				return fmt.Errorf("checkFilesystemConsistentParams:: globalPath: %v isn't isolated device mount path", globalPath)
			}
		}
	}
	return nil
}

func isIsolateDevice(globalPath string) bool {
	globalPathCommand := fmt.Sprintf(GetPathDevice, globalPath)
	pathOut, err := run(globalPathCommand)
	if err != nil {
		reply := "Fail: " + globalPathCommand + ", error: " + err.Error()
		log.Print("Server Fail to run cmd:", reply)
		return false
	}
	globalPathDirCommad := fmt.Sprintf(GetPathDevice, filepath.Dir(globalPath))
	dirOut, err := run(globalPathDirCommad)
	if err != nil {
		reply := "Fail: " + globalPathDirCommad + ", error: " + err.Error()
		log.Print("Server Fail to run cmd:", reply)
		return false
	}

	log.Printf("isIsolateDevice:: pathOut： %s, dirOut: %s", pathOut, dirOut)
	return !strings.EqualFold(pathOut, dirOut)
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

	if err := checkOssfsCmd(cmd); err != nil {
		out := "Fail: " + err.Error()
		log.Printf("Check oss command error: %s", out)
		if _, err := c.Write([]byte(out)); err != nil {
			log.Printf("Check command write error: %s", err.Error())
		}
		return
	}
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

// systemd-run --scope -- /usr/local/bin/ossfs shenzhen
// /var/lib/kubelet/pods/070d1a40-16a4-11ea-842e-00163e062fe1/volumes/kubernetes.io~csi/oss-csi-pv/mount
// -ourl=oss-cn-shenzhen-internal.aliyuncs.com
// -o max_stat_cache_size=0 -o allow_other
func checkOssfsCmd(cmd string) error {
	jindofsPrefix := "systemd-run --scope -- /etc/jindofs-tool/jindofs-fuse "
	if strings.HasPrefix(cmd, jindofsPrefix) {
		if strings.Contains(cmd, ";") {
			return errors.New("Jindofs Options: command cannot contains ; " + cmd)
		}
		cmdParametes := strings.TrimPrefix(cmd, jindofsPrefix)
		cmdParametes = strings.TrimSpace(cmdParametes)
		cmdParametes = strings.Join(strings.Fields(cmdParametes), " ")
		parameteList := strings.Split(cmdParametes, " ")
		for _, value := range parameteList {
			if !strings.HasPrefix(value, "-o") && !strings.HasPrefix(value, "/") {
				return errors.New("Jindofs Options: must start with -o :" + cmd)
			}
		}
		return nil
	}
	ossCmdPrefixList := []string{"systemd-run --scope -- /usr/local/bin/ossfs", "systemd-run --scope -- ossfs", "ossfs"}
	ossCmdPrefix := ""
	for _, cmdPrefix := range ossCmdPrefixList {
		if strings.HasPrefix(cmd, cmdPrefix) {
			ossCmdPrefix = cmdPrefix
			break
		}
	}

	// check oss command options
	if ossCmdPrefix != "" {
		cmdParametes := strings.TrimPrefix(cmd, ossCmdPrefix)
		cmdParametes = strings.TrimSpace(cmdParametes)
		cmdParametes = strings.Join(strings.Fields(cmdParametes), " ")

		parameteList := strings.Split(cmdParametes, " ")
		if len(parameteList) < 3 {
			return errors.New("Oss Options: parameters less than 3: " + cmd)
		}
		if !IsFileExisting(parameteList[1]) {
			return errors.New("Oss Options: mountpoint not exist " + parameteList[1])
		}
		if !strings.HasPrefix(parameteList[2], "-ourl=") {
			return errors.New("Oss Options: url should start with -ourl: " + parameteList[2])
		}
		oFlag := false
		for index, value := range parameteList {
			if index < 3 {
				continue
			}
			if value == "-s" || value == "-d" || value == "--debug" {
				if oFlag {
					return errors.New("Oss Options: no expect string follow -o " + value)
				}
				continue
			}
			if strings.HasPrefix(value, "-o") && len(value) > 2 {
				if oFlag {
					return errors.New("Oss Options: no expect string follow -o " + value)
				}
				continue
			}
			if value == "-o" {
				if oFlag == true {
					return errors.New("Oss Options: inputs must -o string, 2 -o now ")
				}
				oFlag = true
				continue
			}
			if oFlag == true {
				oFlag = false
			} else {
				return errors.New("Oss Options: inputs must -o string, 2 string now ")
			}
		}
		return nil
	}
	return errors.New("Oss Options: options with error prefix: " + cmd)
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

func isUnixDomainSocketLive(socketPath string) error {
	fileInfo, err := os.Stat(socketPath)
	if err != nil || (fileInfo.Mode()&os.ModeSocket == 0) {
		return fmt.Errorf("socket file %s is invalid", socketPath)
	}
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

// EnsureSocketPath ...
func EnsureSocketPath(socketPath string) {
	if IsFileExisting(socketPath) {
		os.Remove(socketPath)
	} else {
		pathDir := filepath.Dir(socketPath)
		if !IsFileExisting(pathDir) {
			os.MkdirAll(pathDir, os.ModePerm)
		}
	}
}
