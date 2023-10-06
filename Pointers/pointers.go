package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i //point to i
	fmt.Printf("The value of i is %v\n", *p) // printing value of i
	*p = 21 // changing value of i via pointer
	fmt.Printf("The value of i is %v\n", i)

	q := &j // point to j 
	fmt.Printf("The value of j is %v\n", *q) // printing value of j
	*q = 43 // changing value pf j via pointer
	fmt.Printf("The new value of j is %v\n", j)
}