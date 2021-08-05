package tree

import (
	"GolangDemo/util"
	"fmt"
)

type Node struct {
	V     interface{}
	Left  *Node
	Right *Node
}

type BST struct {
	Root   *Node
	Length int
}

func Add(value interface{}, t *BST) *BST {

	if t == nil {
		t = &BST{Root: &Node{
			V: value,
		},
			Length: 1}
		//fmt.Println(t.Root.V)
		return t
	}
	if (*t).Root == nil {
		t.Root = &Node{
			V: value,
		}
		t.Length = 1
		return t
	}
	add(value, t.Root)
	t.Length++
	return t
}

func add(value interface{}, node *Node) {
	newNode := &Node{
		V: value,
	}
	res, err := util.Compare(value, node.V)
	if err != nil {
		panic(err.Error())
	}
	if res == 1 && node.Right == nil {
		node.Right = newNode
		return
	} else if res == -1 && node.Left == nil {
		node.Left = newNode
		return
	}
	if res == 1 {
		add(value, node.Right)
	} else if res == -1 {
		add(value, node.Left)
	}
}

func Contains(bst *BST, value interface{}) bool {
	return contains(bst.Root, value)
}

func contains(node *Node, value interface{}) bool {
	if node == nil {
		return false
	}
	res, _ := util.Compare(value, node.V)
	if res == 0 {
		return true
	}
	if res == 1 {
		return contains(node.Right, value)
	}
	return contains(node.Left, value)
}

//前序遍历
func PreOrder(bst *BST) {
	preOrder(bst.Root)
}

/*
	前、中、后序遍历在代码上只是体现在执行内容放在左右孩子的什么位置
*/
func preOrder(node *Node) {
	if node == nil {
		return
	}
	preOrder(node.Left)

	preOrder(node.Right)
	fmt.Println(node.V)

}
