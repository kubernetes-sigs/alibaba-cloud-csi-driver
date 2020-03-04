package agent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

func NewQueryServer() *QueryServer {
	cfg, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	return &QueryServer{
		client: kubeClient,
	}
}

// RunAgent Routers
func (ks *QueryServer) RunQueryServer() {
	socketAddr := &net.UnixAddr{Name: QueryServerSocket, Net: "unix"}
	os.Remove(socketAddr.Name)
	lis, err := net.ListenUnix("unix", socketAddr)
	if err != nil {
		log.Errorf("Listen Unix error: %s", err.Error())
		return
	}

	// set router
	log.Infof("Started Query Server with unix socket: %s", QueryServerSocket)
	http.HandleFunc("/api/v1/volumeinfo", ks.volumeInfoHandler)
	http.HandleFunc("/api/v1/podruntime", ks.podRunTimeHander)
	http.HandleFunc("/api/v1/ping", ks.pingHandler)

	// Server Listen
	svr := &http.Server{Handler: http.DefaultServeMux}
	err = svr.Serve(lis)
	if err != nil {
		log.Errorf("Query Server Starting error: %s", err.Error())
	}
	log.Infof("Query Server Ending ....")
}

// volumeInfoHandler reply with volume options.
func (ks *QueryServer) volumeInfoHandler(w http.ResponseWriter, r *http.Request) {
	reqInfo := QueryRequest{}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Request volumeInfo: Receive request read body error: %s", err.Error())
		fmt.Fprintf(w, "null")
		return
	}
	if err := json.Unmarshal(content, &reqInfo); err != nil {
		log.Errorf("Request volumeInfo: Unmarshal request body(%s) error: %s", string(content), err.Error())
		fmt.Fprintf(w, "null")
		return
	}
	log.Infof("Request volumeInfo: Receive Request with identity: %s", reqInfo.Identity)
	if reqInfo.Identity == "" {
		fmt.Fprintf(w, "null")
		return
	}

	// Response with file content
	fileName := filepath.Join(reqInfo.Identity, CsiPluginRunTimeFlagFile)
	if utils.IsFileExisting(fileName) {
		// Unmarshal file content to map
		fileContent := utils.GetFileContent(fileName)
		fileContent = strings.ToLower(fileContent)
		volInfoMapFrom := map[string]string{}
		if err := json.Unmarshal([]byte(fileContent), &volInfoMapFrom); err != nil {
			log.Errorf("Request volumeInfo: Unmarshal fileContent (%s) error: %s", fileContent, err.Error())
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
			log.Errorf("Request volumeInfo: get error volumeType: %s for identity: %s", volumeType, reqInfo.Identity)
			fmt.Fprintf(w, "null")
			return
		}

		responseStr, err := json.Marshal(volInfoMapResponse)
		if err != nil {
			log.Errorf("Request volumeInfo: Marshal volInfoResp error: %s", err.Error())
			fmt.Fprintf(w, "null")
			return
		}

		// Send response
		fmt.Fprintf(w, string(responseStr))
		log.Infof("Request volumeInfo: Send Successful Response with: %s", responseStr)
		return
	}

	// no found volume
	log.Warnf("Request volumeInfo: Send Fail Response with: no found volume, %s", fileName)
	fmt.Fprintf(w, "no found volume: %s", fileName)
	return

}

// podRunTimeHander used for CSI, get pod runtime info.
func (ks *QueryServer) podRunTimeHander(w http.ResponseWriter, r *http.Request) {
	reqInfo := QueryRequest{}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Request RunTime: Receive Request Body with error: %s", err.Error())
		fmt.Fprintf(w, Unkown_RunTime_Tag)
		return
	}
	if err := json.Unmarshal(content, &reqInfo); err != nil {
		log.Errorf("Request RunTime: Unmarshal request body(%s) error: %s", string(content), err.Error())
		fmt.Fprintf(w, Unkown_RunTime_Tag)
		return
	}

	if reqInfo.PodNameSpace == "" || reqInfo.PodName == "" {
		log.Errorf("Request RunTime: Request with Empty PodInfo: %s, %s", reqInfo.PodName, reqInfo.PodNameSpace)
		fmt.Fprintf(w, Unkown_RunTime_Tag)
		return
	}
	log.Infof("Request RunTime: Receive Request with Pod: %s, NameSpace: %s", reqInfo.PodName, reqInfo.PodNameSpace)

	podInfo, err := ks.client.CoreV1().Pods(reqInfo.PodNameSpace).Get(reqInfo.PodName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("Request RunTime: Get PodInfo(%s, %s) with error: %s", reqInfo.PodName, reqInfo.PodNameSpace, err.Error())
		fmt.Fprintf(w, Unkown_RunTime_Tag)
		return
	}
	if podInfo.Spec.RuntimeClassName == nil {
		log.Infof("Request RunTime: Get with no runtime(nil), %s, %s", reqInfo.PodName, reqInfo.PodNameSpace)
		fmt.Fprintf(w, Runc_RunTime_Tag)
		return
	} else if *podInfo.Spec.RuntimeClassName == "" {
		log.Infof("Request RunTime: Get with empty runtime: %s, %s", reqInfo.PodName, reqInfo.PodNameSpace)
		fmt.Fprintf(w, Runc_RunTime_Tag)
		return
	} else {
		log.Infof("Request RunTime: Send for Pod: %s, %s, with runtime: %s", reqInfo.PodName, reqInfo.PodNameSpace, *podInfo.Spec.RuntimeClassName)
		fmt.Fprintf(w, *podInfo.Spec.RuntimeClassName)
		return
	}
}

// pingHandler ping test
func (ks *QueryServer) pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping successful")
}
