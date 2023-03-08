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
	Value int
}

func (d *Deque) Append(value int) {
	newNode := Node{Value: value, Left: d.Right}

	if d.Left == nil {
		d.Left = &newNode
	}

	if d.Right != nil {
		d.Right.Right = &newNode
	}
	d.Right = &newNode
	d.Size++
}

func (d *Deque) AppendLeft(value int) {
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

func (d *Deque) Count() int {
	return d.Size
}

func (d *Deque) Pop() int {
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

func (d *Deque) PopLeft() int {
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

func (d *Deque) Index(x int) int {
	n := d.Left
	for i := 0; n != nil; i++ {
		if n.Value == x {
			return i
		}
		n = n.Right
	}

	panic(fmt.Sprintf("There are no value %d in the deque", x))
}

func (d *Deque) Insert(i int, x int) {
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

func (d Deque) String() string {
	n := d.Left
	r := ""
	for n != nil {
		r += fmt.Sprintf("%d ", (*n).Value)
		n = (*n).Right
	}

	return r
}

func main() {
	t := Deque{}
	t.Append(11)
	t.Append(12)
	t.Append(13)
	t.Pop()
	t.Insert(0, 0)
	t.Insert(3, 1)
	fmt.Printf("%s\n", t)
}
