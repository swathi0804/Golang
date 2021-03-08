package main

import "fmt"

type User struct {
	name string
	rollno int
	mark int
	place string
}
func main(){
	var struct1 = User{"swathi", 1, 94, "walaja"}
	// fmt.Println(struct1)--getting all the student data
	//fmt.Println(struct1.place, struct1.name)
	fmt.Println(struct1)

	var struct2 = User{name : "san", place : "bangalore"} //declaring value for name and place alone 
	fmt.Println(struct2)
}