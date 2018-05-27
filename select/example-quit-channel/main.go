package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := generator("John sends", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("I'm exiting")
}

func generator(msg string, quit chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				fmt.Println("Quiting")
				return
			}
		}
	}()
	return c
}
