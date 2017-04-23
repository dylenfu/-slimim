package main

import (
	"net/http"
	"fmt"
)

type mutex struct{}

const host = ":9090"

func HttpServerDemo()  {
	http.HandleFunc("hello", handleTest)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test only")
}