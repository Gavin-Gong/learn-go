package example

import (
	"fmt"
	"os"
)

func init() {
	f, _ := os.Create("./tmp/file.xx")
	fmt.Println(f)
	f.Close()
}
