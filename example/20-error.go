package example

import (
	"errors"
	. "fmt"
)

/*
 * 由于go 函数可以返回两个值, 所以一般约定第二个返回值为error, 如果error == nil 就说明没有错误
 */

func compute(n int) (int, error) {
	if n > 0 {
		return n, errors.New("define a error")
	}
	return n, nil

}
func er() {
	if _, e := compute(3); e != nil {
		Println("errors")
	} else {
		Println(e)
	}
}
