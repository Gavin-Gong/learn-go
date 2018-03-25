package example

import(
	"fmt"
)

func init()  {
	arr := [...]int {1, 2, 3, 4, 5, 6}
	// range loop
	for i, v := range arr {
		fmt.Println(i, v)
	}

	a := 2
	for (a < 9) {
		a++
		fmt.Println(a)
	}

	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}

	// infinite loop
	for {
		fmt.Println("infinite loop")
		break // break loop
	}

	// coutine loop
	for i := 0; i < 9; i++ {
		if i% 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}