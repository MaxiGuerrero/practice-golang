package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main(){
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUS\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	go foo()
	go bar()
	wg.Add(runtime.NumGoroutine()-1)
	wg.Wait()
}

func foo(){
	for i := 0; i < 10; i++{
		fmt.Println("foo:",i)
	}
	wg.Done()
}

func bar(){
	for i := 0; i < 10; i++{
		fmt.Println("bar:",i)
	}
	wg.Done()
}