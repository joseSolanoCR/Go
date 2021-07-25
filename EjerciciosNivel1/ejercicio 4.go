package main

import (
	"fmt"
)

type mi_tipo int

var x mi_tipo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
