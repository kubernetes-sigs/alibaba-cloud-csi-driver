package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

const (
	// QueryServerSocket tag, used for queryserver socket
	QueryServerSocket = "/var/run/node-extender-server/volume-query-server.sock"
)

// QueryRequest struct
// Identity: for volumeInfo Request
// PodName/PodNameSpace: PodRunTime Request
type QueryRequest struct {
	Identity     string `json:"identity"`
	PodName      string `json:"podName"`
	PodNameSpace string `json:"podNameSpace"`
}

// QueryServer Kata Server
type QueryServer struct {
	client kubernetes.Interface
}

// NewQueryServer new server
func NewQueryServer() *QueryServer {
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	return &QueryServer{
		client: kubeClient,
	}
}

// RunQueryServer Routers
func (ks *QueryServer) RunQueryServer() {
	socketAddr := &net.UnixAddr{Name: QueryServerSocket, Net: "unix"}
	os.Remove(socketAddr.Name)
	lis, err := net.ListenUnix("unix", socketAddr)
	if err != nil {
		klog.Errorf("Listen Unix error: %s", err.Error())
		return
	}

	// set router
	klog.Infof("Started Query Server with unix socket: %s", QueryServerSocket)
	http.HandleFunc("/api/v1/volumeinfo", ks.volumeInfoHandler)
	//	http.HandleFunc("/api/v1/podruntime", ks.podRunTimeHandler)
	http.HandleFunc("/api/v1/ping", ks.pingHandler)

	// Server Listen
	svr := &http.Server{Handler: http.DefaultServeMux}
	err = svr.Serve(lis)
	if err != nil {
		klog.Errorf("Query Server Starting error: %s", err.Error())
	}
	klog.Infof("Query Server Ending ....")
}

// volumeInfoHandler reply with volume options.
func (ks *QueryServer) volumeInfoHandler(w http.ResponseWriter, r *http.Request) {
	reqInfo := QueryRequest{}
	content, err := io.ReadAll(r.Body)
	if err != nil {
		klog.Errorf("Request volumeInfo: Receive request read body error: %s", err.Error())
		fmt.Fprintf(w, "null")
		return
	}
	if err := json.Unmarshal(content, &reqInfo); err != nil {
		klog.Errorf("Request volumeInfo: Unmarshal request body(%s) error: %s", string(content), err.Error())
		fmt.Fprintf(w, "null")
		return
	}
	klog.Infof("Request volumeInfo: Receive Request with identity: %s", reqInfo.Identity)
	if reqInfo.Identity == "" {
		fmt.Fprintf(w, "null")
		return
	}

	// Response with file content
	fileName := filepath.Join(reqInfo.Identity, utils.CsiPluginRunTimeFlagFile)
	if utils.IsFileExisting(fileName) {
		// Unmarshal file content to map
		fileContent := utils.GetFileContent(fileName)
		fileContent = strings.ToLower(fileContent)
		volInfoMapFrom := map[string]string{}
		if err := json.Unmarshal([]byte(fileContent), &volInfoMapFrom); err != nil {
			klog.Errorf("Request volumeInfo: Unmarshal fileContent (%s) error: %s", fileContent, err.Error())
			fmt.Fprintf(w, "null")
			return
		}
		volumeType := ""
		if value, ok := volInfoMapFrom["volumetype"]; ok {
			volumeType = value
		}
		// copy parts of items to new map
		volInfoMapResponse := map[string]string{}
		// for disk volume type
		if volumeType == "block" {
			if value, ok := volInfoMapFrom["device"]; ok {
				volInfoMapResponse["path"] = value
			}
			if value, ok := volInfoMapFrom["identity"]; ok {
				volInfoMapResponse["identity"] = value
			}
			volInfoMapResponse["volumeType"] = "block"
			// for nas volume type
		} else if volumeType == "nfs" {
			if value, ok := volInfoMapFrom["server"]; ok {
				volInfoMapResponse["server"] = value
			}
			if value, ok := volInfoMapFrom["path"]; ok {
				volInfoMapResponse["path"] = value
			}
			if value, ok := volInfoMapFrom["vers"]; ok {
				volInfoMapResponse["vers"] = value
			} else {
				volInfoMapResponse["vers"] = "3"
			}
			if value, ok := volInfoMapFrom["mode"]; ok {
				volInfoMapResponse["mode"] = value
			} else {
				volInfoMapResponse["mode"] = ""
			}
			if value, ok := volInfoMapFrom["options"]; ok {
				volInfoMapResponse["options"] = value
			} else {
				volInfoMapResponse["options"] = "noresvport,nolock,tcp"
			}
			volInfoMapResponse["volumeType"] = "nfs"
		} else {
			klog.Errorf("Request volumeInfo: get error volumeType: %s for identity: %s", volumeType, reqInfo.Identity)
			fmt.Fprintf(w, "null")
			return
		}

		responseStr, err := json.Marshal(volInfoMapResponse)
		if err != nil {
			klog.Errorf("Request volumeInfo: Marshal volInfoResp error: %s", err.Error())
			fmt.Fprintf(w, "null")
			return
		}

		// Send response
		fmt.Fprint(w, string(responseStr))
		klog.Infof("Request volumeInfo: Send Successful Response with: %s", responseStr)
		return
	}

	// no found volume
	klog.Warningf("Request volumeInfo: Send Fail Response with: no found volume, %s", fileName)
	fmt.Fprintf(w, "no found volume: %s", fileName)
}

// pingHandler ping test
func (ks *QueryServer) pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping successful")
}
