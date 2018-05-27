package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := generator("John sends")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout: //entire conversation timeout
			fmt.Println("You're too slow.")
			return
		}
	}
}

func generator(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
