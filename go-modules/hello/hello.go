package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Render once times hello
	//message, err := greetings.Hello("Khoa")

	//Render multiple hello
	names := []string{"Khoa", "Teo", "Tu"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
