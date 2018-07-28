package main

import (
	"io"
	"os"
	"strings"
)

func rot13(x byte) byte {
    switch {
    case x >= 65 && x <= 77:
        fallthrough
    case x >= 97 && x <= 109:
        x = x + 13
    case x >= 78 && x <= 90:
        fallthrough
    case x >= 110 && x <= 122:
        x = x - 13
    }
    return x
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i:=0; i<n; i++ {
		if b[i] != ' ' {
			b[i] = rot13(b[i])	
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

