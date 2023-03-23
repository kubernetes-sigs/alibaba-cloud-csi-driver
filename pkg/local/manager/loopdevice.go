package manager

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	LOCAL_SPARSE_TEMPLATE_NAME = "ext4.img"
	LOCAL_SPARSE_DEFAULT_SIZE  = 10
)

// MaintainSparseTemplateFile create template dir for loopdeivce
func MaintainSparseTemplateFile(dir string, size string) error {
	ld := NewLoopDevice(dir, size)
	templateDir, templateSize := ld.GetTemplateInfo()
	err := EnsureFolder(templateDir)
	if err != nil {
		return err
	}
	tempFileName := filepath.Join(templateDir, LOCAL_SPARSE_TEMPLATE_NAME)
	_, err = os.Stat(tempFileName)
	if os.IsNotExist(err) {
		err = ld.CreateSparseFile(tempFileName, strconv.Itoa(int(utils.Gi2Bytes(int64(templateSize)))))
		if err != nil {
			return err
		}
	}
	return ld.FormatFile(tempFileName)
}

type LoopDevice interface {
	CreateSparseFile(sparseFile string, sizebit string) error
	FormatFile(fullName string) error
	CustomFormatFile(fullName, fsType string, options []string) error
	CopySparseFile(sourceFile, targetFile string) error
	CreateLoopDevice(sparseFile string) (string, error)
	DeleteLoopDevice(sparseFile string) (string, error)
	ResizeLoopDevice(sparseFile string) error
	FindLoopDeviceBySparseFile(sparseFile string) (string, error)
	GetTemplateInfo() (string, int)
	FileExists(filepath string) error
	GetUsedByteSize() (int64, error)
	GetTempDirTotalCapacity() (int64, error)
}

type NodeLoopDevice struct {
	templateDir  string
	templateSize int
}

func NewLoopDevice(templateDir, templateSize string) LoopDevice {

	size := LOCAL_SPARSE_DEFAULT_SIZE
	if value, err := strconv.Atoi(templateSize); err == nil {
		size = value
	}
	return &NodeLoopDevice{
		templateDir:  templateDir,
		templateSize: size,
	}
}

func (ld *NodeLoopDevice) GetTempDirTotalCapacity() (int64, error) {
	cmd := fmt.Sprintf("%s df %s | awk '!/1K-blocks/{print $2}'", NsenterCmd, ld.templateDir)
	out, err := utils.Run(cmd)
	if err != nil {
		return 0, err
	}
	out = strings.Trim(out, "\n")
	totalCapacity, err := strconv.Atoi(out)
	if err != nil {
		return 0, err
	}
	totalBytes := utils.KBlock2Bytes(int64(totalCapacity))
	return totalBytes, nil
}

func (ld *NodeLoopDevice) GetTemplateInfo() (string, int) {
	return ld.templateDir, ld.templateSize
}

func (ld *NodeLoopDevice) CreateSparseFile(sparseFile string, sizebit string) error {
	cmd := fmt.Sprintf("%s truncate -s %v %s", NsenterCmd, sizebit, sparseFile)
	_, err := utils.Run(cmd)
	return err
}

func (ld *NodeLoopDevice) FormatFile(fullName string) error {
	// fullName := filepath.Join(ld.tmpFileDir, ld.tmpFileName)
	// disable journal: we don't need journal in guest, and save some space
	// disable lazy_itable_init: save I/O when mounting ext4 in guest
	// disable discard: save I/O
	// enable packed_meta_blocks: pack metadata together, so one CoW could cover more metadata
	cmd := fmt.Sprintf("%s mkfs.ext4 -E packed_meta_blocks=1,nodiscard,lazy_itable_init=0 -O ^has_journal -F %s", NsenterCmd, fullName)
	_, err := utils.Run(cmd)
	return err
}

func (ld *NodeLoopDevice) CustomFormatFile(fullName, fsType string, options []string) error {

	if len(fsType) == 0 {
		fsType = "ext4"
	}
	args := []string{}
	if len(options) != 0 {
		for _, opts := range options {
			args = append(args, opts)
		}
	}
	args = append(args, fullName)
	_, err := exec.Command(fmt.Sprintf("%s mkfs.%s", utils.NsenterCmd, fsType), args...).CombinedOutput()
	return err

}

