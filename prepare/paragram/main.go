package main

import (
	. "slimim/prepare/paragram/demos"
)

func main() {
	testCase := 1
	switch testCase {
	case 1:
		StructDemo()
	case 2:
		InterfaceDemo()
	case 3:
		HttpServerDemo()
	case 4:
		BaseDemo()
	}
}
