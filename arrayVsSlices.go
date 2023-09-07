package main

import "fmt"

func main(){
	// Arrays
	var x [5]int // array
	fmt.Println(x)
	x[0] = 1
	fmt.Println(x)

	// Slices
	// type{elements} <- Composite Literal
	d := []int{1,2,3,4,5}
	fmt.Println(d)
	// Run though array with range
	for _ , v := range d {
		fmt.Println(v)
	}
	// Divide a slice
	newSlice := d[1:3]  // generate a new slice divided.
	fmt.Println("array divided: ",newSlice)
	// Add data
	p := []int{}
	k := []int{99,22,33}
	p = append(p,k...)
	fmt.Println(p)
	// Remove data
	deleted := append(k[:1],k[2:3]...)
	fmt.Println(deleted);
	// Using Make for build slice
	making := make([]int, 10, 11)
	fmt.Println("making slice", making)
	fmt.Println("long of making:", len(making))
	fmt.Println("capacity of making:", cap(making))
	making = append(making, 44);
	fmt.Println("making slice", making)
	fmt.Println("long of making:", len(making))
	fmt.Println("capacity of making:", cap(making))
	making = append(making, 45); // When a slice built with making reach its limit and add one element more, go create a new slice with the double of the capacity
	fmt.Println("making slice", making)
	fmt.Println("long of making:", len(making))
	fmt.Println("capacity of making:", cap(making))
}