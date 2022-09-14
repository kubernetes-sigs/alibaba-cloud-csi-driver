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
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
	oidc "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// ConfigPath the secret mount file
	ConfigPath = "/var/addon/token-config"
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
}

// KeyPairArtifacts is cert struct
type KeyPairArtifacts struct {
	Cert    *x509.Certificate
	Key     *rsa.PrivateKey
	CertPEM []byte
	KeyPEM  []byte
}

// CertOption is cert option
type CertOption struct {
	CAName          string
	CAOrganizations []string
	DNSNames        []string
	CommonName      string
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

func getManagedAddonToken() AccessControl {
	tokens := getManagedToken()
	return AccessControl{AccessKeyID: tokens.AccessKeyID, AccessKeySecret: tokens.AccessKeySecret, StsToken: tokens.SecurityToken, UseMode: ManagedToken}
}

// GetAccessControl  1、Read default ak from local file. 2、If local default ak is not exist, then read from STS.
func GetAccessControl() AccessControl {

	oidcToken := getOIDCToken()
	if oidcToken.AccessKeyID != "" {
		log.Info("Get AK: USE OIDC token")
		return oidcToken
	}

	//1、Get AK from Env
	acLocalAK := GetEnvAK()
	if len(acLocalAK.AccessKeyID) != 0 && len(acLocalAK.AccessKeySecret) != 0 {
		log.Info("Get AK: use ENV AK")
		return acLocalAK
	}

	//2、Get AK from Credential Files
	acCredentialAK := getCredentialAK()
	if acCredentialAK.Config != nil && acCredentialAK.Credential != nil {
		log.Info("Get AK: use Credential AK")
		return acCredentialAK
	}

	//3、Get AK from ManagedToken
	acAddonToken := getManagedAddonToken()
	if len(acAddonToken.AccessKeyID) != 0 {
		log.Info("Get AK: use Managed Addon Token")
		return acAddonToken
	}

	//4、Get AK from ECS StsToken
	acStsToken := getStsToken()
	log.Info("Get AK: use ECS RamRole Token")
	return acStsToken

}

var oidcProvider oidc.Provider

func getOIDCToken() AccessControl {

	if os.Getenv("USE_OIDC_AUTH_INNER") != "true" {
		return AccessControl{}
	}
	if oidcProvider != nil {
		log.Infof("getOIDCToken: use exists provider")
		resp, err := oidcProvider.GetStsTokenWithCache()
		if err != nil || resp == nil {
			log.Errorf("getOIDCtoken: failed to assume role with oidc : %++v", err)
			return AccessControl{}
		}
		return AccessControl{AccessKeyID: strings.TrimSpace(resp.Credentials.AccessKeyId), AccessKeySecret: strings.TrimSpace(resp.Credentials.AccessKeySecret), StsToken: strings.TrimSpace(resp.Credentials.SecurityToken), UseMode: OIDCToken}
	}

	regionID := os.Getenv("REGION_ID")
	if regionID == "" {
		regionID = RetryGetMetaData("region-id")
	}
	if regionID == "" {
		log.Error("getOIDCToken: failed to get regionid from metadata server")
		return AccessControl{}
	}
	ownerId := RetryGetMetaData("owner-account-id")
	log.Infof("getOIDCToken: cluster owner id: %v", ownerId)
	if ownerId == "" {
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
		log.Errorf("getOIDCtoken: get empty provider")
		return AccessControl{}
	}
	resp, err := oidcProvider.GetStsTokenWithCache()
	if err != nil || resp == nil {
		log.Errorf("getOIDCtoken: failed to assume role with oidc : %++v", err)
		return AccessControl{}
	}
	return AccessControl{AccessKeyID: strings.TrimSpace(resp.Credentials.AccessKeyId), AccessKeySecret: strings.TrimSpace(resp.Credentials.AccessKeySecret), StsToken: strings.TrimSpace(resp.Credentials.SecurityToken), UseMode: OIDCToken}

}

// GetEnvAK read ak from local ENV
func GetEnvAK() AccessControl {
	accessKeyID, accessSecret := "", ""
	accessKeyID = os.Getenv("ACCESS_KEY_ID")
	accessSecret = os.Getenv("ACCESS_KEY_SECRET")

	return AccessControl{AccessKeyID: strings.TrimSpace(accessKeyID), AccessKeySecret: strings.TrimSpace(accessSecret), UseMode: AccessKey}
}

// GetStsToken get STS token and token from ecs meta server
func getStsToken() AccessControl {
	roleAuth := RoleAuth{}
	subpath := "ram/security-credentials/"
	roleName, err := GetMetaData(subpath)
	if err != nil {
		log.Errorf("GetSTSToken: request roleName with error: %s", err.Error())
		return AccessControl{}
	}

	fullPath := filepath.Join(subpath, roleName)
	roleInfo, err := GetMetaData(fullPath)
	if err != nil {
		log.Errorf("GetSTSToken: request roleInfo with error: %s", err.Error())
		return AccessControl{}
	}

	err = json.Unmarshal([]byte(roleInfo), &roleAuth)
	if err != nil {
		log.Errorf("GetSTSToken: unmarshal roleInfo: %s, with error: %s", roleInfo, err.Error())
		return AccessControl{}
	}
	return AccessControl{AccessKeyID: roleAuth.AccessKeyID, AccessKeySecret: roleAuth.AccessKeySecret, StsToken: roleAuth.SecurityToken, UseMode: EcsRAMRole}
}

// GetManagedToken get ak from csi secret
func getManagedToken() (tokens ManageTokens) {
	var akInfo AKInfo
	if _, err := os.Stat(ConfigPath); err == nil {
		encodeTokenCfg, err := ioutil.ReadFile(ConfigPath)
		if err != nil {
			log.Errorf("failed to read token config, err: %v", err)
			return ManageTokens{}
		}
		err = json.Unmarshal(encodeTokenCfg, &akInfo)
		if err != nil {
			log.Errorf("error unmarshal token config: %v", err)
			return ManageTokens{}
		}
		keyring := akInfo.Keyring
		ak, err := Decrypt(akInfo.AccessKeyID, []byte(keyring))
		if err != nil {
			log.Errorf("failed to decode ak, err: %v", err)
			return ManageTokens{}
		}

		sk, err := Decrypt(akInfo.AccessKeySecret, []byte(keyring))
		if err != nil {
			log.Errorf("failed to decode sk, err: %v", err)
			return ManageTokens{}
		}

		token, err := Decrypt(akInfo.SecurityToken, []byte(keyring))
		if err != nil {
			log.Errorf("failed to decode token, err: %v", err)
			return ManageTokens{}
		}
		layout := "2006-01-02T15:04:05Z"
		t, err := time.Parse(layout, akInfo.Expiration)
		if err != nil {
			log.Errorf("Parse expiration error: %s", err.Error())
		}
		if t.Before(time.Now()) {
			log.Errorf("invalid token which is expired, expiration as: %s", akInfo.Expiration)
		}
		tokens.AccessKeyID = string(ak)
		tokens.AccessKeySecret = string(sk)
		tokens.SecurityToken = string(token)

		if akInfo.RoleAccessKeyID != "" && akInfo.RoleAccessKeySecret != "" {
			roleAK, err := Decrypt(akInfo.RoleAccessKeyID, []byte(keyring))
			if err != nil {
				log.Errorf("failed to decode role ak, err: %v", err)
				return ManageTokens{}
			}
			roleSK, err := Decrypt(akInfo.RoleAccessKeySecret, []byte(keyring))
			if err != nil {
				log.Errorf("failed to decode role sk, err : %v", err)
				return ManageTokens{}
			}
			tokens.RoleAccessKeyID = string(roleAK)
			tokens.RoleAccessKeySecret = string(roleSK)
		}
		tokens.RoleArn = akInfo.RoleArn
	}
	return tokens
}

// PKCS5UnPadding get pkc
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// Decrypt secret Decrypt
func Decrypt(s string, keyring []byte) ([]byte, error) {
	cdata, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Errorf("failed to decode base64 string, err: %v", err)
		return nil, err
	}
	block, err := aes.NewCipher(keyring)
	if err != nil {
		log.Errorf("failed to new cipher, err: %v", err)
		return nil, err
	}
	blockSize := block.BlockSize()

	iv := cdata[:blockSize]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cdata)-blockSize)

	blockMode.CryptBlocks(origData, cdata[blockSize:])

	origData = PKCS5UnPadding(origData)
	return origData, nil
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

