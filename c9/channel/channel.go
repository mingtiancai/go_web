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

func callerA(c chan string) {
	c <- "Hello World!"
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
}

func useSelect() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	for i := 0; i < 5; i++ {
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("Default")
		}
	}
}

func callerA2(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB2(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func useSelect2() {
	a, b := make(chan string), make(chan string)
	go callerA2(a)
	go callerB2(b)
	ok1, ok2 := true, true
	var msg string

	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}

func main() {
	// useChannel()
	// sendRecMessage()
	// useSelect()
	useSelect2()
}
