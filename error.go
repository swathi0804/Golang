package main

import (
    "fmt"
    "os"

)
func file(){
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
}

func fileopen(){
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
}

func main(){
    file()
	fileopen()
}
