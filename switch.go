package main

import "fmt"

func main() {
	num := 1
	switch num {
		case 1: 
			fmt.Println("one")
	    case 2:
			fmt.Println("two")
		default: 
			fmt.Println("none")
	}
}