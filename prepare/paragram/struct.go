package paragram

import (
	. "fmt"
)

func StructDemo() {
	DemoStructFunc1()
	DemoStructFunc2()
	DemoStructFunc3()
}

func DemoStructFunc1() {
	demo := new(DemoStruct)
	demo.test1 = 1
	demo.test2 = "ss"
	Println("demo-struct-1", demo)
}

func DemoStructFunc2() {
	var demo *DemoStruct = new(DemoStruct)
	demo.test1 = 10
	demo.test2 = "sss"
	Println("demo-struct-2", demo)
}

func DemoStructFunc3() {
	demo := &DemoStruct{15, "sssss"}
	Println("demo-struct-3", demo)
}

type DemoStruct struct {
	test1 int
	test2 string
}
