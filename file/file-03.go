package file

import (
	"fmt"
	"os"
)

func F31() {
	fi, err := os.Lstat("main.go")
	if err != nil {
		fmt.Println("error obtaining file/directory properties:", err)
		return
	}
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())
	fmt.Println(fi.Mode())
	fmt.Println(os.ModeSymlink)
	fmt.Println(fi.Mode()&os.ModeSymlink != 0)
}

func F32() {
	fis, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("error read directory:", err)
		return
	}
	for _, fileInfo := range fis {
		fmt.Println(fileInfo.Name())
		fmt.Println(fileInfo.IsDir())
		fmt.Println(!fileInfo.IsDir())
	}
}
