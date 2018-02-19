package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is to test go Pipe")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		myS := scanner.Text()
		fmt.Println(myS)
	}
}
