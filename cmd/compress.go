package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/model"
)

var compressCmd = &cobra.Command{
	Use:   "compress [filename]",
	Short: "Compress files in folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		files, err := getFilteredFiles(folder)
		must(err)

		err = compressFiles(filename, files)
		must(err)
	},
}

func init() {
	compressCmd.Flags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")

	RootCmd.AddCommand(compressCmd)
}

func compressFiles(filename string, files []model.File) error {
	zipfile, err := os.Create(filename + ".zip")
	if err != nil {
		return err
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	// Switch to targeted folder
	if folder != "" {
		os.Chdir(folder)
	}

	currentFolder, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, file := range files {
		relativeFolder, err := filepath.Rel(currentFolder, file.Folder)
		if err != nil {
			return err
		}

		path := filepath.Join(relativeFolder, file.File)
		if err = addFileToArchive(zipWriter, path); err != nil {
			return err
		}

		if enableLog {
			fmt.Printf("%v: Compress %v\n", time.Now().Format(time.RFC3339), path)
		}
	}

	if enableLog {
		fmt.Printf("Output zip file: %v\n", filename+".zip")
	}

	return nil
}

func addFileToArchive(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)

	return err
}
