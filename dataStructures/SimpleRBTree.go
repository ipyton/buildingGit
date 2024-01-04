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
	compare func(string, string) int
	root * TreeNode
}

func flipColors(tree * TreeNode) {
	tree.left.black = true
	tree.right.black = true
	tree.black = false
}

func compare(a string, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func (rbTree * RBTree) put(key string, value string){
	node := rbTree.putRecursively(rbTree.root, key, value)
	node.black = true
}

func (rbTree * RBTree) putRecursively(tree * TreeNode, key string, value string) * TreeNode{
	if tree == nil {
		return newTreeNode(key,value, 1, false)
	}
	result := compare(tree.key, key)
	if result < 0 {
		tree.left = rbTree.putRecursively(tree.left, key, value)
	} else if result > 0{
		tree.right = rbTree.putRecursively(tree.right, key, value)
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

func (node * TreeNode) isEmpty() bool {
	if node.n == 0{
		return true
	}
	return false
}

func reverseFlipColors(tree * TreeNode) {
	tree.left.black = false
	tree.right.black = false
	tree.black = true
}

func moveRedLeft(treeNode * TreeNode) * TreeNode{
	reverseFlipColors(treeNode)
	if !treeNode.right.left.isBlack() { // sibling node is not 2 node
		treeNode.right = rotateRight(treeNode.right)
		treeNode = rotateLeft(treeNode)
	}
	// if not just regard the
	return treeNode
}
func (rbTree *RBTree) balance(node * TreeNode) *TreeNode{
	if !node.right.isBlack() {
		node = rotateLeft(node)
	}
	return node
}

func (rbTree * RBTree) recursivelyDeleteMin(node * TreeNode) * TreeNode {
	if node.left == nil {
		return nil
	}
	//left son is 2 node
	if node.left.isBlack() && node.left.left.isBlack() {
		node = moveRedLeft(node)
	}

	node.left = rbTree.recursivelyDeleteMin(node.left)
	return rbTree.balance(node)
}

func (rbTree * RBTree) deleteMin() {
	if rbTree.root.left.isBlack() && rbTree.root.right.isBlack() {
		rbTree.root.black = false
	}
	rbTree.root = rbTree.recursivelyDeleteMin(rbTree.root)
	if rbTree.root.isEmpty(){
		rbTree.root.black = true
	}
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

func (rbTree * RBTree) moveRedRight(node * TreeNode) * TreeNode{
	reverseFlipColors(node)
	if !node.left.left.isBlack() {
		node = rotateRight(node)
	}
	return node
}

func (rbTree * RBTree) recursivelyDeleteMax(node * TreeNode) * TreeNode {
	if node.left.isBlack() && node.right.isBlack() {
		node.black = false
	}
	if node.right == nil {
		return nil
	}
	if node.right.isBlack() && node.right.left.isBlack() {
		node = rbTree.moveRedRight(node)
	}

	node.right = rbTree.recursivelyDeleteMax(node.right)
	return rbTree.balance(node)
}

func (rbTree * RBTree) deleteMax() {
	if !rbTree.root.left.isBlack() && !rbTree.root.right.isBlack() {
		rbTree.root.black = false
	}
	rbTree.root = rbTree.recursivelyDeleteMax(rbTree.root.right)
	if !rbTree.root.isEmpty() {
		rbTree.root.black = true
	}
}

func (rBTree * RBTree) recursivelyDelete(){

}

func (rbTree * RBTree) delete() {

}


