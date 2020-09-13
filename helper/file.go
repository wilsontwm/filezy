package helper

import (
	"os"
	"path/filepath"

	"github.com/wilsontwm/filezy/model"
)

// HasFile : Check if file exists in the current directory
func HasFile(filename string) bool {
	if info, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return !info.IsDir()
	}
}

// GetFiles :
func GetFiles(folder string, isRecursive bool) ([]model.File, error) {
	var files []model.File

	if isRecursive {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, model.ConstructFile(path))
			}
			return nil
		})

		return files, err
	} else {
		f, err := os.Open(folder)
		defer f.Close()

		if err != nil {
			return files, err
		}

		if fileinfo, err := f.Readdir(-1); err == nil {
			for _, file := range fileinfo {
				if !file.IsDir() {
					files = append(files, model.ConstructFile(file.Name()))
				}
			}
		} else {
			return files, err
		}

	}

	return files, nil
}
