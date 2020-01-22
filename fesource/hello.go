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
	for _, inp := range inputs {
		fmt.Println("Hello again, ", inp.String())
	}
	return nil
}
