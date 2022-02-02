package main

import (
	"fmt"
	"time"
)

func printNumber2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	w <- true
}

func useChannel() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumber2(w1)
	go printLetters2(w2)
	<-w1
	<-w2
}

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >> ", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught << ", num)
	}
}

func sendRecMessage() {
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(1000 * time.Microsecond)
}

func main() {
	// useChannel()
	sendRecMessage()
}
