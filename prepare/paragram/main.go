package main

//注意，同一个包内，调用其他文件的func，在编译的时候"go build ."而不是"go build"
//如果是直接运行main，则为"go run *.go"，而不是"go run main.go"
func main() {
	testCase := 2
	switch testCase {
	case 1:
		StructDemo()
		break
	case 2:
		InterfaceDemo()
		break
	case 3:
		HttpServerDemo()
		break
	case 4:
		BaseDemo()
		break
	default:
		break
	}
}
