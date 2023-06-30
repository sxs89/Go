package main

import "fmt"

func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func main() {
	ch := make(chan string, 1)
	send(ch, "hello world")
	read(ch)

}