func (ld *NodeLoopDevice) CopySparseFile(sourceFile, targetFile string) error {
	cmd := fmt.Sprintf("%s cp %s %s", utils.NsenterCmd, sourceFile, targetFile)
	_, err := utils.Run(cmd)
	return err
	// source, err := os.Open(sourceFile)
	// if err != nil {
	// 	return err
	// }
	// defer source.Close()
	// des, err := os.Create(targetFile)
	// if err != nil {
	// 	return err
	// }
	// defer des.Close()
	// _, err = io.Copy(des, source)
	// return err
}

func (ld *NodeLoopDevice) CreateLoopDevice(sparseFile string) (string, error) {
	if err := ld.FileExists(sparseFile); err != nil {
		return "", fmt.Errorf("CreateLoopDevice: check sparsefile file: %s err:%v, failed to create loopdevice", sparseFile, err)
	}
	cmd := fmt.Sprintf("%s losetup -f %s", utils.NsenterCmd, sparseFile)
	out, err := utils.Run(cmd)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (ld *NodeLoopDevice) DeleteLoopDevice(sparseFile string) (string, error) {
	loopDevice, err := ld.FindLoopDeviceBySparseFile(sparseFile)
	if err != nil {
		return "", fmt.Errorf("DeleteLoopDevice: failed to find loopdevice by sparsefile err: %v", err)
	}

	cmd := fmt.Sprintf("%s losetup -d %s", utils.NsenterCmd, loopDevice)
	out, err := utils.Run(cmd)
	if err != nil {
		return "", err
	}

	sfCmd := fmt.Sprintf("%s rm %s", utils.NsenterCmd, sparseFile)
	out, err = utils.Run(sfCmd)
	if err != nil {
		return "", err
	}

	return out, nil
}

func (ld *NodeLoopDevice) FindLoopDeviceBySparseFile(sparseFile string) (string, error) {
	if err := ld.FileExists(sparseFile); err != nil {
		return "", fmt.Errorf("FindLoopDeviceBySparseFile: check sparsefile file: %s err:%v, failed to find sparsefile", sparseFile, err)
	}
	cmd := fmt.Sprintf("%s losetup | grep '%s' | awk '{print $1}'", utils.NsenterCmd, sparseFile)
	out, err := utils.Run(cmd)
	if err != nil {
		return "", err
	}
	return strings.Trim(out, "\n"), nil
}

func (ld *NodeLoopDevice) GetUsedByteSize() (int64, error) {
	var usedBytes int64
	cmd := fmt.Sprintf("%s ls -l %s | awk '!/total/{print $5}'", utils.NsenterCmd, ld.templateDir)
	out, err := utils.Run(cmd)
	if err != nil {
		return 0, err
	}
	for _, value := range strings.Split(out, "\n") {
		if value == "" {
			continue
		}
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		usedBytes += int64(i)
	}
	return usedBytes, nil
}

func (ld *NodeLoopDevice) ResizeLoopDevice(sparseFile string) error {
	// resize loopdevice
	loopDevice, err := ld.FindLoopDeviceBySparseFile(sparseFile)
	if err != nil {
		return fmt.Errorf("ResizeLoopDevice: failed to find sparsefile: %s err:%v", sparseFile, err)
	}

	cmd := fmt.Sprintf("%s losetup -c %s", NsenterCmd, loopDevice)
	out, err := utils.Run(cmd)
	if err != nil {
		return fmt.Errorf("ResizeLoopDevice: failed to resize loopdevice err:%v, out:%s, cmd: %s", err, out, cmd)
	}

	cmd = fmt.Sprintf("%s resize2fs %s", NsenterCmd, loopDevice)
	out, err = utils.Run(cmd)
	if err != nil {
		return fmt.Errorf("ResizeLoopDevice: failed to resize filesystem err:%v, out:%s", err, out)
	}
	return nil
}

func (ld *NodeLoopDevice) FileExists(filepath string) error {
	cmd := fmt.Sprintf("%s stat %s", NsenterCmd, filepath)
	_, err := utils.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}
