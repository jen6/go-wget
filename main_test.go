package main

import "testing"

var (
	DownloadLink = "http://makeall.ml/vagrant.sh"
	FilePath     = "/Users/songeon/Downloads"
)

func Test_parseFilename(t *testing.T) {
	t.Log("file name parse test")
	file := parseFilename(DownloadLink)
	if file != "vagrant.sh" {
		t.Error("Fail to parse")
	}
}

func Test_makePath(t *testing.T) {
	t.Logf("make path folder teset")
	path, err := makePath(FilePath, DownloadLink)
	if err != nil {
		t.Error("fail in makepath dir")
	}
	t.Log(path)

	t.Logf("make path file test")
	path, err = makePath(FilePath+"/vagrant", DownloadLink)
	if err != nil {
		t.Error("fail in makeapath file")
	}
	t.Log(path)

	t.Logf("make path error test")
	path, err = makePath("/wrong/path", DownloadLink)
	if err == nil {
		t.Error("fail in makepath invaild path")
	}
}

func Test_DownloadFile(t *testing.T) {
	t.Logf("try to download " + DownloadLink)
	err := downloadFile(FilePath, DownloadLink)
	if err != nil {
		t.Error("fail to download")
	}
}
