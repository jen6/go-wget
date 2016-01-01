package main

import (
	"bytes"
	"errors"
	"github.com/gosurui/uiprogress"
	"io"
)

type Progress struct {
	io.Reader
	Total int64 //전체 파일 크기 content-size
	Recv  int64 //현재까지 받은 파일 크기
}

func (ptr *Progress) Read(p []byte) (int, error) {
}

func (ptr *Progress) GetBar() uiprogress.Bar {
}
