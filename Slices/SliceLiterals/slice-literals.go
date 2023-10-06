package main

import "fmt"

func main() {
	q := [6]int{2,3,5,7,9,11}
	fmt.Println(q)

	r := [6]bool{true, false, false, true, false, true}
	fmt.Println(r)

s := []struct{
	i int
	b bool
}{
	{2, true},
	{3, false},
	{3, false},
	{4, true},
	{5, false},
	{6, true},
}
	fmt.Println(s)
}