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