package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func F41() {
	fn := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s, %s, %v, %v\n", path, info.Name(), info.Mode(), info.ModTime())
		return nil
	}
	_ = filepath.Walk(".", fn)
}
