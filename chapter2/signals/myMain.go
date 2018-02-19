package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func CatchSig(ch chan os.Signal, done chan bool) {
	sig := <-ch
	fmt.Println("signal received: ", sig)
	switch sig {
	case syscall.SIGINT:
		fmt.Println("receives a SIGINT")
	case syscall.SIGTERM:
		fmt.Println("receives SIGTERM")
	default:
		fmt.Println("receives something")
	}
	done <- true
}
func main() {
	fmt.Println(os.Getwd())

	signals := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go CatchSig(signals, done)
	<-done
	fmt.Println("Done")

}
