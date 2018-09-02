package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer func() {
		switch p := recover(); p {
		case nil:
		default:
			fmt.Println("catch:", p)
			defer printStack()
		}
	}()
	// f(3)
	num := noReturn()
	fmt.Println(num)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func noReturn() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = 1
		}
	}()
	panic("return")
}
