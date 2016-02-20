package main

import (
	"fmt"
)

func Greet(name string, prefix string) {
	fmt.Println(prefix + "Hello, " + name)
}


func GreetNames(names []string, prefix string) {
	for _, n := range names {
		Greet(n, prefix)
	}
}

func main(){
	names := []string {
		"Dice",
		"Caplan",
		"Cornelius",
	}
	
	comm := make(chan string)
	
	go func () {
		GreetNames(names, "<GO>: ")
		comm <- "Finished greetings names"
	}()
	
	GreetNames(names, "<Main>: ")
	fmt.Println(<-comm)	
}