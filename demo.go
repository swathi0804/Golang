package main

import "fmt"

func owner(name string, age int) int  {
	a := name
	b := age
	return b
}
func main(){
	var result = owner{"san", 24}
	fmt.Println(result)
}

