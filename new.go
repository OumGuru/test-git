package main

import (
	"fmt"
	"/OFFICE(old)/prime/dog/main.go"
)

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}

type canine struct {
	name string
	age  int
}
