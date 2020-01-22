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
		doc := js.Global().Get("document")
		btn := doc.Call("createElement", "button")
		btn.Set("innerHTML", inp.String())

		a := js.Global().Get("buttonToInput")
		btn.Call("addEventListener", "click", a)

		doc.Get("body").Call("appendChild", btn)
	}
	return nil
}

func printValue(name string, v js.Value) {
	var ret string
	if v == js.Undefined() {
		ret = " is undefined"
	} else if v == js.Null() {
		ret = " is null"
	} else if v.Type() == js.TypeBoolean {
		ret = " bool value = " + fmt.Sprint(v.Bool())
	} else if v.Type() == js.TypeNumber {
		ret = " num value = " + fmt.Sprint(v.Float())
	} else if v.Type() == js.TypeString {
		ret = " str value = " + v.String()
	} else if v.Type() == js.TypeFunction {
		ret = " is a func "
	} else if v.Type() == js.TypeObject {
		ret = " is an object "
	} else {
		ret = " is TBD "
	}
	fmt.Println(name + ret)
}
