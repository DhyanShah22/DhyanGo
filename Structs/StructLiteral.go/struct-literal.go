package main

import "fmt" 

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{3, 5}
	v3 = Vertex{}
	p = &Vertex{5,7}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}