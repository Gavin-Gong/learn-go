package example

import (
	"fmt"
	"io/ioutil"
)

func init() {
	data, _ := ioutil.ReadDir("./")
	// fmt.Println(data, err)
	for _, v := range data {
		fmt.Println(v.Name())
	}
	fmt.Println("")
}
