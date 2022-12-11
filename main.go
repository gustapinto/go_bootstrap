package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var TEMP_PATH = "/tmp/"
var VALID_TEMPLATES = []string{
	"basic",
	"api",
}

func contains[T comparable](container []T, item T) bool {
	for _, i := range container {
		if i == item {
			return true
		}
	}

	return false
}

func unzipIntoPath(template string, path string) error {
	reader, err := zip.OpenReader(filepath.Join(TEMP_PATH, template+".zip"))
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	for _, file := range reader.File {
		if err := unzipFile(file, path); err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(file *zip.File, destination string) error {
	path := filepath.Join(destination, file.Name)

	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}

		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
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

func main() {
	// projectName := flag.String("project", "", "The project folder name")
	template := flag.String("template", "basic", "The template to use")
	flag.Parse()

	if !contains(VALID_TEMPLATES, *template) {
		fmt.Println("Invalid template name, please use one of the following:")

		for _, t := range VALID_TEMPLATES {
			fmt.Printf("- %s\n", t)
		}
	}

	unzipIntoPath("basic", "/tmp/sla")
}
