package main

import (
	// "fmt"
	"piscine"
	// student "student"
	// "github.com/01-edu/go-tests/lib/challenge"
	// "github.com/01-edu/go-tests/solutions"
)

func parentListTransp(root *piscine.TreeNode) string {
	if root == nil {
		return ""
	}

	var parent string

	if root.Parent == nil {
		parent = "nil"
	} else {
		parent = root.Parent.Data
	}

	r := "Node: " + root.Data + " Parent: " + parent + "\n"
	r += parentListTransp(root.Left) + parentListTransp(root.Right)
	return r
}

func FormatTree_transp(root *piscine.TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree_transp(root, "")
	return res
}

func formatSubTree_transp(root *piscine.TreeNode, prefix string) string {
	if root == nil {
		return ""
	}

	var res string

	hasLeft := root.Left != nil
	hasRight := root.Right != nil

	if !hasLeft && !hasRight {
		return res
	}

	res += prefix
	if hasLeft && hasRight {
		res += "├── "
	}

	if !hasLeft && hasRight {
		res += "└── "
	}

	if hasRight {
		printStrand := (hasLeft && hasRight && (root.Right.Right != nil || root.Right.Left != nil))
		newPrefix := prefix
		if printStrand {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		res += root.Right.Data + "\n"
		res += formatSubTree_transp(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree_transp(root.Left, prefix+"    ")
	}
	return res
}
