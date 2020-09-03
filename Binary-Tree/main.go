package main

import "fmt"

//Node tree implementation
type Node struct {
	val   int
	left  *Node
	right *Node
}

//PostorderTraversal post order
func PostorderTraversal(node *Node) {

	if node == nil {
		return
	}

	PostorderTraversal(node.left)
	PostorderTraversal(node.right)
	fmt.Printf("%d ", node.val)
}

func main() {
	/*
		      1
			/   \
		   7	  4
		  / \    /  \
		3   5    6   5
	*/

	node := Node{1, nil, nil}
	node.left = &Node{7, nil, nil}
	node.left.left = &Node{3, nil, nil}
	node.left.right = &Node{5, nil, nil}
	node.right = &Node{4, nil, nil}
	node.right.left = &Node{6, nil, nil}
	node.right.right = &Node{5, nil, nil}

	//LRM
	PostorderTraversal(&node)
}
