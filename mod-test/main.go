package main

import (
	"fmt"
	"log"
	"mymodule/mypackage"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello!")

	hellos, err := mypackage.Hellos([]string{"xfy", "dfy"})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range hellos {
		fmt.Println(v)
	}
}
