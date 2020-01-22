package main

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"
)

//GOOS=js GOARCH=wasm go build -o ../frontend/benchmark.wasm

func benchmark(f func()) func(this js.Value, inputs []js.Value) interface{} {
	return func(this js.Value, inputs []js.Value) interface{} {
		t := time.Now()
		f()
		fmt.Println("go:", time.Now().Sub(t))
		return nil
	}
}

func runLoop1000000() {
	for i := 0; i < 1000000; i++ {

	}
}

func concat10000chars() {
	var s string
	for i := 0; i < 100000; i++ {
		s += "s"
	}
}

func concat10000chars2() {
	var builder strings.Builder

	for i := 0; i < 100000; i++ {
		builder.WriteString("s")
	}
}

func push10000() {
	var mp []int
	for i := 0; i < 100000; i++ {
		mp = append(mp, i)

	}
}

var killSignal chan bool

func init() {
	killSignal = make(chan bool)
}

func main() {
	href := js.Global().Get("location").Get("href")
	js.Global().Set("runLoop1000000Go", js.FuncOf(benchmark(runLoop1000000)))
	js.Global().Set("concat10000charsGo", js.FuncOf(benchmark(concat10000chars)))
	js.Global().Set("concat10000charsGo2", js.FuncOf(benchmark(concat10000chars2)))
	js.Global().Set("push10000Go", js.FuncOf(benchmark(push10000)))

	println(href.String())

	<-killSignal
}
