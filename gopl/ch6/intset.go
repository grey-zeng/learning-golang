package main

import (
	"bytes"
	"fmt"
)

// 拆分成64每组的hash
type Intset struct {
	words []uint64
}

func (s *Intset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *Intset) Add(x int) {
	word, bit := x/64, uint(x%64)
	// 初始化直到有为止
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *Intset) UnionWith(t *Intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *Intset) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 练习题
func (s *Intset) Len() int {
	res := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res++
			}
		}
	}
	return res
}

func (s *Intset) Remove(x int) {
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] &= ^(1 << bit)
	}
}

func (s *Intset) Clear() {
	s.words = s.words[:0:0]
	// 不可以使用 s.words[:0] 容易导致内存泄漏
}

func (s *Intset) Copy() *Intset {
	var ns = Intset{make([]uint64, len(s.words))}
	copy(ns.words, s.words)
	return &ns
}

func main() {
	var x, y Intset
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	nx := x.Copy()
	x.Remove(9)
	fmt.Println("x", x.String())
	fmt.Println("nx", nx.String())

	y.Add(9)
	y.Add(42)
	fmt.Println("y", y.String())
}
