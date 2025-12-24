package utils

import (
	"os"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

var (
	// Store the read-write lock corresponding to each file path
	fileLocks = sync.Map{}
)

func getFileLock(path string) *sync.RWMutex {
	lock, _ := fileLocks.LoadOrStore(path, &sync.RWMutex{})
	return lock.(*sync.RWMutex)
}

// WriteFileWithLock safely writes data to file with locking
func WriteFileWithLock(path string, data []byte, perm os.FileMode) (done bool, err error) {
	lock := getFileLock(path)

	// First try to acquire read lock to check if content is consistent
	lock.RLock()
	if existingData, err := os.ReadFile(path); err == nil {
		// If file exists and content is the same, return directly to avoid redundant write
		if string(existingData) == string(data) {
			lock.RUnlock()
			return false, nil
		}
	}
	lock.RUnlock()

	// Content is different or file does not exist, need to write new content
	// Acquire write lock
	lock.Lock()
	defer lock.Unlock()

	// Check content again (double-checked locking pattern)
	if existingData, err := os.ReadFile(path); err == nil {
		if string(existingData) == string(data) {
			return false, nil
		}
	}

	// Perform write operation
	err = utils.WriteAndSyncFile(path, data, perm)
	if err == nil {
		done = true
	}
	return
}
