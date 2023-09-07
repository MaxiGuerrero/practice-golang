package main

import (
	"fmt"
	"sync"
)

func main(){
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even,odd)
	go get(even,odd,fanin)

	for v:= range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}

func send(e,o chan<- int){
	for i:= 0; i < 100; i++{
		if i%2 == 0 {
			e <- i
		}else{
			o <- i
		}
	}
	close(e)
	close(o)
}

func get(e,o <-chan int,fanin chan<- int){
	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		for v:= range e{
			fanin <- v
		}
		wg.Done()
	}()

	go func(){
		for v:= range o{
			fanin <- v
		}
		wg.Done()
	}()
	
	wg.Wait()
	close(fanin)
}