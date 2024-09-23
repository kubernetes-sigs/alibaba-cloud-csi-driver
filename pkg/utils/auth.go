/*
Copyright 2020 The Alibaba Cloud Authors.

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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	cre "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
	crev2 "github.com/aliyun/credentials-go/credentials"
	oidc "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/auth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/crypto"
	"k8s.io/klog/v2"
)

const (
	// ConfigPath the secret mount file
	ConfigPath = "/var/addon/token-config"

	addonTokenExpirationFormat = "2006-01-02T15:04:05Z"
	addonTokenExpirationScale  = 0.9
)

// AKInfo access key info
type AKInfo struct {
	// AccessKeyId access key id
	AccessKeyID string `json:"access.key.id"`
	// AccessKeySecret access key secret
	AccessKeySecret string `json:"access.key.secret"`
	// SecurityToken security token
	SecurityToken string `json:"security.token"`
	// Expiration expiration duration
	Expiration string `json:"expiration"`
	// Keyring key ring
	Keyring string `json:"keyring"`
	// RoleAccessKeyId key
	RoleAccessKeyID string `json:"role.access.key.id"`
	// RoleAccessKeySecret key
	RoleAccessKeySecret string `json:"role.access.key.secret"`
	// RoleArn key
	RoleArn string `json:"role.arn"`
}

// ManageTokens 定义资源账号 和 角色扮演账号
type ManageTokens struct {
	// AccessKeyId key
	AccessKeyID string
	// AccessKeySecret key
	AccessKeySecret string
	// SecurityToken key
	SecurityToken string

	// RoleAccessKeyId key
	RoleAccessKeyID string
	// RoleAccessKeySecret key
	RoleAccessKeySecret string
	// RoleArn key
	RoleArn string

	// expire time
	ExpireAt time.Time
}

// AccessControlMode is int, represents different modes
type AccessControlMode int

// AccessControlMode includes AccessKey, ManagedToken, EcsRamRole, Credential, RoleArnToken, five types of access control
const (
	AccessKey AccessControlMode = iota
	ManagedToken
	EcsRAMRole
	Credential
	RoleArnToken
	OIDCToken
)

// AccessControl is access control option
type AccessControl struct {
	AccessKeyID     string
	AccessKeySecret string
	StsToken        string
	RoleArn         string
	Config          *sdk.Config
	Credential      auth.Credential
	UseMode         AccessControlMode
}

// CheckRequestArgs is check string is valid in args map
func CheckRequestArgs(m map[string]string) (bool, error) {
	valid := true
	var msg string
	for _, value := range m {
		if strings.Contains(value, "&") || strings.Contains(value, "|") || strings.Contains(value, ";") ||
			strings.Contains(value, "$") || strings.Contains(value, "'") || strings.Contains(value, "`") ||
			strings.Contains(value, "(") || strings.Contains(value, ")") {
			valid = false
			msg = msg + fmt.Sprintf("Args %s has illegal access.", value)
		}
	}
	return valid, errors.New(msg)
}

func ValidateRequest(m map[string]string) (bool, error) {
	valid := true
	var msg string
	for _, value := range m {
		if strings.Contains(value, "&") || strings.Contains(value, "|") || strings.Contains(value, ";") ||
			strings.Contains(value, "$") || strings.Contains(value, "'") || strings.Contains(value, "`") ||
			strings.Contains(value, "(") || strings.Contains(value, ")") {
			valid = false
			msg = msg + fmt.Sprintf("ValidateRequest: Args %s has illegal access.", value)
		}
		if pathValid, _ := ValidatePath(value); !pathValid {
			msg = msg + fmt.Sprintf("ValidateRequest: Args %s has illegal path", value)
			valid = false
		}
	}
	return valid, errors.New(msg)
}

// ValidatePath is check path string is valid
func ValidatePath(path string) (bool, error) {
	var msg string
	if strings.Contains(path, "../") || strings.Contains(path, "/..") || strings.Contains(path, "..") {
		msg = msg + fmt.Sprintf("Path %s has illegal access.", path)
		return false, errors.New(msg)
	}
	if strings.Contains(path, "./") || strings.Contains(path, "/.") {
		msg = msg + fmt.Sprintf("Path %s has illegal access.", path)
		return false, errors.New(msg)
	}

	return true, nil
}

func CheckRequest(m map[string]string, path string) (bool, error) {
	valid, err := CheckRequestArgs(m)
	if !valid {
		return valid, err
	}
	valid, err = ValidatePath(path)
	if !valid {
		return valid, err
	}
	return valid, nil
}

func getManagedAddonToken() AccessControl {
	tokens := getManagedToken()
	scheme := "https"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "http"
	}
	config := sdk.NewConfig().WithScheme(scheme)
	credent := &cre.StsTokenCredential{
		AccessKeyId:       tokens.AccessKeyID,
		AccessKeySecret:   tokens.AccessKeySecret,
		AccessKeyStsToken: tokens.SecurityToken,
	}

	return AccessControl{AccessKeyID: tokens.AccessKeyID, AccessKeySecret: tokens.AccessKeySecret, StsToken: tokens.SecurityToken, UseMode: ManagedToken, Config: config, Credential: credent}
}

// GetAccessControl  1、Read default ak from local file. 2、If local default ak is not exist, then read from STS.
func GetAccessControl() AccessControl {

	oidcToken := getOIDCToken()
	if oidcToken.AccessKeyID != "" {
		klog.Info("Get AK: USE OIDC token")
		return oidcToken
	}

	//1、Get AK from Env
	acLocalAK := GetEnvAK()
	if len(acLocalAK.AccessKeyID) != 0 && len(acLocalAK.AccessKeySecret) != 0 {
		klog.Info("Get AK: use ENV AK")
		return acLocalAK
	}

	//2、Get AK from Credential Files
	acCredentialAK := getCredentialAK()
	if acCredentialAK.Config != nil && acCredentialAK.Credential != nil {
		klog.Info("Get AK: use Credential AK")
		return acCredentialAK
	}

	//3、Get AK from ManagedToken
	acAddonToken := getManagedAddonToken()
	if len(acAddonToken.AccessKeyID) != 0 {
		klog.Info("Get AK: use Managed Addon Token")
		return acAddonToken
	}

	//4、Get AK from ECS StsToken
	acStsToken := getStsToken()
	klog.Info("Get AK: use ECS RamRole Token")
	return acStsToken
}

var oidcProvider oidc.Provider

func getOIDCToken() AccessControl {

	if os.Getenv("USE_OIDC_AUTH_INNER") != "true" {
		return AccessControl{}
	}
	if oidcProvider != nil {
		klog.Infof("getOIDCToken: use exists provider")
		resp, err := oidcProvider.GetStsTokenWithCache()
		if err != nil || resp == nil {
			klog.Errorf("getOIDCtoken: failed to assume role with oidc : %++v", err)
			return AccessControl{}
		}
		return AccessControl{AccessKeyID: strings.TrimSpace(resp.Credentials.AccessKeyId), AccessKeySecret: strings.TrimSpace(resp.Credentials.AccessKeySecret), StsToken: strings.TrimSpace(resp.Credentials.SecurityToken), UseMode: OIDCToken}
	}

	regionID := os.Getenv("REGION_ID")
	if regionID == "" {
		regionID = RetryGetMetaData("region-id")
	}
	if regionID == "" {
		klog.Error("getOIDCToken: failed to get regionid from metadata server")
		return AccessControl{}
	}
	ownerId := os.Getenv("ACCOUNT_ID")
	if ownerId == "" {
		ownerId = RetryGetMetaData("owner-account-id")
	}
	klog.Infof("getOIDCToken: cluster owner id: %v", ownerId)
	if ownerId == "" {
		klog.Error("getOIDCToken: failed to get cluster owner id from metadata server")
		return AccessControl{}
	}

	oidcProvider = oidc.NewOIDCProviderVPC(
		regionID,
		"alibaba-cloud-csi-controller",
		"alibaba-cloud-csi-controller-oidc-provider",
		"alibaba-cloud-csi-controller-oidc-role",
		ownerId,
		time.Duration(1000)*time.Second)
	if oidcProvider == nil {
		klog.Errorf("getOIDCtoken: get empty provider")
		return AccessControl{}
	}
	resp, err := oidcProvider.GetStsTokenWithCache()
	if err != nil || resp == nil {
		klog.Errorf("getOIDCtoken: failed to assume role with oidc : %++v", err)
		return AccessControl{}
	}
	return AccessControl{AccessKeyID: strings.TrimSpace(resp.Credentials.AccessKeyId), AccessKeySecret: strings.TrimSpace(resp.Credentials.AccessKeySecret), StsToken: strings.TrimSpace(resp.Credentials.SecurityToken), UseMode: OIDCToken}

}

// GetEnvAK read ak from local ENV
func GetEnvAK() AccessControl {
	accessKeyID, accessSecret := "", ""
	accessKeyID = os.Getenv("ACCESS_KEY_ID")
	accessSecret = os.Getenv("ACCESS_KEY_SECRET")

	scheme := "https"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "http"
	}
	config := sdk.NewConfig().WithScheme(scheme)

	credent := &cre.AccessKeyCredential{
		AccessKeyId:     strings.TrimSpace(accessKeyID),
		AccessKeySecret: strings.TrimSpace(accessSecret),
	}
	return AccessControl{AccessKeyID: strings.TrimSpace(accessKeyID), AccessKeySecret: strings.TrimSpace(accessSecret), UseMode: AccessKey, Config: config, Credential: credent}
}

// GetStsToken get STS token and token from ecs meta server
func getStsToken() AccessControl {
	roleAuth := RoleAuth{}
	subpath := "ram/security-credentials/"
	roleName, err := GetMetaData(subpath)
	if err != nil {
		klog.Errorf("GetSTSToken: request roleName with error: %s", err.Error())
		return AccessControl{}
	}

	fullPath := filepath.Join(subpath, roleName)
	roleInfo, err := GetMetaData(fullPath)
	if err != nil {
		klog.Errorf("GetSTSToken: request roleInfo with error: %s", err.Error())
		return AccessControl{}
	}

	err = json.Unmarshal([]byte(roleInfo), &roleAuth)
	if err != nil {
		klog.Errorf("GetSTSToken: unmarshal roleInfo: %s, with error: %s", roleInfo, err.Error())
		return AccessControl{}
	}
	scheme := "https"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "http"
	}
	config := sdk.NewConfig().WithScheme(scheme)
	credent := &cre.StsTokenCredential{
		AccessKeyId:       roleAuth.AccessKeyID,
		AccessKeySecret:   roleAuth.AccessKeySecret,
		AccessKeyStsToken: roleAuth.SecurityToken,
	}
	return AccessControl{AccessKeyID: roleAuth.AccessKeyID, AccessKeySecret: roleAuth.AccessKeySecret, StsToken: roleAuth.SecurityToken, UseMode: EcsRAMRole, Config: config, Credential: credent}
}

// GetManagedToken get ak from csi secret
func getManagedToken() (tokens ManageTokens) {
	var akInfo crypto.RamToken
	if _, err := os.Stat(ConfigPath); err == nil {
		encodeTokenCfg, err := os.ReadFile(ConfigPath)
		if err != nil {
			klog.Errorf("failed to read token config, err: %v", err)
			return ManageTokens{}
		}
		err = json.Unmarshal(encodeTokenCfg, &akInfo)
		if err != nil {
			klog.Errorf("failed to unmarshal token config: %v", err)
			return ManageTokens{}
		}
		newToken, err := crypto.RamTokenDecrypt(&akInfo)
		if err != nil {
			klog.Errorf("failed to decrypt new token: %v", err)
			return ManageTokens{}
		}
		expireAt, err := time.Parse("2006-01-02T15:04:05Z", newToken.Expiration)
		if err != nil {
			klog.Errorf("failed to parse expiration: %q: %v", newToken.Expiration, err)
			return ManageTokens{}
		}
		tokens.AccessKeyID = newToken.AccessKeyId
		tokens.AccessKeySecret = newToken.AccessKeySecret
		tokens.SecurityToken = newToken.SecurityToken
		tokens.ExpireAt = expireAt
	}
	return tokens
}

// GetDefaultRoleAK  返回角色扮演账号AK, SK, role arn
func GetDefaultRoleAK() AccessControl {
	roleAccessKeyID, roleAccessKeySecret, roleArn := os.Getenv("ROLE_ACCESS_KEY_ID"), os.Getenv("ROLE_ACCESS_KEY_SECRET"), os.Getenv("ROLE_ARN")
	if len(roleAccessKeyID) == 0 || len(roleAccessKeySecret) == 0 || len(roleArn) == 0 {
		tokens := getManagedToken()
		return AccessControl{AccessKeyID: tokens.RoleAccessKeyID, AccessKeySecret: tokens.RoleAccessKeySecret, RoleArn: tokens.RoleArn, UseMode: ManagedToken}
	}
	return AccessControl{AccessKeyID: roleAccessKeyID, AccessKeySecret: roleAccessKeySecret, RoleArn: roleArn, UseMode: RoleArnToken}
}

// getCredentialAK get credential and config from credential files.
func getCredentialAK() AccessControl {
	envProvider := provider.NewEnvProvider()
	profileProvider := provider.NewProfileProvider()
	pc := provider.NewProviderChain([]provider.Provider{envProvider, profileProvider})
	credential, err := pc.Resolve()
	if err != nil {
		if !strings.Contains(err.Error(), "No credential found") {
			klog.Errorf("Failed to resolve an authentication provider: %v", err)
		}
	}
	scheme := "https"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "http"
	}
	config := sdk.NewConfig().WithScheme(scheme)
	return AccessControl{Config: config, Credential: credential, UseMode: Credential}
}

type managedAddonTokenCredv2 struct {
	sync.Mutex
	token        *ManageTokens
	lastUpdateAt time.Time
	scale        float64
}

func newManagedAddonTokenCredv2() *managedAddonTokenCredv2 {
	return &managedAddonTokenCredv2{
		scale: addonTokenExpirationScale,
	}
}

func (cred *managedAddonTokenCredv2) needUpdate() bool {
	if cred.token == nil {
		return true
	}
	duration := time.Since(cred.lastUpdateAt)
	expiration := cred.token.ExpireAt.Sub(cred.lastUpdateAt)
	return duration >= time.Duration(float64(expiration)*cred.scale)
}

func (cred *managedAddonTokenCredv2) updateAndGet() ManageTokens {
	cred.Lock()
	defer cred.Unlock()
	if cred.needUpdate() {
		tokens := getManagedToken()
		cred.token = &tokens
		cred.lastUpdateAt = time.Now()
	}
	return *cred.token
}

func (cred *managedAddonTokenCredv2) GetAccessKeyId() (*string, error) {
	token := cred.updateAndGet()
	return &token.AccessKeyID, nil
}

func (cred *managedAddonTokenCredv2) GetAccessKeySecret() (*string, error) {
	token := cred.updateAndGet()
	return &token.AccessKeySecret, nil
}

func (cred *managedAddonTokenCredv2) GetSecurityToken() (*string, error) {
	token := cred.updateAndGet()
	return &token.SecurityToken, nil
}

func (cred *managedAddonTokenCredv2) GetCredential() (*crev2.CredentialModel, error) {
	token := cred.updateAndGet()
	return &crev2.CredentialModel{
		AccessKeyId:     &token.AccessKeyID,
		AccessKeySecret: &token.AccessKeySecret,
		SecurityToken:   &token.SecurityToken,
		BearerToken:     tea.String(""),
		Type:            tea.String("sts"),
	}, nil
}

func (cred *managedAddonTokenCredv2) GetBearerToken() *string {
	return tea.String("")
}

func (cred *managedAddonTokenCredv2) GetType() *string {
	return tea.String("sts")
}

func GetCredentialV2() (crev2.Credential, error) {
	// env variable
	acLocalAK := GetEnvAK()
	if len(acLocalAK.AccessKeyID) != 0 && len(acLocalAK.AccessKeySecret) != 0 {
		klog.Info("credential v2: using ak from env variables")
		config := new(crev2.Config).SetType("access_key").
			SetAccessKeyId(acLocalAK.AccessKeyID).
			SetAccessKeySecret(acLocalAK.AccessKeySecret)
		return crev2.NewCredential(config)
	}

	// try default credential chain
	cred, err := crev2.NewCredential(nil)
	if err == nil {
		klog.Info("credential v2: using default credential chain")
		return cred, nil
	}

	// managed addon token
	_, err = os.Stat(ConfigPath)
	if err == nil {
		klog.Info("credential v2: using managed addon token")
		return newManagedAddonTokenCredv2(), nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}

	// ecs ram role
	klog.Info("credential v2: using ecs_ram_role")
	config := new(crev2.Config).SetType("ecs_ram_role")
	return crev2.NewCredential(config)
}
