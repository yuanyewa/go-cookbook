package main

import (
	"fmt"
	"io"
	"os"
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

}
