package main

import "fmt"

type Action interface {
	salute()
}

type Person struct {
	message string
}

func (p Person) salute() {
	fmt.Println(p.message)
}

func main(){
	var maxi Action = Person{message: "hola papus"};
	maxi.salute();
	
}