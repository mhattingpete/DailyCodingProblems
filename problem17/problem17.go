package main

import (
    "fmt"
    "strings"
    "regexp"
)

/*
Suppose we represent our file system by a string in the following manner:

The string "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" represents:

dir
    subdir1
    subdir2
        file.ext

The directory dir contains an empty sub-directory subdir1 and a sub-directory subdir2 containing a file file.ext.

The string "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" represents:

dir
    subdir1
        file1.ext
        subsubdir1
    subdir2
        subsubdir2
            file2.ext

The directory dir contains two sub-directories subdir1 and subdir2. subdir1 contains a file file1.ext and an empty second-level sub-directory subsubdir1. subdir2 contains a second-level sub-directory subsubdir2 containing a file file2.ext.

We are interested in finding the longest (number of characters) absolute path to a file within our file system. For example, in the second example above, the longest absolute path is "dir/subdir2/subsubdir2/file2.ext", and its length is 32 (not including the double quotes).

Given a string representing the file system in the above format, return the length of the longest absolute path to a file in the abstracted file system. If there is no file in the system, return 0.

Note:

The name of a file contains at least a period and an extension.

The name of a directory or sub-directory will not contain a period.
*/

type Tree struct {
	Nodes []*Tree
	Value string
}

func (t *Tree) addNode(s string) int {
	n := &Tree{}
	n.Value = s
	t.Nodes = append(t.Nodes, n)
	return len(t.Nodes)-1
}

func traverseNodes(t *Tree, padding string, p string, hasRightSibling bool) string {
	s := ""
	if t != nil {
		s += "\n" + padding + p + t.Value
		nextPadding := padding
		if hasRightSibling == true {
			nextPadding += "│  "
		} else {
			nextPadding += "   "
		}
		pointerRight := "└──Node:"
		var pointerLeft string
		if len(t.Nodes) > 1 {
			pointerLeft = "├──Node:"
		} else {
			pointerLeft = "└──Node:"
		}
		for i := 0; i < len(t.Nodes)-1; i++ {
			s += traverseNodes(t.Nodes[i], nextPadding, pointerLeft, len(t.Nodes) > 1)
		}
		if idx := len(t.Nodes)-1; idx >= 0 {
			s += traverseNodes(t.Nodes[idx], nextPadding, pointerRight, false)
		}
	}
	return s
}

func (t *Tree) print() string {
	if t == nil {
		return ""
	}
	s := "Root:"
	s += t.Value
	pointerRight := "└──Node:"
	var pointerLeft string
	if len(t.Nodes) > 1 {
		pointerLeft = "├──Node:"
	} else {
		pointerLeft = "└──Node:"
	}
	for i := 0; i < len(t.Nodes)-1; i++ {
		s += traverseNodes(t.Nodes[i], "", pointerLeft, len(t.Nodes) > 1)
	}
	if idx := len(t.Nodes)-1; idx >= 0 {
		s += traverseNodes(t.Nodes[idx], "", pointerRight, false)
	}
	return s
}

func buildTree(dirPath string) *Tree {
	var idx int
	root := &Tree{}
	re := regexp.MustCompile("\n\t[a-zA-Z]")
	splitted := re.Split(dirPath, -1)
	root.Value = splitted[0]
	if len(splitted) > 1 {
		splitted = splitted[1:]
		for i := 0; i < len(splitted); i++ {
			if s := splitted[i]; strings.Contains(splitted[i], ".") {
				idx = root.addNode(s)
				root.Nodes[idx].buildSubTree(s)
			}
		}
	}
	return root
}

func (t *Tree) buildSubTree(dirPath string) {
	if strings.Contains(dirPath, ".") {
		var idx int
		dirPath = strings.ReplaceAll(dirPath, "\t\t", "\t")
		dirPath = "s"+dirPath
		re := regexp.MustCompile("\n\t[a-z]")
		splitted := re.Split(dirPath, -1)
		if len(splitted) > 1 {
			t.Value = splitted[0]
		} else {
			t.Value = "f"+splitted[0][1:]
		}
		if len(splitted) > 1 {
			splitted = splitted[1:]
			for i := 0; i < len(splitted); i++ {
				if s := splitted[i]; strings.Contains(splitted[i], ".") {
					idx = t.addNode(s)
					t.Nodes[idx].buildSubTree(s)
				}
			}
		}
	}
}

func (t *Tree) findSubTreeLenght() (string, int) {
	if t != nil {
		var maxPath, cpath string
		var length, count int
		path := t.Value
		maxLen := 0
		for i := 0; i < len(t.Nodes); i++ {
			cpath, length = t.Nodes[i].findSubTreeLenght()
			if length >= maxLen {
				maxPath = cpath
				maxLen = length
			}
		}
		path += "/" + maxPath
		count = len(path)
		return path, count
	}
	return "", 0
}

func longestPath(s string) (string, int) {
	var maxPath, cpath string
	var length, count int
	maxLen := 0
	t := buildTree(s)
	path := t.Value
	fmt.Println(t.print())	
	for i := 0; i < len(t.Nodes); i++ {
		cpath, length = t.Nodes[i].findSubTreeLenght()
		if length >= maxLen {
			maxPath = cpath
			maxLen = length
		}
	}
	path += "/" + maxPath
	path = path[:len(path)-1]
	count = len(path)
	return path, count
}

func main() {
	s := "dir\n\tsubdir3\n\t\tfile3.ext\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	path, count := longestPath(s)
	fmt.Println(path, count)
}