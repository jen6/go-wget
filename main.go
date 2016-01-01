package main

import (
	"github.com/gosuri/uiprogress"
	"io"
	"net/http"
	"os"
)

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

	uiprogress.Start()
	bar := uiprogress.AddBar(int(resp.ContentLength))
	bar.AppendCompleted()
	bar.PrependElapsed()
	downprogress := &Progress{
		Reader: resp.Body,
		Total:  resp.ContentLength,
		Recv:   0,
		Bar:    bar,
	}

	_, err = io.Copy(out, downprogress)
	if err != nil {
		return err
	}
	return nil
}

func main() {
}
