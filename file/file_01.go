package file

import (
	"fmt"
	"os"
)

func F1() {
	f, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, _ = f.Write([]byte("Hello World"))

	_, _ = f.WriteString("\nHello World")
}

func F2() {
	f, err := os.OpenFile("demo.txt", os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	bs := make([]byte, 10)
	n, err := f.Read(bs)
	fmt.Printf("read content: %s, err: %v\n", string(bs[:n]), err)

	n, err = f.WriteString("abc123")
	fmt.Printf("written bytes number: %d, err: %v\n", n, err)
}
