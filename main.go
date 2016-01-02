package main

import (
	"fmt"
	"github.com/gosuri/uiprogress"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	DefaultPath = "./"
)

func parseFilename(url string) string {
	r := regexp.MustCompile("[\\w\\.]+")
	result := r.FindAllString(url, -1)
	return result[len(result)-1]
}

func makePath(filepath, url string) (string, error) {
	finfo, err := os.Stat(filepath)
	if err == nil {
		if finfo.IsDir() {
			//파일일 경우
			return filepath + "/" + parseFilename(url), nil
		} else {
			//디렉토리일경우
			return filepath, nil
		}
	}

	f, err := os.Create(filepath)
	if err != nil {
		log.Println("File path is not exist")
		return filepath, err
	}
	defer f.Close()
	return filepath, nil
}

func downloadFile(filepath string, url string) (err error) {

	if filepath == "" {
		filepath = DefaultPath
	}

	path, err := makePath(filepath, url)
	fmt.Println(path)
	if err != nil {
		fmt.Println("die1")
		os.Exit(1)
	}

	out, err := os.Create(path)
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
func parseArg(args []string) (string, string) {
	ArgLen := len(args)
	switch {
	case ArgLen <= 1:
		usage()
		return "", ""
	case ArgLen == 2:
		return args[1], ""
	case ArgLen >= 3:
		return args[1], args[2]
	}
	return "", ""
}

func usage() {
	fmt.Println("Usage: wget http://url.com/what/you/want /path/of/file")
}

func main() {
	url, filepath := parseArg(os.Args)
	if url == "" {
		os.Exit(1)
	}
	downloadFile(filepath, url)
}
