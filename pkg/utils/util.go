/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

//DefaultOptions used for global ak
type DefaultOptions struct {
	Global struct {
		KubernetesClusterTag string
		AccessKeyID          string `json:"accessKeyID"`
		AccessKeySecret      string `json:"accessKeySecret"`
		Region               string `json:"region"`
	}
}

const (
	USER_AKID      = "/etc/.volumeak/akId"
	USER_AKSECRET  = "/etc/.volumeak/akSecret"
	METADATA_URL   = "http://100.100.100.200/latest/meta-data/"
	REGIONID_TAG   = "region-id"
	INSTANCEID_TAG = "instance-id"
	DEFAULT_REGION = "cn-hangzhou"
)

var KUBERNETES_ALICLOUD_IDENTITY = fmt.Sprintf("Kubernetes.Alicloud/CsiPlugin")

// Define STS Token Response
type RoleAuth struct {
	AccessKeyId     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

func Succeed(a ...interface{}) Result {
	return Result{
		Status:  "Success",
		Message: fmt.Sprint(a...),
	}
}

func NotSupport(a ...interface{}) Result {
	return Result{
		Status:  "Not supported",
		Message: fmt.Sprint(a...),
	}
}

func Fail(a ...interface{}) Result {
	return Result{
		Status:  "Failure",
		Message: fmt.Sprint(a...),
	}
}

// Result
type Result struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Device  string `json:"device,omitempty"`
}

// run shell command
func Run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: " + cmd + ", with out: " + string(out) + ", with error: " + err.Error())
	}
	return string(out), nil
}

func CreateDest(dest string) error {
	fi, err := os.Lstat(dest)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0777); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%v already exist but it's not a directory", dest)
	}
	return nil
}

func IsMounted(mountPath string) bool {
	cmd := fmt.Sprintf("mount | grep %s | grep -v grep | wc -l", mountPath)
	out, err := Run(cmd)
	if err != nil {
		log.Infof("IsMounted: exec error: %s, %s", cmd, err.Error())
		return false
	}
	if strings.TrimSpace(out) == "0" {
		return false
	}
	return true
}

func Umount(mountPath string) bool {
	cmd := fmt.Sprintf("umount %s", mountPath)
	_, err := Run(cmd)
	if err != nil {
		return false
	}
	return true
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

// Get regionid instanceid;
func GetRegionAndInstanceId() (string, string, error) {
	regionId, err := GetMetaData(REGIONID_TAG)
	if err != nil {
		return "", "", err
	}
	instanceId, err := GetMetaData(INSTANCEID_TAG)
	if err != nil {
		return "", "", err
	}
	return regionId, instanceId, nil
}

//GetMetaData get metadata from ecs meta-server
func GetMetaData(resource string) (string, error) {
	resp, err := http.Get(METADATA_URL + resource)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetRegionIdAndInstanceId(nodeName string) (string, string, error) {
	strs := strings.SplitN(nodeName, ".", 2)
	if len(strs) < 2 {
		return "", "", fmt.Errorf("failed to get regionID and instanceId from nodeName")
	}
	return strs[0], strs[1], nil
}

// save json data to file
func WriteJosnFile(obj interface{}, file string) error {
	maps := make(map[string]interface{})
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() != "" {
			maps[t.Field(i).Name] = v.Field(i).String()
		}
	}
	rankingsJson, _ := json.Marshal(maps)
	if err := ioutil.WriteFile(file, rankingsJson, 0644); err != nil {
		return err
	}
	return nil
}

// parse json to struct
func ReadJsonFile(file string) (map[string]string, error) {
	jsonObj := map[string]string{}
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &jsonObj)
	if err != nil {
		return nil, err
	}
	return jsonObj, nil
}

// GetLocalAK read ossfs ak from local or from secret file
func GetLocalAK() (string, string) {
	accessKeyID, accessSecret := "", ""
	//accessKeyID, accessSecret = GetLocalAK()
	//if accessKeyID == "" || accessSecret == "" {
	if IsFileExisting(USER_AKID) && IsFileExisting(USER_AKSECRET) {
		raw, err := ioutil.ReadFile(USER_AKID)
		if err != nil {
			log.Errorf("Read User AK ID file error: %s", err.Error())
			return "", ""
		}
		accessKeyID = string(raw)

		raw, err = ioutil.ReadFile(USER_AKSECRET)
		if err != nil {
			log.Errorf("Read User AK Secret file error: %s", err.Error())
			return "", ""
		}
		accessSecret = string(raw)
	}

	//}
	return strings.TrimSpace(accessKeyID), strings.TrimSpace(accessSecret)
}

// GetDefaultAK read default ak from local file or from STS
func GetDefaultAK() (string, string, string) {
	accessKeyID, accessSecret := GetLocalAK()

	accessToken := ""
	if accessKeyID == "" || accessSecret == "" {
		accessKeyID, accessSecret, accessToken = GetSTSAK()
	}

	return accessKeyID, accessSecret, accessToken

}

// GetSTSAK get STS AK and token from ecs meta server
func GetSTSAK() (string, string, string) {
	roleAuth := RoleAuth{}
	subpath := "ram/security-credentials/"
	roleName, err := GetMetaData(subpath)
	if err != nil {
		log.Errorf("GetSTSToken: request roleName with error: %s", err.Error())
		return "", "", ""
	}

	fullPath := filepath.Join(subpath, roleName)
	roleInfo, err := GetMetaData(fullPath)
	if err != nil {
		log.Errorf("GetSTSToken: request roleInfo with error: %s", err.Error())
		return "", "", ""
	}

	err = json.Unmarshal([]byte(roleInfo), &roleAuth)
	if err != nil {
		log.Errorf("GetSTSToken: unmarshal roleInfo: %s, with error: %s", roleInfo, err.Error())
		return "", "", ""
	}
	return roleAuth.AccessKeyId, roleAuth.AccessKeySecret, roleAuth.SecurityToken
}

func NewEcsClient(accessKeyId, accessKeySecret, accessToken string) (ecsClient *ecs.Client) {
	var err error
	if accessToken == "" {
		ecsClient, err = ecs.NewClientWithAccessKey(DEFAULT_REGION, accessKeyId, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		ecsClient, err = ecs.NewClientWithStsToken(DEFAULT_REGION, accessKeyId, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}
	return
}
