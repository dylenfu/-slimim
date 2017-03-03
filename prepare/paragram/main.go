package main

import (
	. "fmt"
	. "slimim/prepare/paragram/demos"
)

func main() {
	Println(":::hello slim-im:::")

	testCase := 3
	switch testCase {
	case 1:
		StructDemo()
	case 2:
		InterfaceDemo()
	case 3:
		HttpServerDemo()
	}
}
