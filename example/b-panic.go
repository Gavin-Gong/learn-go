package example

import (
	"fmt"
	"os"
)

func pnc() {
	f, _ := os.Create("./tmp/file.xx")
	fmt.Println(f)
	f.Close()
}
