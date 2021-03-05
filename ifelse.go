package main

import "fmt"

func main(){
    var num = 20
    if num <= 2 {
        fmt.Println(num, "is the smallest num")
    } else if num < 12 {
        fmt.Println(num, "is the medium num")
    } else {
        fmt.Println(num, "is the largest num")
	}
}