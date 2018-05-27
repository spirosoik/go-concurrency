package main

import "fmt"

func main() {
	j := generator("John sends")
	d := generator("Doe sends")
	for i := 0; i < 100; i++ {
		fmt.Printf(<-j)
		fmt.Printf(<-d)
	}
	fmt.Printf("I'm done and quiting")
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
