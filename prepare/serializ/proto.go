package serializ

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func ProtoDemo() {
	test := new(Test)
	test.Id = proto.Int32(1)
	test.Opt = proto.Int32(2)
	test.Str = proto.String("test")

	fmt.Println("test.Id", test.GetId())
	fmt.Println("test.Opt", test.GetOpt())
	fmt.Println("test.Str", test.GetStr())
}
