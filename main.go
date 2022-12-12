package main

import (
	"flag"
	"fmt"
	"go_bootstrap/downloader"
	"go_bootstrap/util"
	"log"
	"os"
	"path/filepath"
)

var BASE_URL = "https://github.com/gustapinto/go_bootstrap/raw/main/templates/"
var VALID_TEMPLATES = []string{
	"basic",
	"api",
}

func main() {
	projectName := flag.String("project", "./", "The project folder name")
	template := flag.String("template", "basic", "The template to use")
	flag.Parse()

	if !util.Contains(VALID_TEMPLATES, *template) {
		fmt.Println("Invalid template name, please use one of the following:")

		for _, t := range VALID_TEMPLATES {
			fmt.Printf("- %s\n", t)
		}
	}

	templateZipFileName := *template + ".zip"
	tempFilePath, err := filepath.Abs(filepath.Join(os.TempDir(), templateZipFileName))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := downloader.DownloadFromURL(BASE_URL+templateZipFileName, tempFilePath); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := util.Unzip(tempFilePath, *projectName); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := os.Rename(*template, *projectName); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
