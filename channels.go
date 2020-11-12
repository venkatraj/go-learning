package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		// channels are un buffered
		// meanking that if we don't read again after reading `ping`
		// then `pong will not be sent`

		messages <- "pong"
	}()

	msg := <-messages
	fmt.Println(msg) // ping

	msg = <-messages
	fmt.Println(msg) // pong

}
