package file

import (
	"fmt"
	"os"
)

func F21() {
	err := os.WriteFile("demo.txt", []byte("Hello World"), 0666)
	if err != nil {
		fmt.Println("write file error:", err)
	}

	bs, err := os.ReadFile("demo.txt")
	if err != nil {
		fmt.Println("read file error:", err)
	} else {
		fmt.Println("read content:", string(bs))
	}
}

func F22() {
	file := "demo.txt"
	fmt.Println(IsExist(file))
	fmt.Println(IsFile(file))
	fmt.Println(IsDir(file))
	fmt.Println(IsSymlink(file))
}

func IsExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func IsFile(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && !fi.IsDir()
}

func IsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

func IsSymlink(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.Mode()&os.ModeSymlink != 0
}
