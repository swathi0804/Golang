package main

import (
	"fmt"
	"net/http"
	"log"
)
func main(){
    request1()
}
func homePage(response http.ResponseWriter, r *http.Request){
	fmt.Fprintf(response, "welcome to web api home page")
	fmt.Println("Endpoint: home page")
}
func aboutMe(response http.ResponseWriter, r *http.Request){
	fmt.Fprintf(response, "Hai everyone, I'm swathi from kloudone")
	fmt.Println("Endpoint: home page")
}
func request1(){
    http.HandleFunc("/",homePage)
	http.HandleFunc("/",aboutMe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
