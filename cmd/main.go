package main

import (
	"log"
	"tabl"
)

func main() {
	err := tabl.CreateTabl("foo")
	if err != nil {
		log.Fatal(err)
	}
}
