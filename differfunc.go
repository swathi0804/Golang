package main

import "fmt"

func a() {
	fmt.Println("I am A")
	defer b()
	fmt.Println("A ends...")
}
func b() {
	fmt.Println("I am in B")
}
func main() {
	fmt.Println("main begins ")
	//defer a()
	// defer fmt.Println("I am in main")
	a()
	fmt.Println("main ends")
}