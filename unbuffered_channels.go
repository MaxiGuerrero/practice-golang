package main

import "fmt"

func main(){
	ca := make(chan int) // Unbuffered channel

	go func(){
		ca <- 42
	}()

	fmt.Println(<-ca)

}