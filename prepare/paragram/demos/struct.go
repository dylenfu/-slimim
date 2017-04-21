package demos

func StructDemo() {
	DemoStructFunc1()
	DemoStructFunc2()
	DemoStructFunc3()
	DemoStructFunc4()
	DemoStructFunc5()
}

func DemoStructFunc1() {
	demo := new(DemoStruct)
	demo.test1 = 1
	demo.test2 = "ss"
	println("demo-struct-1", demo)
}

func DemoStructFunc2() {
	var demo *DemoStruct = new(DemoStruct)
	demo.test1 = 10
	demo.test2 = "sss"
	println("demo-struct-2", demo)
}

func DemoStructFunc3() {
	demo := &DemoStruct{15, "sssss"}
	println("demo-struct-3", demo)
}

func DemoStructFunc4() {
	demos := []DemoStruct{
		DemoStruct{11,"11"},
		DemoStruct{12,"12"}}
	println("demo-struct-4",len(demos))
}

func DemoStructFunc5() {
	demo := DemoStruct{1,"1"}
	println("demo-struct-5", demo.test1)
}

type DemoStruct struct {
	test1 int
	test2 string
}
