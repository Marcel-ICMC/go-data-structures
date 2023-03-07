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

func (d *Deque) PrintList() {
	n := (*d).Left
	for n != nil {
		fmt.Printf("%d ", (*n).Value)
		n = (*n).Right
	}

	fmt.Println()
}

func main() {
	t := Deque{}
	t.Append(11)
	t.Append(12)
	t.Append(12)
	t.Append(12)
	t.Append(12)
	t.AppendLeft(120)
	t.AppendLeft(120)
	t.AppendLeft(120)
	t.PrintList()
}
