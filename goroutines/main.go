package main

import (
	"fmt"
	"time"
)

func main() {
	// a channel allows main and goroutine communicate
	c := make(chan bool)
	people := [2]string{"fernok", "Lyndis"}
	for _, person := range people {
		go isTired(person, c)
	}
	// when you receive something from a channel,
	// the main function waits until you get a reply
	// "receiving a message" is a blocking operation
	// compiler does nothing beyond getting message from a channel
	// until a message is actually received
	resultOne := <-c
	resultTwo := <-c
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
	// fmt.Println(<-c)
	/*
		loops can be used to get multiple messages from channel
		for i := 0; i < len(people); i++ {
			fmt.Println(<-c)
		}
	*/
	go countNumber("fernok")
	countNumber("Lyndis")
	// go routine is alive as long as the main function is running
	// the main function does not wait for go routine
}

func countNumber(gostring string) {
	for i := 0; i < 10; i++ {
		fmt.Println(gostring, "goes!", i)
		time.Sleep(time.Second)
	}
}

func isTired(person string, c chan bool) {
	fmt.Println(person)
	time.Sleep(time.Second * 5)
	c <- true
}
