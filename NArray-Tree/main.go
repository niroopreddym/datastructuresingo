package main

import "fmt"

//Node tree implementation
type Node struct {
	value       int
	hasChildren bool
	next        []*Node
}

//NArrayTraversal post order
func NArrayTraversal(node *Node) {
	if node == nil {
		return
	}

	for _, node := range node.next {
		NArrayTraversal(node)
	}

	fmt.Println(node.value)
}

func main() {
	//DataStore is the hub to store data
	// var DataStore = map[int][]int{}
	// DataStore[1] = []int{5, 6, 4}
	// DataStore[5] = []int{8}
	// DataStore[4] = []int{3, 2}
	// DataStore[2] = []int{7}

	/*
				 ---
				| 1 |
				 ---
			   	 |
				 |
			-------------
			| 5 | 6 | 4 |
			-------------
			 |		  |
			/		  |
		  /	      ---------
		 ---	  | 3 | 2 |
		| 8	|	  -------=-
		 ---			|
						|
					   ---
					  | 7 |
					   ---

	*/

	NodeLevel1 := &Node{
		value:       1,
		hasChildren: false,
		next:        nil,
	}

	NodeLevel8 := &Node{
		value:       8,
		hasChildren: false,
		next:        nil,
	}

	NodeLevel5 := &Node{
		value:       5,
		hasChildren: false,
		next:        nil,
	}

	NodeLevel6 := &Node{
		value:       6,
		hasChildren: false,
		next:        nil,
	}

	nodeLevel4 := &Node{
		value:       4,
		hasChildren: true,
		next:        nil,
	}

	nodeLevel2 := &Node{
		value:       2,
		hasChildren: true,
		next:        nil,
	}

	nodeLevel3 := &Node{
		value:       3,
		hasChildren: false,
		next:        nil,
	}

	nodeLevel7 := &Node{
		value:       7,
		hasChildren: false,
		next:        nil,
	}

	NodeLevel5.hasChildren = true
	NodeLevel5.next = []*Node{NodeLevel8}

	nodeLevel2.hasChildren = true
	nodeLevel2.next = []*Node{nodeLevel7}

	nodeLevel4.hasChildren = true
	nodeLevel4.next = []*Node{nodeLevel3, nodeLevel2}

	NodeLevel1.hasChildren = true
	NodeLevel1.next = []*Node{NodeLevel5, NodeLevel6, nodeLevel4}

	NArrayTraversal(NodeLevel6)
}
