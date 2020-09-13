package model

import (
	"path/filepath"
	"strings"
)

type File struct {
	FullPath string // The full path including the folder
	Folder   string // The folder of the file
	File     string // The file name including extension
	FileName string // The file name excluding extension
	Ext      string // Extension of the file
}

// Construct from string
func ConstructFile(path string) File {
	base := filepath.Base(path)
	ext := filepath.Ext(path)

	return File{
		FullPath: path,
		File:     base,
		Folder:   strings.TrimSuffix(path, base),
		FileName: strings.TrimSuffix(base, ext),
		Ext:      ext,
	}
}
