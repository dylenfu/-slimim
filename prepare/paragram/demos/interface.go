package demos

import (
	"fmt"
)

func InterfaceDemo() {
	struct1 := &DemoInterfaceStruct1{1, "111"}
	struct2 := &DemoInterfaceStruct2{"222", 3.05}
	struct1.DemoInterfacePrintInt()
	struct1.DemoInterfacePrintString()
	struct2.DemoInterfacePrintString()
	struct2.DemoInterfacePrintFloat()
}

func (s *DemoInterfaceStruct1) DemoInterfacePrintInt() {
	fmt.Println("print int", s.data1)
}

func (s *DemoInterfaceStruct1) DemoInterfacePrintString() {
	fmt.Println("print string", s.data2)
}

func (s *DemoInterfaceStruct2) DemoInterfacePrintString() {
	fmt.Println("print string", s.data1)
}

func (s *DemoInterfaceStruct2) DemoInterfacePrintFloat() {
	fmt.Println("print float", s.data2)
}

type DemoInterfaceStruct1 struct {
	data1 int
	data2 string
}

type DemoInterfaceStruct2 struct {
	data1 string
	data2 float32
}

type DemoInterface1 interface {
	DemoInterfacePrintInt()
	DemoInterfacePrintString()
}

type DemoInterface2 interface {
	DemoInterfacePrintString()
	DemoInterfacePrintFloat()
}
