package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("example_dir"); err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("aFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	value := []byte("this is the content")
	f.Write(value)
	f.Close()

	f, _ = os.Open("aFile.txt")
	io.Copy(os.Stdout, f)
	f.Close()
	os.Chdir("..")
	if err = os.RemoveAll("example_dir"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getwd())

	f1, _ := os.Create("file1.txt")
	f1.Write([]byte("this is me222222222222222"))
	f2, _ := os.Create("file2.txt")
	capFile(f1, f2)
	f1.Close()
	f2.Close()
}

func capFile(f1 *os.File, f2 *os.File) {
	f1.Seek(0, 0)
	fmt.Println()
	tmp := new(bytes.Buffer)
	//tmp := &bytes.Buffer{}
	io.Copy(tmp, f1)
	fmt.Println(tmp)
	s := strings.ToUpper(tmp.String())
	fmt.Println(s)
	io.Copy(f2, strings.NewReader(s))
}
