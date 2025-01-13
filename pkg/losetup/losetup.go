package losetup

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"k8s.io/klog/v2"
)

type LoopDevice struct {
	Name      string      `json:"name"`
	Sizelimit json.Number `json:"sizelimit"`
	Offset    json.Number `json:"offset"`
	Autoclear bool        `json:"autoclear"`
	RO        bool        `json:"ro"`
	BackFile  string      `json:"back-file"`
	Dio       bool        `json:"dio"`
	LogSec    json.Number `json:"log-sec"`
}

type option []string

func WithAssociatedFile(file string) option {
	return []string{"--associated", file}
}

func List(opts ...option) ([]LoopDevice, error) {
	args := []string{"--list", "--json", "--noheadings"}
	for _, opt := range opts {
		args = append(args, opt...)
	}
	output, err := execLosetupWithArgs(args...)
	if err != nil {
		return nil, err
	}
	// empty output when no devices found
	if len(output) == 0 {
		return nil, nil
	}
	var result struct {
		LoopDevices []LoopDevice `json:"loopdevices"`
	}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, err
	}
	return result.LoopDevices, nil
}

func Resize(name string) error {
	_, err := execLosetupWithArgs("--set-capacity", name)
	return err
}

func execLosetupWithArgs(arg ...string) ([]byte, error) {
	return execWithLog("losetup", arg...)
}

func execWithLog(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("failed to execute %q: %w", cmd, err)
	}
	klog.V(2).InfoS("command finished", "command", cmd, "output", string(output))
	return output, nil
}
