package example

/**
 * 接口是一系列函数签名的集合
 */

type IsMe struct{}
type Method interface {
	Code()
	Write() string
}

func (m IsMe) Code() {

}

func (m IsMe) Write() string {
	return "xx"
}

func init() {
	me := IsMe{}
	// me.
	me.Code()

	// 指定接口
	func(g Method) {
		g.Code()
	}(me)
}
