package example

import (
	. "fmt"
)

func init() {
	//
	x, y := 1, 2
	a, b := retval(x, y)
	Println(a, b)
}

func retval(x, y int) (int, int) {
	return y, x
}
