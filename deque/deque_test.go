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

	if actualValue := d.Get(0); actualValue != 19 {
		t.Errorf("Got left most value %v expected %v", actualValue, 19)
	}

	if actualValue := d.Get(d.Size() - 1); actualValue != 18 {
		t.Errorf("Got right most value %v expected %v", actualValue, 18)
	}

	if actualValue := d.Get(4); actualValue != 11 {
		t.Errorf("Got value %v at index 4 expected %v", actualValue, 11)
	}
}

func TestAppend(t *testing.T) {
	d := Deque{}
	a := []any{10, 2, 3, 4, 5}
	d.Append(a...)

	if actualSize := d.Size(); actualSize != len(a) {
		t.Errorf("Got size %v expected %v", actualSize, len(a))
	}

	if actualValue := d.Get(0); actualValue != a[0] {
		t.Errorf("Got left most value %v expected %v", actualValue, a[0])
	}

	if actualValue := d.Get(d.Size() - 1); actualValue != a[len(a)-1] {
		t.Errorf("Got right most value %v expected %v", actualValue, a[len(a)-1])
	}

	if actualValue := d.Get(4); actualValue != a[4] {
		t.Errorf("Got value %v at index 4 expected %v", actualValue, a[4])
	}
}

func TestAppendLeft(t *testing.T) {
	d := Deque{}
	a := []any{1, 2, 3, 4, 5}

	// This is equivalent to
	// d.Append(a...)
	// d.Rotate()
	d.AppendLeft(a...)

	if actualSize := d.Size(); actualSize != len(a) {
		t.Errorf("Got size %v expected %v", actualSize, len(a))
	}

	if actualValue := d.Get(d.Size() - 1); actualValue != a[0] {
		t.Errorf("Got right most value %v expected %v", actualValue, a[len(a)-1])
	}

	if actualValue := d.Get(4); actualValue != a[len(a)-5] {
		t.Errorf("Got value %v at index 4 expected %v", actualValue, a[len(a)-5])
	}
}

func TestPop(t *testing.T) {
	d := Deque{}
	d.Append("a", "b", "c")
	d.Pop()

	expected_size := 2
	expected_value := "b"
	if actualSize := d.Size(); actualSize != expected_size {
		t.Errorf("Got size %v expected %v", actualSize, expected_size)
	}

	if actualValue := d.Get(d.Size() - 1); actualValue != expected_value {
		t.Errorf("Got right most value %v expected %s", actualValue, expected_value)
	}

	d.Pop()
	expected_size = 1
	expected_value = "a"
	if actualSize := d.Size(); actualSize != expected_size {
		t.Errorf("Got size %v expected %v", actualSize, expected_size)
	}

	if actualValue := d.Get(d.Size() - 1); actualValue != expected_value {
		t.Errorf("Got right most value %v expected %s", actualValue, expected_value)
	}
}

func TestPopLeft(t *testing.T) {
	d := Deque{}
	d.Append("a", "b", "c")
	d.PopLeft()

	expected_size := 2
	expected_value := "b"
	if actualSize := d.Size(); actualSize != expected_size {
		t.Errorf("Got size %v expected %v", actualSize, expected_size)
	}

	if actualValue := d.Get(0); actualValue != expected_value {
		t.Errorf("Got left most value %v expected %s", actualValue, expected_value)
	}

	d.PopLeft()
	expected_size = 1
	expected_value = "c"
	if actualSize := d.Size(); actualSize != expected_size {
		t.Errorf("Got size %v expected %v", actualSize, expected_size)
	}

	if actualValue := d.Get(0); actualValue != expected_value {
		t.Errorf("Got left most value %v expected %s", actualValue, expected_value)
	}
}

func TestReverse(t *testing.T) {
	d := Deque{}
	d.Append("a", "b", "c")
	d_r := d.Copy()
	d_r.Reverse()

	if d_r.Size() != d.Size() {
		t.Errorf("Got size %v expected %v", d_r.Size(), d.Size())
	}

	for i := 0; i < d_r.Size(); i++ {
		d_value := d.Pop()
		d_r_value := d_r.PopLeft()
		if d_value != d_r_value{
			t.Errorf("Got value %v at index %d expected %v", d_r_value, i, d_value)
		}
	}
}
