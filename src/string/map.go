package main

import "fmt"

func Add(){
	m :=make(map[int]int)
	m[0]=2
	m[1]=2
	m[2]=3
	m[3]=4
	delete(m,2)
	fmt.Println(m)
}