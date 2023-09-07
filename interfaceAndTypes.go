package main

import "fmt"

var a string = "hola"

var void string

func main(){
	//...interface{}
	testInterface("hola como andas", 12)
	testNumbers(1,2,3,4)

	const y string= "hola" // Long way

	var a string = "hola" // Long way
	a = "b";
	x := 1; // Short way

	testNumbers(x)
	testInterface(a)

	fmt.Println(void)
}

// The type interface{} is a generic type that it can recive in a function, this is equal of "any" on javascript.
// if precede of triple dot "..." it mean that the argument can recive N arguments of genereics type.
func testInterface(arg1 ...interface{}){
	fmt.Println(arg1...);
}

func testNumbers(numbers ...int){
	fmt.Println(numbers)
}
