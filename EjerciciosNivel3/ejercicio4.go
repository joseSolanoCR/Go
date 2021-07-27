package main

import (
	"fmt"
)

func main() {
	nac := 1979
	for {
		if nac > 2021 {
			break
		}
		fmt.Println(nac)
		nac++
	}
}
