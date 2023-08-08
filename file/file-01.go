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
