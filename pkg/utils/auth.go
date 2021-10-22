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

type KeyPairArtifacts struct {
	Cert    *x509.Certificate
	Key     *rsa.PrivateKey
	CertPEM []byte
	KeyPEM  []byte
}

type CertOption struct {
	CAName          string
	CAOrganizations []string
	DNSNames        []string
	CommonName      string
}

// GetDefaultAK read default ak from local file or from STS
func GetDefaultAK() (string, string, string) {
	accessKeyID, accessSecret := GetLocalAK()

	accessToken := ""
	if accessKeyID == "" || accessSecret == "" {
		tokens := GetManagedToken()
		accessKeyID, accessSecret, accessToken = tokens.AccessKeyID, tokens.AccessKeySecret, tokens.SecurityToken
		if accessKeyID != "" {
			log.Infof("Get AK: use Managed AK")
		}
	} else {
		log.Infof("Get AK: use Local AK")
	}

	if accessKeyID == "" || accessSecret == "" {
		accessKeyID, accessSecret, accessToken = GetSTSAK()
		log.Infof("Get AK: use STS")
	}

	return accessKeyID, accessSecret, accessToken
}

// GetLocalAK read ossfs ak from local or from secret file
func GetLocalAK() (string, string) {
	accessKeyID, accessSecret := "", ""
	accessKeyID = os.Getenv("ACCESS_KEY_ID")
	accessSecret = os.Getenv("ACCESS_KEY_SECRET")

	return strings.TrimSpace(accessKeyID), strings.TrimSpace(accessSecret)
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
	return roleAuth.AccessKeyID, roleAuth.AccessKeySecret, roleAuth.SecurityToken
}

// GetManagedToken get ak from csi secret
func GetManagedToken() (tokens ManageTokens) {
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
func GetDefaultRoleAK() (string, string, string) {
	accessKeyID, accessSecret, roleArn := os.Getenv("ROLE_ACCESS_KEY_ID"), os.Getenv("ROLE_ACCESS_KEY_SECRET"), os.Getenv("ROLE_ARN")
	if accessKeyID == "" || accessSecret == "" || roleArn == "" {
		tokens := GetManagedToken()
		accessKeyID, accessSecret, roleArn = tokens.RoleAccessKeyID, tokens.RoleAccessKeySecret, tokens.RoleArn
	}
	return accessKeyID, accessSecret, roleArn
}

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

func NewServerTLSFromFile(caFile, certFile, keyFile string)(credentials.TransportCredentials, error) {
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
		return nil,  errors.New("failed to append client certs")
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
	return creds, nil
}
