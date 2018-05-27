package main

import "fmt"

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
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d \n", msg, i)
		}
	}()
	return c
}
