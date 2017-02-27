package main

import (
	. "fmt"
	"slimim/prepare/paragram"
)

func main() {
	Println(":::hello slim-im:::")

	testCase := 2
	switch testCase {
	case 1:
		paragram.StructDemo()
	case 2:
		paragram.InterfaceDemo()
	}
}
