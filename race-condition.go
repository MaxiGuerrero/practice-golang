package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	fmt.Println("Cantidad cpu:",runtime.NumCPU())
	fmt.Println("Numero de Gorutinas:",runtime.NumGoroutine())
	counter := 0

	const gs = 100;

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(gs)

	for i := 0; i<gs; i++{
		go func(){
			mutex.Lock()
			v:= counter
			v++
			runtime.Gosched()
			counter=v
			wg.Done()
			mutex.Unlock()
		}()
		fmt.Println("Numero de Gorutinas:",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:",counter)
}