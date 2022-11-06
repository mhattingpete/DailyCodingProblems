package main

import (
    "fmt"
    "encoding/json"
)

type Node struct {
	value string
	left *Node
	right *Node
}

func (n *Node) Serialize() string {
	var leftNode string
	var rightNode string
	if (n.left != nil) {
		leftNode = n.left.Serialize()
	}
	if (n.right != nil) {
		rightNode = n.right.Serialize()
	}
	data := map[string]interface{}{
		"value": n.value,
		"left": leftNode,
		"right": rightNode,
	}
	jsondata, _ := json.Marshal(data)
	return string(jsondata)
}

func Deserialize(nodeString string) Node {
	var n Node
	var result map[string]interface{}
	json.Unmarshal([]byte(nodeString), &result)
	n.value = result["value"].(string)
	if result["left"] != "" {
		leftNode := Deserialize(result["left"].(string))
		n.left = &leftNode
	}
	if result["right"] != "" {
		rightNode := Deserialize(result["right"].(string))
		n.right = &rightNode
	}
	return n
}

func main() {
	n := Node{"root",&Node{"left",&Node{"left.left",nil,nil},nil},&Node{"right",nil,nil}}
    fmt.Println(n.left.left.value) 
    fmt.Println(n.Serialize())
    n2 := Deserialize(n.Serialize())
    fmt.Println(n2.left.left.value)
}