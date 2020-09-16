package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
			file, err := filepath.Abs(path)
			if err != nil {
				return nil
			}

			if !info.IsDir() {
				files = append(files, model.ConstructFile(file))
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
					folder, err := filepath.Abs(folder)
					if err != nil {
						return files, err
					}
					files = append(files, model.ConstructFile(folder+"\\"+file.Name()))
				}
			}
		} else {
			return files, err
		}

	}

	return files, nil
}

// GetNewFilePath :
func GetNewFilePath(file model.File, sourceFolder, targetFolder string) (string, error) {
	if file.FullPath == "" {
		return "", fmt.Errorf("Original file is not valid.")
	}

	subfolder := GetSubfolder(file, sourceFolder)

	return targetFolder + subfolder + file.File, nil
}

// GetSubfolder :
func GetSubfolder(file model.File, sourceFolder string) string {
	return strings.TrimPrefix(file.Folder, sourceFolder)
}
