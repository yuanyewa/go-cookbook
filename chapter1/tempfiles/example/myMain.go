package main

import (
	"io/ioutil"
	"fmt"
	"os"
)
func main() {
	t, _ := ioutil.TempDir("", "tmp")
	fmt.Println(os.Getwd())
	tf, _ := ioutil.TempFile(t, "tmp")
	fmt.Println(tf.Name)
	os.RemoveAll(t)
}