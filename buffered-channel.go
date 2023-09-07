package main

import "fmt"

func main(){
	ca := make(chan int,1) // buffered channel

	ca <- 42

	fmt.Println(<-ca)

}