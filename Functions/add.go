package main

import "fmt"

func add(x int, y int) int {                    // x,y int is also used if same datatype
	return x+y
}

func main() {
	fmt.Println(add(42,46))
}