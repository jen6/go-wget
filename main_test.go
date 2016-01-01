package main

import "testing"

var (
	DownloadLink = "http://makeall.ml/vagrant.sh"
	FilePath     = "./vagrant.sh"
)

func TestDownloadFile(t *testing.T) {
	t.Logf("try to download" + DownloadLink)
	err := downloadFile(FilePath, DownloadLink)
	if err != nil {
		t.Error("fail to download")
	}
}
