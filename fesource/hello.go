package main

import (
	"fmt"
	"syscall/js"
)

var killSignal chan bool

func init() {
	killSignal = make(chan bool)
}

func main() {
	fmt.Println("Hello, WASM !!")
	js.Global().Set("helloAgain", js.FuncOf(helloAgain))
	<-killSignal
}

func helloAgain(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("Hello again, WASM !!")
	return nil
}
