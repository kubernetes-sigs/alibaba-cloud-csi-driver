package metric

import (
	"context"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
)

//UDSClient is unix domain socket client with server client
type UDSClient struct {
	client http.Client
}

func NewUDSClient(socketPath string) *UDSClient {
	httpClient := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}
	return &UDSClient{httpClient}
}

func (u *UDSClient) Get(url string) (*http.Response, error) {
	resp, err := u.client.Get(url)
	if err != nil {
		log.Errorf("Get url %s is failed, err: %s ", url, err)
		return nil, err
	}
	return resp, nil
}

//ReadBody to read body from response
func (u *UDSClient) ReadBody(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Get resp.body is failed, resp:%+v, err:%s", resp, err.Error())
		return string("")
	}
	return string(body)
}

//Close to close response
func (u *UDSClient) Close(resp *http.Response) {
	if resp != nil {
		_ = resp.Body.Close()
	}
	u.client.CloseIdleConnections()
}
