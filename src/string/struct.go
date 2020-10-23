package main

import "fmt"

func init(){
	fmt.Println("init doing")
}

type Person1 struct{
	name string
	age int
	sex bool
	school string
}


type Man interface{
	say()
	Eat()
}

func (p Person1)Say(){
	fmt.Println("person say")
}

func (p Person1)Eat(){
	fmt.Println("person eat")
}