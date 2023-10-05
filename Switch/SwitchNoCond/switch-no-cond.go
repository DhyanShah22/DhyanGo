package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("Good Morning!")
	case t.Hour() < 17 :
		fmt.Println("Good Afternoon!")
	case t.Hour() < 21 :
		fmt.Println("You are the best coder Dhyan.")
	default: 
		fmt.Println("Good Evening!!")
	}
}