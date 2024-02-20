package credentials

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/signers"
	alicred "github.com/aliyun/credentials-go/credentials"
	"github.com/sirupsen/logrus"
)

type SignerAdaptor struct {
	cred alicred.Credential
}

func NewSignerAdaptor(cred alicred.Credential) *SignerAdaptor {
	return &SignerAdaptor{cred: cred}
}

func (s *SignerAdaptor) GetName() string {
	return "HMAC-SHA1"
}

func (s *SignerAdaptor) GetType() string {
	return ""
}

func (s *SignerAdaptor) GetVersion() string {
	return "1.0"
}

func (s *SignerAdaptor) GetAccessKeyId() (string, error) {
	ak, err := s.cred.GetAccessKeyId()
	if ak == nil {
		return "", err
	}
	return *ak, err
}

func (s *SignerAdaptor) GetExtraParam() map[string]string {
	sts, err := s.cred.GetSecurityToken()
	if err != nil {
		logrus.Errorf("get security token failed: %v", err)
		return nil
	}
	if sts != nil {
		return map[string]string{"SecurityToken": *sts}
	}
	return nil
}

func (s *SignerAdaptor) Sign(stringToSign, secretSuffix string) string {
	sk, err := s.cred.GetAccessKeySecret()
	if err != nil {
		logrus.Errorf("get access key secret failed: %v", err)
		return ""
	}
	secret := *sk + secretSuffix
	return signers.ShaHmac1(stringToSign, secret)
}
