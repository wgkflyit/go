package main

import "fmt"

//new struct
func NewPerson(){
	p := new(Person1)
	fmt.Println(p)
	fmt.Println(*p)
}