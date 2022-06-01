package main

import (
	"bytes"
	"io"
)

type MyBytesStr struct {
	Storage []byte
}

type MyBytesInterFace interface {
	Read(value byte) byte
	Write(value byte) byte
}

func (my *MyBytesInterFace) Read(value byte) byte {

	io.WriteString(value)
	io.ReadAll(value)
}

func main() {

	bytes.Buffer{}
}
