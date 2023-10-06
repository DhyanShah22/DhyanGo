package main 

import "fmt"

func main() {
	names:= [5]string {
		"Dhyan",
		"Shubh",
		"Rig",
		"Moxil",
		"Priyansh",
	}

	fmt.Println(names)

	a := names[1:3]
	b := names[2:5]
	fmt.Println(a, b)

	b[0] = "Dhyan ka beta"
	fmt.Println(a,b)
	fmt.Println(names)
}