// CreateCACert function is create cacert
func CreateCACert(option CertOption, begin, end time.Time) (*KeyPairArtifacts, error) {
	templ := &x509.Certificate{
		SerialNumber: big.NewInt(0),
		Subject: pkix.Name{
			CommonName:   option.CAName,
			Organization: option.CAOrganizations,
		},
		DNSNames:              option.DNSNames,
		NotBefore:             begin,
		NotAfter:              end,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("generating key: %s", err)
	}
	der, err := x509.CreateCertificate(rand.Reader, templ, templ, key.Public(), key)
	if err != nil {
		return nil, fmt.Errorf("creating certificate: %s", err)
	}
	certPEM, keyPEM, err := pemEncode(der, key)
	if err != nil {
		return nil, fmt.Errorf("encoding PEM: %s", err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, fmt.Errorf("parsing certificate: %s", err)
	}

	return &KeyPairArtifacts{Cert: cert, Key: key, CertPEM: certPEM, KeyPEM: keyPEM}, nil
}

// CreateCertPEM function is create cacert pem
func CreateCertPEM(option CertOption, ca *KeyPairArtifacts, begin, end time.Time, isClient bool) ([]byte, []byte, error) {
	sn, err := genSerialNum()
	if err != nil {
		return nil, nil, err
	}
	eks := []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	dnsNames := option.DNSNames
	if isClient {
		eks = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}
		dnsNames = nil
	}
	templ := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: option.CommonName,
		},
		DNSNames:              dnsNames,
		NotBefore:             begin,
		NotAfter:              end,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           eks,
		BasicConstraintsValid: true,
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("generating key: %s", err)
	}
	der, err := x509.CreateCertificate(rand.Reader, templ, ca.Cert, key.Public(), ca.Key)
	if err != nil {
		return nil, nil, fmt.Errorf("creating certificate: %s", err)
	}
	certPEM, keyPEM, err := pemEncode(der, key)
	if err != nil {
		return nil, nil, fmt.Errorf("encoding PEM: %s", err)
	}
	return certPEM, keyPEM, nil
}

