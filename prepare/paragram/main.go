package main

//注意，同一个包内，调用其他文件的func，在编译的时候"go build ."而不是"go build"
//如果是直接运行main，则为"go run *.go"，而不是"go run main.go"
//这里break其实没什么用
func main() {
	testCase := 3
	switch testCase {
	case 1:
		demos.StructDemo()
	case 2:
		demos.InterfaceDemo()
	case 3:
		HttpServerDemo()
	case 4:
		BaseDemo()
	default:
		break
	}
}
