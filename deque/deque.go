package deque

import (
	"fmt"
)

type Deque struct {
	left  *Node
	right *Node
	size  int
}

type Node struct {
	left  *Node
	right *Node
	value any
}

func (d *Deque) Empty() bool {
	return d.size == 0
}

func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) Append(values ...any) {
	for _, value := range values {
		newNode := &Node{value: value, left: d.right}

		if d.left == nil {
			d.left = newNode
		}

		if d.right != nil {
			d.right.right = newNode
		}
		d.right = newNode
		d.size++
	}
}

func (d *Deque) AppendLeft(values ...any) {
	for _, value := range values {
		newNode := Node{value: value, right: d.left}

		if d.right == nil {
			d.right = &newNode
		}

		if d.left != nil {
			d.left.left = &newNode
		}
		d.left = &newNode
		d.size++
	}
}

func (d *Deque) Count() int {
	return d.size
}

func (d *Deque) Copy() *Deque {
	copy := Deque{}
	for n := d.left; n != nil; n = n.right {
		copy.Append(n.value)
	}
	return &copy
}

func (d *Deque) Pop() any {
	if d.right != nil {
		r := d.right.value

		d.right = d.right.left
		d.size--

		if d.right == nil {
			d.left = nil
		} else {
			d.right.right = nil
		}
		return r
	}

	panic("No nodes to pop!")
}

func (d *Deque) PopLeft() any {
	if d.left != nil {
		r := d.left.value

		d.left = d.left.right
		d.size--

		if d.left == nil {
			d.right = nil
		} else {
			d.left.left = nil
		}
		return r
	}

	panic("No nodes to pop!")
}

func (d *Deque) Get(l int) any {
	if l >= d.size {
		panic(fmt.Sprintf("Index %d unreachable, actual deque size is %d", l, d.size))
	}

	n := d.left
	for i := 0; i < l; i++ {
		n = n.right
	}
	return n.value
}

func (d *Deque) Index(x any) int {
	n := d.left
	for i := 0; n != nil; i++ {
		if n.value == x {
			return i
		}
		n = n.right
	}

	panic(fmt.Sprintf("There are no value %d in the deque", x))
}

func (d *Deque) Insert(i int, x any) {
	if i > d.size {
		panic("Position unreachable!")
	} else if i == 0 {
		d.AppendLeft(x)
		return
	} else if i == d.size {
		d.Append(x)
		return
	}

	n := d.left
	for j := 1; j < i; j++ {
		n = n.right
	}

	d.size++
	new := Node{value: x, left: n, right: n.right}
	n.right.left = &new
	n.right = &new
}

func (d *Deque) Remove(x any) {
	for n := d.left; n != nil; n = n.right {
		if n.value == x {
			if n.left == nil {
				d.left = n.right
			} else {
				n.left.right = n.right
			}

			if n.right == nil {
				d.right = n.left
			} else {
				n.right.left = n.left
			}

			d.size--
			return
		}
	}

	panic(fmt.Sprintf("There are no value %d in the deque", x))
}

func (d *Deque) Revese() {
	for n := d.left; n != nil; n = n.left {
		n.left, n.right = n.right, n.left
	}
	d.left, d.right = d.right, d.left
}

func (d *Deque) rotate(n *Node) {
	if d.left != n {
		n.left.right = nil

		d.left.left = d.right
		d.right.right = d.left

		d.left = n
		d.right, n.left = n.left, nil
	}
}

func (d *Deque) Rotate(i int) {
	j := i % d.size
	if j > d.size/2 {
		j -= d.size
	}

	n := d.left
	if j > 0 {
		n = d.right
		for k := 1; k < j; k++ {
			n = n.left
		}
	} else if j < 0 {
		for k := 1; k < -j; k++ {
			n = n.right
		}
	}

	d.rotate(n)
}

func (d Deque) String() string {
	n := d.left
	r := ""
	for n != nil {
		r += fmt.Sprint((*n).value) + " "
		n = (*n).right
	}

	return r
}

func main() {
	t := Deque{}

	var test_dots []any
	for i := 0; i < 20; i++ {
		test_dots = append(test_dots, i)
	}

	t.Append(test_dots...)
	fmt.Printf("%s\n", t)
	t.Pop()

	t.Remove(15)
	fmt.Printf("%s\n", t)

	t.Revese()
	fmt.Printf("%s\n", t)

	t.Rotate(1)
	fmt.Printf("%s\n", t)

	t.Rotate(20)
	fmt.Printf("%s\n", t)

	t.Rotate(-1)
	fmt.Printf("%s\n", t)

	t.Rotate(0)
	fmt.Printf("%s\n", t)

	t.Rotate(5)
	fmt.Printf("%s\n", t)

	a := t.Copy()
	fmt.Printf("%s\n", a)
}
