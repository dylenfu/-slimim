package main

import (
	"fmt"
	"slimim/prepare/paragram/demos"
)

func main() {
	fmt.Println(":::hello slim-im:::")
	testCase := 3
	switch testCase {
	case 1:
		demos.StructDemo()
	case 2:
		demos.InterfaceDemo()
	case 3:
		demos.HttpServerDemo()
	}
}
