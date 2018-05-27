package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(generator("John sends"), generator("Doe sends"))
	for i := 0; i < 100; i++ {
		fmt.Printf(<-c)
	}
	fmt.Printf("I'm done and quiting")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func generator(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d \n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
