package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// DoRequest Http Post Request
func DoRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Get Response StatusCode %d, Response: %s", resp.StatusCode, body)
		return nil, errors.New(msg)
	}

	return body, nil
}
