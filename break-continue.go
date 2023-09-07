package main

import "fmt"

func main(){
	for i := 0; i < 500; i++ {
		if i>100 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}