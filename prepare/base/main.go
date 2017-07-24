package main

import (
	"flag"
	. "slimim/prepare/base/abc"
	"slimim/prepare/base/bcd"
	_ "slimim/prepare/base/bcd"
	"strconv"
)

var (
	testCase = flag.String("tc", "loop", "choose testcase")
)

func main() {
	flag.Parse()

	switch *testCase {
	case "loop":
		dLoop()
	case "package":
		dPackage()
	}
}

func dLoop() {
	for i := 0; i < 10; i++ {
		println("line " + strconv.Itoa(i))
	}

	t := []int{1, 2, 3}
	for k, i := range t {
		println(k, i)
	}
}

func dPackage() {
	Ipublish()
	println(bcd.B_POP)
}
