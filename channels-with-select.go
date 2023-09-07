package main

import "fmt"

func main(){
	even := make(chan int)
	odd := make(chan int)
	exit := make(chan int)

	go send(even,odd,exit)
	get(even,odd,exit)

	fmt.Println("Finalizando.")
}

func send(e,o,ex chan<- int){
	for i:= 0; i < 100; i++{
		if i%2 == 0 {
			e <- i
		}else{
			o <- i
		}
	}
	close(ex)
}

func get(e,o,ex <-chan int){
	for {
		select {
		case v := <-e:
			fmt.Println("From even channel: ", v)
		case v := <-o:
			fmt.Println("From odd channel: ", v)
		case v, ok := <-ex:
			if !ok {
				fmt.Println("From comma ok channel: ", v, ok)
				return
			} else{
				fmt.Println("From comma ok channel: ", v)
			}
		return
		}
	}
}