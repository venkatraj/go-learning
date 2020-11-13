package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Channel1 Message"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Channel2 Message"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received1: ", msg1)
		case msg2 := <-c2:
			fmt.Println("Received2:", msg2)
		}
	}

}
