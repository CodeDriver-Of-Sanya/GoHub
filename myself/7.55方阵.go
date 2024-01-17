package main

import (
	"fmt"
)

func main() {

	for j := 1; j <= 5; j++ {
		for i := 1; i <= 5; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
