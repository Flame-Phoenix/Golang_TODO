package main

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string
	age int
}

func (p *person) speak() {
	fmt.Println("Hello")
} 

type human interface {
	speak()
}

func saySomething(human human){
	human.speak()
}

func main() {

	person := person{
		firstname: "Nick",
		lastname: "Nickolaou",
		age: 2,
	}

	saySomething(&person)

}
