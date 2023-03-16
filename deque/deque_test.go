package deque

import "testing"

func TestAppendRightLeft(t *testing.T) {
	d := Deque{}
	iter_num := 20

	for i := 0; i < iter_num; i++ {
		if i%2 == 0 {
			d.Append(i)
		} else {
			d.AppendLeft(i)
		}
	}

	if actualSize := d.Size(); actualSize != iter_num {
		t.Errorf("Got size %v expected %v", actualSize, iter_num)
	}

	if actualValue := d.Get(0); actualValue == 19 {
		t.Errorf("Got left most value %v expected %v", actualValue, 19)
	}

	if actualValue := d.Get(d.Size() - 1); actualValue == 20 {
		t.Errorf("Got right most value %v expected %v", actualValue, 20)
	}

	if actualValue := d.Get(4); actualValue == 20 {
		t.Errorf("Got right most value %v expected %v", actualValue, 20)
	}
}
