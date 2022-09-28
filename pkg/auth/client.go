package auth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"sync"
)

// Client is the sdk client struct, each func corresponds to an OpenAPI
type Client struct {
	sdk.Client
	lock sync.RWMutex
}

func (c *Client) getSigner() auth.Signer {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.GetSigner()
}

func (c *Client) setSigner(signer auth.Signer) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.SetSigner(signer)
}

// NewClientWithStsToken is a shortcut to create sdk client with sts token
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken)
	return
}
