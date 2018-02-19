package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)

	err = fmt.Errorf("this is an error %s", "oh something is wrong")
	fmt.Println(err)
}
