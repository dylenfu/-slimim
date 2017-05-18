package main

import (
	. "slimim/prepare/paragram/demos"
)

func main() {
	testCase := 1
	switch testCase {
	case 1:
		demos.StructDemo()
	case 2:
		demos.InterfaceDemo()
	case 3:
		HttpServerDemo()
	case 4:
		BaseDemo()
	}
}
