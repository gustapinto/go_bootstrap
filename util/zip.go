package util

import (
	"archive/zip"
	"io"
	"os"
)

// Unzip unzips a zip file into the destination folder
func Unzip(zipFilePath, destination string) error {
	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if err := unzipFile(file, destination); err != nil {
			return err
		}
	}

	return nil
}

// unzipFile unzips a single zip file
func unzipFile(file *zip.File, destination string) error {
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(file.Name, os.ModePerm); err != nil {
			return err
		}

		return nil
	}

	destinationFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	zippedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}

	return nil
}
