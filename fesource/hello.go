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
	var ret string
	if len(inputs) == 0 {
		ret = "no one"
	} else {
		for _, inp := range inputs {
			ret += "," + inp.String()
		}
		ret = ret[1:]
	}

	d := js.Global().Get("document")
	elem := d.Call("getElementById", "replyInp")
	elem.Set("value", ret)

	elem = d.Call("getElementById", "replyDiv")
	elem.Set("innerHTML", "Hello <b>"+ret+"</b>")

	b := elem.Call("querySelector", "b")

	return b.Get("innerText")
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
