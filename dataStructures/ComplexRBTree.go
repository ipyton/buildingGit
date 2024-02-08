package dataStructures


type RBTNode[KEY any,VALUE any] struct {
	parent * RBTNode[KEY,VALUE]
	left * RBTNode[KEY,VALUE]
	right * RBTNode[KEY, VALUE]
	key KEY
	value VALUE
}

type ComplexRBTree[KEY any, VALUE any] struct {
	root * RBTNode
	empty * RBTNode
}