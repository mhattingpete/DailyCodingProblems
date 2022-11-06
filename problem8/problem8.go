package main

import (
    "fmt"
    "strconv"
)

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

func (t *Tree) getValue() string {
	return strconv.Itoa(t.Value)
}

func (t *Tree) addLeftNode(v int) {
	n := &Tree{Left:nil,Value:v,Right:nil}
	t.Left = n
}

func (t *Tree) addRightNode(v int) {
	n := &Tree{Left:nil,Value:v,Right:nil}
	t.Right = n
}

func (t *Tree) nodesEqual() bool {
	if t != nil {
		var lv, rv int
		v := t.Value
		if t.Left != nil && t.Right != nil {
			lv = t.Left.Value
			rv = t.Right.Value
			return lv == v && rv == v
		} else if t.Left != nil {
			lv = t.Left.Value
			return lv == v
		} else if t.Right != nil {
			rv = t.Right.Value
			return rv == v
		} else {
			return true
		}
	}
	return false
}

func (t *Tree) univalTree() int {
	c := 0
	if t != nil {
		if t.nodesEqual() == true {
			c++
		}
		if t.Left != nil {
			c += t.Left.univalTree()
		}
		if t.Right != nil {
			c += t.Right.univalTree()
		}
	}
	return c
}

func traverseNodes(t *Tree, padding string, p string, hasRightSibling bool) string {
	s := ""
	if t != nil {
		s += "\n" + padding + p + t.getValue()
		nextPadding := padding
		if hasRightSibling == true {
			nextPadding += "│  "
		} else {
			nextPadding += "   "
		}
		pointerRight := "└──Right:"
		var pointerLeft string
		if t.Right != nil {
			pointerLeft = "├──Left:"
		} else {
			pointerLeft = "└──Left:"
		}
		s += traverseNodes(t.Left, nextPadding, pointerLeft, t.Right != nil)
		s += traverseNodes(t.Right, nextPadding, pointerRight, false)
	}
	return s
}

func (t *Tree) traversePreOrder() string {
	if t == nil {
		return ""
	}
	s := "Root:"
	s += t.getValue()
	pointerRight := "└──Right:"
	var pointerLeft string
	if t.Right != nil {
		pointerLeft = "├──Left:"
	} else {
		pointerLeft = "└──Left:"
	}
	s += traverseNodes(t.Left, "", pointerLeft, t.Right != nil)
	s += traverseNodes(t.Right, "", pointerRight, false)
	return s
}

func main() {
	t := Tree{Left:nil,Value:0,Right:nil}
	t.addLeftNode(1)
	t.addRightNode(0)
	t.Right.addLeftNode(1)
	t.Right.addRightNode(0)
	t.Right.Left.addLeftNode(1)
	t.Right.Left.addRightNode(1)
	fmt.Println(t.traversePreOrder())
	fmt.Println(t.univalTree())
}