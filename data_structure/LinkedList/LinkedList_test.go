package LinkedList

import "testing"

func Test_InsertFirst(t *testing.T)  {
	list := LinkedList{}
	list.InsertFirst(1)
	list.InsertFirst(2)
	
	if list.head.data != 2 {
		t.Errorf("except 2 but get %s", list.head.data)
	}
	if list.head.next.data != 1 {
		t.Errorf("except 1 but get %s", list.head.next.data)
	}
}
func Test_InsertLast(t *testing.T)  {
	list := LinkedList{}
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertLast(3)
	if list.head.next.next.data != 3 {
		t.Errorf("except 3 but get %s", list.head.next.next.data)
	}
}
// func Test_AddByIdx_1(t *testing.T) {
// 	list := LinkedList{}
// 	list.InsertLast(1)
// 	list.InsertLast(2)
// 	list.InsertLast(3)
// 	list.AddByIdx(0, -1)
// 	if list.head.data != -1 {
// 		t.Errorf("excepted -1 but get %s", list.head.data)
// 	}
// }

func Test_GetFirst(t *testing.T) {
	list := LinkedList{}
	list.InsertFirst(2)
	v, ok := list.GetFirst()
	if !ok || v != 2 {
		t.Errorf("except 2 but get%s", v)
	}
}

func Test_GetLast(t *testing.T) {
	list := LinkedList{}
	list.InsertFirst(2)
	v, ok := list.GetLast()
	if !ok || v != 2 {
		t.Errorf("except 2 but get%s", v)
	}
}

func Test_GetSize(t *testing.T) {
	list := LinkedList{}
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(5)
	size := list.GetSize()
	if size != 3 {
		t.Errorf("except 2 but get%d", size)
	}
}
func Test_GetValues(t *testing.T) {
	list := LinkedList{}
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(5)
	vals := list.GetValues()

	var slice = []interface{}{5,3,2}
	for key, v := range slice {
		if v != slice[key] {
			t.Errorf("except %s but get%s", slice, vals)
		}
	}
}