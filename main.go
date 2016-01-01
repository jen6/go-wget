package main

import (
	"github.com/gosurui/uiprogress"
	"io"
	"net/http"
	"os"
)

var steps = []string{"Downloading File", "Finish"}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	downprogress := &Progress{Reader: resp.Body, Total: resp.ContentLength, Recv: 0}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
}
