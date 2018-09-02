package main

import (
	"fmt"
	"github.com/learning-golang/gopl/ch2/popcount"
)

func main() {
	fmt.Println("I am main")
	res := popcount.PopCount(20)
	fmt.Println(res)
}
