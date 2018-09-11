package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 直接照搬了源码的设计
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

type LimitedReader struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {
	// var str string
	// fmt.Printf("请输入内容：")
	// fmt.Scanf("%s", &str)
	// fmt.Printf("输入了：%s", str)
	r := strings.NewReader("Hello World!")
	lr := io.LimitReader(r, 5)

	n, err := io.Copy(os.Stdout, lr)  // Hello
	fmt.Printf("\n%d   %v\n", n, err) // 5   <nil>
}
