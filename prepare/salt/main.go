package main

import (
	"os"
)

func main() {
	testCase := 1
	switch testCase {
	case 1: osArgsTest()
	}
}

// go run main.go -t1 "test1" -t2 "test2"
// result:
// os.Args[0] /tmp/go-build382341349/command-line-arguments/_obj/exe/main
// os.Args[1] -t1
// os.Args[2] test1
func osArgsTest() {
	println("os.Args[0]", string(os.Args[0]))
	println("os.Args[1]", string(os.Args[1]))
	println("os.Args[2]", string(os.Args[2]))
}
