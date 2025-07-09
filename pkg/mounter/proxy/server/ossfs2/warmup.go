package ossfs2

import (
	"io"
	"io/fs"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
	"sync"
)

func worker(id int, entryChan <-chan string, bytesReadChan chan<- int64, perFileMaxBytes int64, wg *sync.WaitGroup) {
	defer wg.Done()

	// Function to read a single file (used if the entry is a file)
	readSingleFile := func(filePath string, workerBytesRead *int64) error {
		klog.Infof("Worker %d: Reading single file: %s", id, filePath)
		file, err := os.Open(filePath)
		if err != nil {
			klog.Errorf("Worker %d: Error opening file %s: %v", id, filePath, err)
			return err
		}
		defer file.Close()

		buffer := make([]byte, perFileMaxBytes)
		for {
			n, readErr := file.Read(buffer)
			if n > 0 {
				*workerBytesRead += int64(n)
			}
			if readErr == io.EOF {
				break
			}
			if readErr != nil {
				klog.Errorf("Worker %d: Error reading file %s: %v", id, filePath, readErr)
				return readErr
			}
		}
		return nil
	}

	walkAndRead := func(startPath string, workerBytesRead *int64) error {
		return filepath.WalkDir(startPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				klog.Errorf("Worker %d: Error accessing path %s: %v", id, path, err)
				return nil
			}

			// Skip the starting directory itself, as we're processing its contents
			if path == startPath && d.IsDir() {
				return nil
			}

			if !d.IsDir() {
				return readSingleFile(path, workerBytesRead)
			} else {
				klog.Infof("Worker %d: Visiting directory: %s", id, path)
			}
			return nil
		})
	}

	for entryPath := range entryChan {
		klog.Infof("Worker %d received task for: %s", id, entryPath)
		var workerBytesRead int64

		info, err := os.Stat(entryPath)
		if err != nil {
			klog.Errorf("Worker %d: Error stating entry %s: %v", id, entryPath, err)
			// Report 0 bytes for this failed entry and continue
			bytesReadChan <- 0
			continue
		}

		if info.IsDir() {
			// Process directory recursively
			err = walkAndRead(entryPath, &workerBytesRead)
			if err != nil {
				klog.Errorf("Worker %d: Error during recursive walk for %s: %v", id, entryPath, err)
			}
			klog.Infof("Worker %d finished directory task: %s (Read %d MiB)", id, entryPath, workerBytesRead/(1024*1024))

		} else {
			// Process single file
			err = readSingleFile(entryPath, &workerBytesRead)
			if err != nil {
				klog.Errorf("Worker %d: Error reading single file %s: %v", id, entryPath, err)
			}
			klog.Infof("Worker %d finished file task: %s (Read %d MiB)", id, entryPath, workerBytesRead/(1024*1024))
		}

		// Report bytes read for this task (either recursive walk or single file read)
		bytesReadChan <- workerBytesRead
	}

	klog.Infof("Worker %d finished", id)
}
