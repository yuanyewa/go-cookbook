package main

import (
	"fmt"
	"flag"
)

func main(){
	var isAwesome bool
	var val int
	flag.BoolVar(&isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&val, "v",100, "enter a value")
	flag.Parse()
	fmt.Println(val)
}