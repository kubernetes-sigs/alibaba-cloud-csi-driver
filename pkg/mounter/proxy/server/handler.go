package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

const (
	defaultHandleTimeout = time.Second * 10
)

func Handle(conn net.Conn, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	unixConn, ok := conn.(*net.UnixConn)
	if !ok {
		return errors.New("failed to cast conn to *net.UnixConn")
	}
	connf, err := unixConn.File()
	if err != nil {
		return err
	}
	socket := int(connf.Fd())
	defer connf.Close()

	err = utils.WaitFdReadable(socket, timeout)
	if err != nil {
		return err
	}

	klog.V(4).InfoS("Start to recvmsg")
	p := make([]byte, proxy.MaxMsgSize)
	n, _, _, _, err := unix.Recvmsg(socket, p, nil, 0)
	if err != nil {
		return fmt.Errorf("recvmsg: %w", err)
	}
	klog.V(4).InfoS("Succeeded to recvmsg", "n", n)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var resp proxy.Response
	req, err := parseRawRequest(p[:n])
	if err != nil {
		resp = proxy.Response{
			Error: err.Error(),
		}
	} else {
		resp = handle(ctx, req)
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("encode response: %w", err)
	}

	klog.V(4).InfoS("Start to sendmsg")
	err = unix.Sendmsg(socket, append(data, proxy.MessageEnd), nil, nil, 0)
	if err != nil {
		return fmt.Errorf("sendmsg: %w", err)
	}
	klog.InfoS("Succeeded to sendmsg", "msg", string(data))
	return nil
}

type rawRequest struct {
	Header proxy.Header    `json:"header"`
	Body   json.RawMessage `json:"body,omitempty"`
}

func parseRawRequest(data []byte) (*rawRequest, error) {
	end := bytes.IndexByte(data, proxy.MessageEnd)
	if end == -1 {
		return nil, errors.New("invalid message")
	}
	var req rawRequest
	return &req, json.Unmarshal(data[:end], &req)
}

func handle(ctx context.Context, req *rawRequest) proxy.Response {
	switch req.Header.Method {
	case proxy.Mount:
		var mountReq proxy.MountRequest
		err := json.Unmarshal(req.Body, &mountReq)
		if err != nil {
			return proxy.Response{
				Error: err.Error(),
			}
		}
		err = handleMountRequest(ctx, &mountReq)
		if err != nil {
			return proxy.Response{
				Error: err.Error(),
			}
		}
	default:
		return proxy.Response{
			Error: "invalid method",
		}
	}
	return proxy.Response{}
}

func Init(driverNames []string) {
	for _, name := range sets.New(driverNames...).UnsortedList() {
		if driver, ok := nameToDriver[name]; ok {
			for _, fstype := range driver.Fstypes() {
				fstypeToDriver[fstype] = driver
			}
			driver.Init()
		}
	}
}

func Terminate(driverNames []string) {
	var wg sync.WaitGroup
	for _, name := range sets.New(driverNames...).UnsortedList() {
		if driver, ok := nameToDriver[name]; ok {
			wg.Add(1)
			go func(m Driver) {
				defer wg.Done()
				m.Terminate()
			}(driver)
		}
	}
	wg.Wait()
}
