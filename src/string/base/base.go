package base



type IEat interface{
	eat()
}


type IWalk interface {
	walk()
}

type Person struct{
	Name string
	Age int
	sex bool
	School string
	IDCard string
}