package dataStructures

type Comparable struct {

}


type TreeNode struct{
	black bool
	key string
	value string
	left * TreeNode
	right * TreeNode
	n int
}


func newTreeNode(key string, value string, size int,black bool) * TreeNode {
	return & TreeNode{key: key, value: value, n:size, black: black}
}



func (this * TreeNode) isBlack() bool {
	if this == nil {
		return true
	}
	return this.black
}



func size(node * TreeNode) int {
	if node == nil {
		return 0
	}
	return node.n
}




type RBTree struct {
	compare func()
}

func flipColors(tree * TreeNode) {

}

func compare(a string, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func (rbTree * RBTree) putRecursively(key string, value string){

}

func (rbTree * RBTree) put(tree * TreeNode, key string, value string) * TreeNode{
	if tree == nil {
		return newTreeNode(key,value, 1, false)
	}
	result := compare(tree.key, key)
	if result < 0 {
		tree.left = rbTree.put(tree.left, key, value)
	} else if result > 0{
		tree.right = rbTree.put(tree.right, key, value)
	} else {
		tree.value = value
	}

	if tree.left.isBlack() && !tree.right.isBlack() {
		tree = rotateLeft(tree)
	}
	if !tree.left.isBlack() && !tree.left.left.isBlack() {
		tree = rotateRight(tree)
	}
	if !tree.left.isBlack() && !tree.right.isBlack() {
		flipColors(tree)
	}

	tree.n = size(tree.left)  + size(tree.right) + 1
	return tree
}

func (rbTree * RBTree) deleteMin(){

}

func (rbTree * RBTree) deleteMax() {

}

func (rbTree * RBTree) delete() {

}

func rotateLeft(node * TreeNode) * TreeNode {
	right := node.right
	node.right = right.left
	right.left = node
	right.black = node.black
	node.black = false
	right.n = node.n
	node.n = 1 + size(node.left) + size(node.right)
	return right
}


func rotateRight(node * TreeNode) * TreeNode {
	left := node.left
	node.left = left.right
	left.right = node
	left.black = node.black
	node.black = false
	left.n = node.n
	node.n = 1 + size(node.left) + size(node.right)
	return left
}










