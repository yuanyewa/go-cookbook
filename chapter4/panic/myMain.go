package main

import (
	"fmt"
)

func Panic() {
	z := 0
	a := 1 / z
	fmt.Println("the value of a is ", a)
}

func Catcher() {
	defer func() {
		fmt.Println("I'm panic, ", recover())
	}()
}
func main() {
	Catcher()
}
