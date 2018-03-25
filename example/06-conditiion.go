package example

import(
	"fmt"
)

func init() {
	if num := 1; num > 0 {
		fmt.Println(">0", num)
	} else if num < -9 {
		fmt.Println(num)
	} else {

	}
}