func pemEncode(certificateDER []byte, key *rsa.PrivateKey) ([]byte, []byte, error) {
	certBuf := &bytes.Buffer{}
	if err := pem.Encode(certBuf, &pem.Block{Type: "CERTIFICATE", Bytes: certificateDER}); err != nil {
		return nil, nil, fmt.Errorf("encoding cert: %s", err)
	}
	keyBuf := &bytes.Buffer{}
	if err := pem.Encode(keyBuf, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}); err != nil {
		return nil, nil, fmt.Errorf("encoding key: %s", err)
	}
	return certBuf.Bytes(), keyBuf.Bytes(), nil
}

func genSerialNum() (*big.Int, error) {
	serialNumLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNum, err := rand.Int(rand.Reader, serialNumLimit)
	if err != nil {
		return nil, fmt.Errorf("serial number generation failure (%v)", err)
	}
	return serialNum, nil
}

// NewClientTLSFromFile function is new client with tls
func NewClientTLSFromFile(serverName, caFile, certFile, keyFile string) (credentials.TransportCredentials, error) {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append client certs")
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})
	return creds, nil
}

// NewServerTLSFromFile function is new server with tls
func NewServerTLSFromFile(caFile, certFile, keyFile string) (credentials.TransportCredentials, error) {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append client certs")
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
	return creds, nil
}

// getCredentialAK get credential and config from credential files.
func getCredentialAK() AccessControl {
	envProvider := provider.NewEnvProvider()
	profileProvider := provider.NewProfileProvider()
	pc := provider.NewProviderChain([]provider.Provider{envProvider, profileProvider})
	credential, err := pc.Resolve()
	if err != nil {
		if !strings.Contains(err.Error(), "No credential found") {
			log.Errorf("Failed to resolve an authentication provider: %v", err)
		}
	}
	config := sdk.NewConfig().WithScheme("https")
	return AccessControl{Config: config, Credential: credential, UseMode: Credential}
}
