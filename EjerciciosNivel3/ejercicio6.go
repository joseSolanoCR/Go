package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for a := 1; a < 10; a++ {
		x := rand.Intn(10)
		if x%a == 0 {
			fmt.Println("ganador")
		}
	}
}
