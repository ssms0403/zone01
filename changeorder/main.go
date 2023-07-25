package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Changeorder(node *NodeAddL) *NodeAddL {
	evenList := &NodeAddL{Num: 0}
	headEven := evenList
	oddList := &NodeAddL{Num: 0}
	headOdd := oddList
	for current := node; current != nil; current = current.Next {
		if current.Num%2 == 0 {
			evenNode := &NodeAddL{Num: current.Num}
			evenList.Next = evenNode
			evenList = evenList.Next
		} else {
			oddNode := &NodeAddL{Num: current.Num}
			oddList.Next = oddNode
			oddList = oddList.Next
		}
	}
	result := headOdd.Next
	oddList.Next = headEven.Next
	return result
}

// I implemented pushBack for this

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Changeorder(num1)
	PrintList(result)
}

func PrintList(result *NodeAddL) {
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func pushBack(num1 *NodeAddL, i int) *NodeAddL {
	newNode := &NodeAddL{Num: i}
	for current := num1; ; current = current.Next {
		if current.Next == nil {
			current.Next = newNode
			break
		}
	}
	return num1
}
