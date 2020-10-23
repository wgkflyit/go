package main

import (
	"fmt"
	"string/base"
	"string/file"
)

func main(){
	fmt.Println("hello world")
	p := Person1{}
	p.name="wuguangkkuo"
	p.Say()
	fmt.Println("end")
	MakeSlice()
	MakeMap()
	MakeChannel()
	NewPerson()
	ChageMap(TestMaps)
	fmt.Println(TestMaps)
	delete(TestMaps,"aaa")

	p1 := base.Person{}
	fmt.Println(p1)

	fmt.Println(TestMaps)

	file.TestFile()

}
var TestMaps =make(map[string]string)

func ChageMap(m map[string]string){
	fmt.Println(m)
	m["aaa"]="aaa"
}