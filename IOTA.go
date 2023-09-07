package main

import "fmt"

// IOTA is a count of integer constants without tyhpe, that start at 0.
const (
	a = iota
	b
	c
)

const (
	d = iota
	f
	g
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)
}