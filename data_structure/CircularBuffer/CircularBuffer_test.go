package CircularBuffer

import(
	"testing"
)

var cb = CircularBuffer{}

func Test_insertValue(t *testing.T) {
	// cb.inser
	for i := 0; i < 10; i++ {
		cb.InsertValue(i)
	}
	excepted := [10]interface{}{0,1,2,3,4,5,6,7,8,9}
	
	if excepted != cb.data {
		t.Errorf("except %s, but get %s", excepted, cb.data)
	}
}
func Test_getValuesFromPos(t *testing.T) {
	excepted := [10]interface{}{9,0,1,2,3,4,5,6,7,8}
	
	if vals, _ := cb.GetValuesFromPos(9); excepted != vals {
		t.Errorf("except %s, but get %s", excepted, cb.data)
	}
}

 