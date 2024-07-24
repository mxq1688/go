package main

import (
	"fmt"
	"main/utils"
	"net/http"
)

type HelloHandlerStruct struct {
	content string
}

func (handler *HelloHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.content)
}

func main() {
	Say()
	// utils.print()
	utils.PrintText("调用子目录下的包")
	http.Handle("/", &HelloHandlerStruct{content: "Hello World！"})
	http.ListenAndServe(":8000", nil)
}
