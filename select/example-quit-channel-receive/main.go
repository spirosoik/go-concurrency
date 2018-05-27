package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	c := generator("John sends", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	fmt.Printf("John says %q\n", <-quit)
}

func generator(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				cleanup()
				quit <- "See you"
				return
			}
		}
	}()
	return c
}

func cleanup() {

}
