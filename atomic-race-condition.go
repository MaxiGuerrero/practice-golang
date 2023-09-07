package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){
	fmt.Println("Cantidad cpu:",runtime.NumCPU())
	fmt.Println("Numero de Gorutinas:",runtime.NumGoroutine())
	var counter int64

	const gs = 100;

	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i<gs; i++{
		go func(){
			atomic.AddInt64(&counter,1)
			runtime.Gosched()
			fmt.Println("count:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Cuenta:",counter)
}