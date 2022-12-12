package downloader

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// Download fetches and download a resource into the destiantion path using
// simple HTTP calls
func DownloadFromURL(url, destination string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Request failed with status " + response.Status)
	}

	file, err := os.Create(destination)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, response.Body); err != nil {
		return err
	}

	return nil
}
