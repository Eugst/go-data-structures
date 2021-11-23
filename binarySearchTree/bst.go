package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n == nil {
		return
	}
	if val < n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) Find(val int) *Node {
	if n == nil {
		return nil
	}
	if val == n.Val {
		return n
	}
	if val < n.Val {
		return n.Left.Find(val)
	}
	return n.Right.Find(val)
}

func (n *Node) Delete(val int, parent *Node) {
	if n == nil {
		return
	}
	if val < n.Val {
		n.Left.Delete(val, n)
	} else if val > n.Val {
		n.Right.Delete(val, n)
	} else {
		if n.Left == nil && n.Right == nil {
			if parent.Left == n {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		} else if n.Left == nil {
			if parent.Left == n {
				parent.Left = n.Right
			} else {
				parent.Right = n.Right
			}
			// n = n.Right
		} else if n.Right == nil {
			if parent.Left == n {
				parent.Left = n.Left
			} else {
				parent.Right = n.Left
			}

		} else {
			min := n.Right.Min()
			n.Val = min.Val
			n.Right.Delete(min.Val, n)
		}
	}
}

func (n *Node) Min() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.Min()
}

func (n *Node) FindMin() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.FindMin()
}

func (n Node) Print() {
	if n.Left != nil {
		n.Left.Print()
	}
	fmt.Println(n.Val)
	if n.Right != nil {
		n.Right.Print()
	}
}
func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

func main() {
	root := &Node{Val: 5}
	root.Insert(3)
	root.Insert(2)
	root.Insert(4)
	root.Insert(8)
	root.Insert(6)
	root.Insert(7)
	root.Insert(1)
	root.Insert(12)
	root.Insert(11)
	root.Insert(10)
	root.Insert(13)
	// root.Print()
	print(os.Stdout, root, 0, 'M')
	root.Delete(12, nil)
	//root.Print()
	print(os.Stdout, root, 0, 'M')
}
