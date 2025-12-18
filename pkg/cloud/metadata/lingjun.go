package metadata

import (
	"encoding/json"
	"errors"
	"os"

	"k8s.io/klog/v2"
)

type LingjunMetaData struct {
	RegionId     string `json:"RegionId"`
	ZoneId       string `json:"ZoneId"`
	NodeId       string `json:"NodeId"`
	InstanceType string `json:"InstanceType"`
}

func NewLingJunMetadata(lingjunConfigFile string) (*LingjunMetaData, error) {
	lm := LingjunMetaData{}
	data, err := os.ReadFile(lingjunConfigFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			klog.V(3).InfoS("lingjun metadata file does not exist, use env metadata only", "file", lingjunConfigFile)
		} else {
			klog.ErrorS(err, "Failed to read lingjun metadatafile", "file", lingjunConfigFile)
		}
		return nil, err
	}
	err = json.Unmarshal(data, &lm)
	if err != nil {
		return nil, err
	}
	return &lm, nil
}

func (m *LingjunMetaData) Get(key MetadataKey) (string, error) {

	switch key.String() {
	case "RegionID":
		return m.RegionId, nil
	case "ZoneID":
		return m.ZoneId, nil
	case "InstanceID":
		return m.NodeId, nil
	case "InstanceType":
		return m.InstanceType, nil
	default:
		return "", ErrUnknownMetadataKey
	}
}
