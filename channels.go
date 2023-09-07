package main

import "fmt"

func main(){
	channel := make(chan int)

	go enviar(channel)
	recibir(channel)

	fmt.Println("Finalizando.")

}

func enviar(c chan<- int){
	c <- 42
}

func recibir(c <-chan int){
	fmt.Println(<-c)
}