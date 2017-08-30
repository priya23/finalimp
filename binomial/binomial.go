package binomial

import (
	"fmt"
	//"github.com/priya23/finalimp"
)

type BinomialHeap struct {
	forest_head *BinomialHeapNode

	size int
}

func NewBinomialHeap() *BinomialHeap {
	return &BinomialHeap{
		forest_head: nil,
		size:        0,
	}
}

func (k *BinomialHeap) Give(val string, pro int) {
	k.Insert(pro)
}

func (k *BinomialHeap) Take() int {
	return k.Pop()
}
func (bh *BinomialHeap) Insert(value int) {
	bh.size += 1

	newnode := newBinomialHeapNode(value)
	bh.insert(newnode)
}

func (bh *BinomialHeap) Pop() int {
	// Assume the queue is not empty.
	bh.size -= 1

	minnode := getMinimumNode(bh.forest_head)
	removeFromLinkedList(&bh.forest_head, minnode)

	for _, child := range nodeIterator(minnode.children_head) {
		removeFromLinkedList(&minnode.children_head, child)
		bh.insert(child)
	}

	return minnode.value
}

func (bh *BinomialHeap) Peek() int {
	return getMinimumNode(bh.forest_head).value
}

func (bh *BinomialHeap) Size() int {
	return bh.size
}

func (bh *BinomialHeap) Merge(other *BinomialHeap) {
	bh.size += other.size

	for _, child := range nodeIterator(other.forest_head) {
		removeFromLinkedList(&other.forest_head, child)
		bh.insert(child)
	}
}

func (bh *BinomialHeap) insert(newnode *BinomialHeapNode) {
	srnode := getNodeWithOrder(bh.forest_head, newnode.order)

	if srnode == nil {
		insertIntoLinkedList(&bh.forest_head, newnode)
	} else {
		removeFromLinkedList(&bh.forest_head, srnode)
		linkednode := linkNodes(srnode, newnode)
		bh.insert(linkednode)
	}
}

func (bh *BinomialHeap) PrintValue() {
	if bh.forest_head == nil {
		fmt.Print("heap is empty.")
	}

	for _, node := range nodeIterator(bh.forest_head) {
		node.print_recursive(0)
	}
}

const (
	PRINT_LEVEL_INCR = 4
)

type BinomialHeapNode struct {
	value int

	parent        *BinomialHeapNode
	children_head *BinomialHeapNode
	rightsibling  *BinomialHeapNode

	order int
}

func newBinomialHeapNode(value int) *BinomialHeapNode {
	return &BinomialHeapNode{
		value:         value,
		parent:        nil,
		children_head: nil,
		rightsibling:  nil,
		order:         0,
	}
}

func (bn *BinomialHeapNode) adopt(other *BinomialHeapNode) {

	insertIntoLinkedList(&bn.children_head, other)

	// Parent relations
	other.parent = bn
}

func linkNodes(n1 *BinomialHeapNode, n2 *BinomialHeapNode) *BinomialHeapNode {
	if n1.value < n2.value {
		n1.order += 1
		n1.adopt(n2)
		return n1
	} else {
		n2.order += 1
		n2.adopt(n1)
		return n2
	}
}

func (bn *BinomialHeapNode) print_single() {

	fmt.Printf("Value: %d Order: %d\n", bn.value, bn.order)
}

func (bn *BinomialHeapNode) print_recursive(level int) {
	printSpaces(level)
	bn.print_single()

	for _, child := range nodeIterator(bn.children_head) {
		child.print_recursive(level + PRINT_LEVEL_INCR)
	}
}

const (
	INIT_ARRAYSIZE = 4
)

func insertIntoLinkedList(head **BinomialHeapNode, node *BinomialHeapNode) {
	var prev *BinomialHeapNode
	var next *BinomialHeapNode

	prev = nil
	next = *head

	for next != nil && node.order < next.order {
		prev = next
		next = next.rightsibling
	}

	if prev == nil && next == nil { // linked list is empty
		*head = node
	} else if prev == nil && next != nil { // linked list is not empty and our new node has higher rank than the node pointed by head.
		node.rightsibling = *head
		*head = node
	} else if prev != nil && next == nil { // We got to the end of the list, and our newnode has the smallest rank.
		prev.rightsibling = node
	} else if prev != nil && next != nil { // our node has found a place for itself somewhere in the list.
		prev.rightsibling = node
		node.rightsibling = next
	}
}

func removeFromLinkedList(head **BinomialHeapNode, node *BinomialHeapNode) {
	// Assume the node is present in the list.
	leftsib := getLeftsibling(*head, node)

	if leftsib == nil {
		// We are removing the head of this list.
		*head = node.rightsibling // this can set to nil.
	} else {
		leftsib.rightsibling = node.rightsibling
	}
	node.rightsibling = nil
}

func getLeftsibling(head *BinomialHeapNode, node *BinomialHeapNode) *BinomialHeapNode {
	// Assume the node is present in the list.

	if head == node {
		return nil
	}

	checknode := head

	for checknode.rightsibling != node {
		checknode = checknode.rightsibling
	}

	return checknode
}

func getNodeWithOrder(head *BinomialHeapNode, order int) *BinomialHeapNode {
	checknode := head

	for checknode != nil {
		if checknode.order == order {
			return checknode
		}
		checknode = checknode.rightsibling
	}
	return nil
}

func getMinimumNode(head *BinomialHeapNode) *BinomialHeapNode {
	// Assume there exists at least 1 node.
	minnode := head
	checknode := head.rightsibling

	for checknode != nil {
		if checknode.value < minnode.value {
			minnode = checknode
		}
		checknode = checknode.rightsibling
	}
	return minnode
}

func nodeIterator(head *BinomialHeapNode) []*BinomialHeapNode {
	arr := make([]*BinomialHeapNode, 0, INIT_ARRAYSIZE)

	trnode := head
	for trnode != nil {
		arr = append(arr, trnode)
		trnode = trnode.rightsibling
	}
	return arr
}

func printSpaces(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Print(" ")
	}
}
