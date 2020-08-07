package main

type animal interface {
	Shut()
}

type Cat struct {

}

func(c * Cat) Shut() {

}

type Dog struct {

}

func(d * Dog) Shut() {

}

func main() {
	// interface是一种类型

}
