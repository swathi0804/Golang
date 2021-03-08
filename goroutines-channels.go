package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)    //making the 1st channel
	c2 := make(chan string)    // making the 2nd channel

	go func() {                // go key word denotes goroutine
		for {
			c1 <- "every 500ms"    // sends msg to the channel
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every two second"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// fmt.Println(<-c1)  ---it will run and wait for c2, still c2 takes 2 sec to execute 
		// fmt.Println(<-c2)   
		select {                    // to avoid the above method we are using select 
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
