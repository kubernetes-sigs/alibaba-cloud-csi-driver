package dara

import (
	"os"
)

// File struct to represent the file
type DaraFile struct {
	path     string
	fileInfo os.FileInfo
	file     *os.File
	position int64
}

// NewFile creates a new instance of File
func NewDaraFile(path string) *DaraFile {
	return &DaraFile{
		path:     path,
		position: 0,
	}
}

// Path returns the path of the file
func (tf *DaraFile) Path() string {
	return tf.path
}

// CreateTime returns the creation time of the file
func (tf *DaraFile) CreateTime() (*Date, error) {
	if tf.fileInfo == nil {
		var err error
		tf.fileInfo, err = os.Stat(tf.path)
		if err != nil {
			return nil, err
		}
	}
	return &Date{tf.fileInfo.ModTime()}, nil
}

// ModifyTime returns the modification time of the file
func (tf *DaraFile) ModifyTime() (*Date, error) {
	if tf.fileInfo == nil {
		var err error
		tf.fileInfo, err = os.Stat(tf.path)
		if err != nil {
			return nil, err
		}
	}
	return &Date{tf.fileInfo.ModTime()}, nil
}

// Length returns the size of the file
func (tf *DaraFile) Length() (int64, error) {
	if tf.fileInfo == nil {
		var err error
		tf.fileInfo, err = os.Stat(tf.path)
		if err != nil {
			return 0, err
		}
	}
	return tf.fileInfo.Size(), nil
}

// Read reads a specified number of bytes from the file
func (tf *DaraFile) Read(size int) ([]byte, error) {
	if tf.file == nil {
		file, err := os.OpenFile(tf.path, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, err
		}
		tf.file = file
	}

	fileInfo, err := tf.file.Stat()
	if err != nil {
		return nil, err
	}

	// 获取文件大小
	fileSize := fileInfo.Size()

	// 计算可以读取的实际大小
	if tf.position >= fileSize {
		return nil, nil // End of file reached
	}

	// 确保 size 不超过剩余文件大小
	actualSize := size
	if tf.position+int64(size) > fileSize {
		actualSize = int(fileSize - tf.position)
	}

	buf := make([]byte, actualSize)
	bytesRead, err := tf.file.ReadAt(buf, tf.position)
	if err != nil {
		return nil, err
	}
	tf.position += int64(bytesRead)
	return buf[:bytesRead], nil
}

// Write writes data to the file
func (tf *DaraFile) Write(data []byte) error {
	if tf.file == nil {
		file, err := os.OpenFile(tf.path, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		tf.file = file
	}

	_, err := tf.file.Write(data)
	if err != nil {
		return err
	}

	tf.fileInfo, err = os.Stat(tf.path) // Update fileInfo after write
	return err
}

// Close closes the file
func (tf *DaraFile) Close() error {
	if tf.file == nil {
		return nil
	}
	return tf.file.Close()
}

// Exists checks if the file exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

// CreateReadStream would typically return an os.File or similar
func CreateReadStream(path string) (*os.File, error) {
	return os.Open(path)
}

// CreateWriteStream would typically return an os.File or similar
func CreateWriteStream(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
}
