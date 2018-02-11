package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	in := bytes.NewBuffer([]byte("this is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input string"))
	in2 := bytes.NewBuffer([]byte("this is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input stringthis is a input string"))
	in3 := bytes.NewBuffer([]byte("a new input"))
	go func() {
		out := &bytes.Buffer{}
		io.Copy(out, in)
		fmt.Println("this one")
	}()
	go func() {
		out2 := &bytes.Buffer{}
		buf := make([]byte, 32)
		w := io.MultiWriter(out2, os.Stdout)
		io.CopyBuffer(w, in2, buf)
		fmt.Println("this two")
	}()
	time.Sleep(time.Second * 2)
	ww1 := &bytes.Buffer{}
	ww2 := &bytes.Buffer{}
	w := io.MultiWriter(ww1, ww2)
	io.Copy(w, in3)
	fmt.Println(ww1, ww2)

}
