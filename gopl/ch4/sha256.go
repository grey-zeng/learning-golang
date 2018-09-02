package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, input := range os.Args[1:] {
			hashVal := sha256.Sum256([]byte(input))
			fmt.Printf("%s the hash256 val is %x\n", input, hashVal)
		}
		return
	}
	fmt.Println("vim-go")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	diff(c1, c2)
}

func diff(c1, c2 [32]uint8) {
	m := make(map[uint8]int)
	for i, a1 := range c1 {
		a2 := c2[i]
		if a1 != a2 {
			m[a2]++
		}
	}
	for key, value := range m {
		fmt.Printf("%x\t=>%d\n", key, value)
	}

}
