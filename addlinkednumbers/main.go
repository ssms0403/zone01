package main

import (
	"fmt"
)

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {
	result := &NodeAddL{Num: num1.Num + num2.Num}
	prev, head := result, result
	for current1, current2 := num1.Next, num2.Next; current1 != nil && current2 != nil; current1, current2 = current1.Next, current2.Next {
		add := current1.Num + current2.Num
		if add <= 9 {
			result.Next = &NodeAddL{Num: add}
		} else {
			result.Next = &NodeAddL{Num: 0}
			prev.Num += add - 9
		}
		result = result.Next
		prev = result
	}
	return head
}

func pushFront(node *NodeAddL, num int) *NodeAddL {
	return &NodeAddL{Next: node, Num: num}
}

func main() {
	// 3 -> 1 -> 5
	num1 := &NodeAddL{Num: 8}
	num1 = pushFront(num1, 1)
	num1 = pushFront(num1, 9)
	PrintList(num1)

	// 5 -> 9 -> 2
	num2 := &NodeAddL{Num: 1}
	num2 = pushFront(num2, 7)
	num2 = pushFront(num2, 9)
	PrintList(num2)

	// 9 -> 0 -> 7
	result := AddLinkedNumbers(num1, num2)
	PrintList(result)
}

func PrintList(num1 *NodeAddL) {
	for tmp := num1; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
