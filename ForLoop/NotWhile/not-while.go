package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum+= 1
	}
	fmt.Println(sum)
}