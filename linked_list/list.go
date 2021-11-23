package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	len  int
}

// Add items to the end
func (l *List) Push(val int) {
	node := ListNode{Val: val}
	cur := l.Head
	if cur == nil {
		l.Head = &node
	} else {
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &node
	}
	l.len++
}

// Remove an item from the end
func (l *List) Pop() int {
	if l.Head == nil {
		return 0
	}
	cur := l.Head
	prev := l.Head
	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}
	val := cur.Val
	prev.Next = nil
	l.len--
	return val
}

// Remove an item from the beginning
func (l *List) shift() int {
	if l.Head == nil {
		return 0
	}
	val := l.Head.Val
	l.Head = l.Head.Next
	l.len--
	return val
}

// Add items to the beginning
func (l *List) Unshift(val int) {
	node := ListNode{Val: val}
	node.Next = l.Head
	l.Head = &node
	l.len++
}

func (l *List) deleteByVal(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		l.len--
		return
	}
	node := l.Head
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			l.len--
			return
		}
		node = node.Next
	}
}

func (l *List) Len() int {
	return l.len
}

func (l List) Print() {
	for l.Head != nil {
		print(l.Head.Val, " ")
		l.Head = l.Head.Next
	}
	println()
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	l.Print()
	fmt.Println(l.Pop())
	l.Print()
	fmt.Println(l.shift())
	l.Print()
	l.Unshift(7)
	l.Print()
	l.deleteByVal(3)
	l.Print()
}
