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
	data, err := os.ReadFile(lingjunConfigFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			klog.V(3).InfoS("lingjun metadata file does not exist, use env metadata only", "file", lingjunConfigFile)
		} else {
			klog.ErrorS(err, "Failed to read lingjun metadata file", "file", lingjunConfigFile)
		}
		return nil, err
	}
	var lm LingjunMetaData
	err = json.Unmarshal(data, &lm)
	if err != nil {
		klog.ErrorS(err, "Failed to parse lingjun metadata file", "file", lingjunConfigFile)
		return nil, err
	}
	return &lm, nil
}

func (m *LingjunMetaData) GetAny(_ *mcontext, key MetadataKey) (any, error) {
	switch key {
	case RegionID:
		return m.RegionId, nil
	case ZoneID:
		return m.ZoneId, nil
	case InstanceID:
		return m.NodeId, nil
	case InstanceType:
		return m.InstanceType, nil
	case machineKind:
		return MachineKindLingjun, nil
	default:
		return "", ErrUnknownMetadataKey
	}
}
