package oss

import (
	"io/ioutil"
	"net/http"
)

const (
	METADATA_URL = "http://100.100.100.200/latest/meta-data/"
	INSTANCE_ID  = "instance-id"
)

// get host regionid, zoneid
func GetMetaData(resource string) string {
	resp, err := http.Get(METADATA_URL + resource)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
