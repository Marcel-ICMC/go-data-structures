package main

import (
	"fmt"
)

type Deque struct {
	Left  *Node
	Right *Node
	Size  int
}

type Node struct {
	Left  *Node
	Right *Node
	Value any
}

func (d *Deque) Append(values ...any) {
	for _, value := range values {
		newNode := &Node{Value: value, Left: d.Right}

		if d.Left == nil {
			d.Left = newNode
		}

		if d.Right != nil {
			d.Right.Right = newNode
		}
		d.Right = newNode
		d.Size++
	}
}

func (d *Deque) AppendLeft(values ...any) {
	for _, value := range values {
		newNode := Node{Value: value, Right: d.Left}

		if d.Right == nil {
			d.Right = &newNode
		}

		if d.Left != nil {
			d.Left.Left = &newNode
		}
		d.Left = &newNode
		d.Size++
	}
}

func (d *Deque) Count() int {
	return d.Size
}

func (d *Deque) Copy() *Deque {
	copy := Deque{}
	for n := d.Left; n != nil; n = n.Right {
		copy.Append(n.Value)
	}
	return &copy
}

func (d *Deque) Pop() any {
	if d.Right != nil {
		r := d.Right.Value

		d.Right = d.Right.Left
		d.Size--

		if d.Right == nil {
			d.Left = nil
		} else {
			d.Right.Right = nil
		}
		return r
	}

	panic("No nodes to pop!")
}

func (d *Deque) PopLeft() any {
	if d.Left != nil {
		r := d.Left.Value

		d.Left = d.Left.Right
		d.Size--

		if d.Left == nil {
			d.Right = nil
		} else {
			d.Left.Left = nil
		}
		return r
	}

	panic("No nodes to pop!")
}

func (d *Deque) Index(x any) int {
	n := d.Left
	for i := 0; n != nil; i++ {
		if n.Value == x {
			return i
		}
		n = n.Right
	}

	panic(fmt.Sprintf("There are no value %d in the deque", x))
}

func (d *Deque) Insert(i int, x any) {
	if i > d.Size {
		panic("Position unreachable!")
	} else if i == 0 {
		d.AppendLeft(x)
		return
	} else if i == d.Size {
		d.Append(x)
		return
	}

	n := d.Left
	for j := 1; j < i; j++ {
		n = n.Right
	}

	d.Size++
	new := Node{Value: x, Left: n, Right: n.Right}
	n.Right.Left = &new
	n.Right = &new
}

func (d *Deque) Remove(x any) {
	for n := d.Left; n != nil; n = n.Right {
		if n.Value == x {
			if n.Left == nil {
				d.Left = n.Right
			} else {
				n.Left.Right = n.Right
			}

			if n.Right == nil {
				d.Right = n.Left
			} else {
				n.Right.Left = n.Left
			}

			d.Size--
			return
		}
	}

	panic(fmt.Sprintf("There are no value %d in the deque", x))
}

func (d *Deque) Revese() {
	for n := d.Left; n != nil; n = n.Left {
		n.Left, n.Right = n.Right, n.Left
	}
	d.Left, d.Right = d.Right, d.Left
}

func (d *Deque) rotate(n *Node) {
	if d.Left != n {
		n.Left.Right = nil

		d.Left.Left = d.Right
		d.Right.Right = d.Left

		d.Left = n
		d.Right, n.Left = n.Left, nil
	}
}

func (d *Deque) Rotate(i int) {
	j := i % d.Size
	if j > d.Size/2 {
		j -= d.Size
	}

	n := d.Left
	if j > 0 {
		n = d.Right
		for k := 1; k < j; k++ {
			n = n.Left
		}
	} else if j < 0 {
		for k := 1; k < -j; k++ {
			n = n.Right
		}
	}

	d.rotate(n)
}

func (d Deque) String() string {
	n := d.Left
	r := ""
	for n != nil {
		r += fmt.Sprint((*n).Value) + " "
		n = (*n).Right
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
