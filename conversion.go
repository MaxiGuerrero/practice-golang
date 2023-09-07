package main

import "fmt"

type dinero int

func main(){
	var number int

	var dinero dinero = 1

	number = int(dinero)

	fmt.Println(number)
}