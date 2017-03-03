package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	. "slimim/prepare/serializ/protodata"
)

func main() {
	Proto2Demo()
	Proto3Demo1()
}

func Proto2Demo() {
	fmt.Println("=====proto2 demo=====")
	test := &Test{}
	test.Id = proto.Int32(1)
	test.Opt = proto.Int32(2)
	test.Str = proto.String("test")

	fmt.Println("test.Id", test.GetId())
	fmt.Println("test.Opt", test.GetOpt())
	fmt.Println("test.Str", test.GetStr())
}

func Proto3Demo1() {
	fmt.Println("=====proto3 demo1=====")
	test := &Test1{}
	test.Ed = Test1_X
	test.Page = 3
	test.Names = []string{"name1", "name2"}
	test.Users = map[string]int32{"name1":32, "name2":33}

	fmt.Println("test.Ed", test.Ed)
	fmt.Println("test.Page", test.Page)
	fmt.Println("test.Names", test.Names)
	fmt.Println("test.Users", test.Users)
}
