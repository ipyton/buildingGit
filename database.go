package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func writeObject(){
	var in bytes.Buffer
	writer := zlib.NewWriter(&in)
	writer.Write([]byte("shiofh"))
	writer.Close()

	var out bytes.Buffer
	r, _:=zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(out.Bytes())

}

func readObjects() {

}
