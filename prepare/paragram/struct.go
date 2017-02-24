package paragram

import (
	. "fmt"
)

func DemoStruct1() {
	demo := new(DemoStruct)
	demo.test1 = 1
	demo.test2 = "ss"
	Println("demo-struct-1", demo)
}

func DemoStruct2() {
	var demo *DemoStruct = new(DemoStruct)
	demo.test1 = 10
	demo.test2 = "sss"
	Println("demo-struct-2", demo)
}

func DemoStruct3() {
	demo := &DemoStruct{15, "sssss"}
	Println("demo-struct-3", demo)
}

type DemoStruct struct {
	test1 int
	test2 string
}
