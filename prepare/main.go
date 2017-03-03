package main

import (
	. "fmt"
	"slimim/prepare/paragram"
	"slimim/prepare/serializ"
)

func main() {
	Println(":::hello slim-im:::")

	testCase := 3
	switch testCase {
	case 1:
		paragram.StructDemo()
	case 2:
		paragram.InterfaceDemo()
	case 3:
		serializ.ProtoDemo()
	case 4:
		paragram.HttpServerDemo()
	}
}
