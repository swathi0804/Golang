package main

import (
	"fmt"
	"time"
)

func pinger(pinger <- chan int, ponger chan<- int) {  //pinger prints a
	for {                         // ping and waits for a pong
		<- pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) { //ponger prints
	for {                         // a pong and waits for a ping
		<- ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)   //makes the ping channel 
	pong := make(chan int)  // makes the pong channel 

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1 // starts the ping/pon by sending to ping channel

	// for{
	// 	time.Sleep(time.Second) //block the main thread until an interrupt
	// }

	select{}  // instead of using for we can use select also 
}