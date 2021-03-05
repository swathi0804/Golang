package main

import (
	"fmt"
	"math"
)

func main(){
	var num float64 = 24
	var intresult = math.Sqrt(num)
	var result = math.Ceil(intresult)  //up o/p 5
	// var result = math.Floor(intresult)  // down o/p 4
	// var result = math.Round(intresult)  // o/p r
	fmt.Println("the output is %f ", result)
}



