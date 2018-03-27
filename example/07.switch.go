package example

import (
	"fmt"
)

func init() {
	// case: 1
	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("3/4")
	default:
		fmt.Println("not  1 2 3 4")
	}

	// 类似if语句
	m := 3
	switch {
	case m > 2:
		fmt.Println(m)
	case m < 4:
		fmt.Println(m)
	}

	// 类型switch

}

func Type(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unkown type")
	}
}
