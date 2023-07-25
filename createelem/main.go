package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
}

func (n *Node) CreateElem(value int) {
	n.Data = value
}

func main() {
	n := &Node{}
	n.CreateElem(1234)
	fmt.Println(n)
}
