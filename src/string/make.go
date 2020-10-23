package main


/**
make 内置函数的三个初始化例子，slice ，map，chan
 */
import (
	"fmt"
	"strconv"
	"time"
)

func MakeSlice(){
	pa := make([]int,4,8)
	pa[1]=10

	for i:=1;i<8;i++{
		pa=append(pa,i)
	}

	fmt.Println(pa)
}


func MakeMap(){
	m := make(map[string]string)
	m["aaa"]="aaaa"
	for i:=0;i<10;i++{
		index :=strconv.Itoa(i)
		m[index]="aa"
	}

	fmt.Println(m)
	fmt.Println(m["3"])

}

func MakeChannel(){
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second)
}

 func consumer(ch chan int){
		a :=<-ch
		fmt.Println(a)
	}

func producer(ch chan int){
		ch<- 999
}


