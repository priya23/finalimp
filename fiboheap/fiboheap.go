package fiboheap

import (
	"fmt"
)

type FibHeap struct {
	roots       *Fibtree
	nodes       int
	degree      int
	numoftress  int
	minheap     *Fibtree
	numofmarked int
}

type Fibtree struct {
	Parent, Child, next, prev *Fibtree
	Degree                    int
	Mark                      bool
	key                       int
	value                     string
}

func CreateNewHeap() *FibHeap {
	heap := new(FibHeap)
	heap.roots = nil
	heap.nodes = 0
	heap.degree = 0
	heap.minheap = nil
	heap.numoftress = 0
	heap.numofmarked = 0
	return heap
}

func CreateSingeltonTree(ket int, val string) *Fibtree {
	tree1 := new(Fibtree)
	tree1.Degree = 0
	tree1.Parent = nil
	tree1.Child = nil
	tree1.next = tree1
	tree1.prev = tree1
	tree1.Mark = false
	tree1.key = ket
	tree1.value = val
	return tree1
}
func (fb *FibHeap) Give(val string, pro int) {
	fb.Insert(pro, val)
}
func (fb *FibHeap) Take() int {
	return fb.DeleteMin()
}
func (fb *FibHeap) Insert(ket int, val string) {
	tree := CreateSingeltonTree(ket, val)
	fb.minheap = mergeintolist(fb.minheap, tree)
	fb.nodes += 1
}

func mergeintolist(one, two *Fibtree) *Fibtree {
	if one == nil && two == nil {
		return nil
	} else if one == nil && two != nil {
		return two
	} else if one != nil && two == nil {
		return one
	} else {
		oneNext := one.next
		one.next = two.next
		one.next.prev = one
		two.next = oneNext
		two.next.prev = two
		//fmt.Println("one is ", one.key)
		//fmt.Println("two is", two.key)
		if one.key < two.key {
			//fmt.Println("inside one")
			return one
		}
		return two
	}
}
func printSpaces(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Print(" ")
	}
}
func PrintChild(val int) {
	printSpaces(val)
}
func (fb *FibHeap) PrintRoot() {
	k := fb.minheap
	y := k.next
	//fmt.Println("root is ", k)
	for y != k {
		//fmt.Println("root is ", y)
		y = y.next
	}
}

func (fb *FibHeap) DeleteMin() int {
	fb.nodes -= 1
	min := fb.minheap
	if min.next == min {
		fb.minheap = nil
	} else {
		fb.minheap.prev.next = fb.minheap.next
		fb.minheap.next.prev = fb.minheap.prev
		fb.minheap = fb.minheap.next //assign arbitary
	}
	if min.Child != nil {
		curr := min.Child
		for curr != min.Child {
			curr.Parent = nil
			curr = curr.next
		}
	}
	fb.minheap = mergeintolist(fb.minheap, min.Child)
	if fb.minheap == nil {
		return min.key
	}
	treeSlice := make([]*Fibtree, 0, fb.nodes)
	toVisit := make([]*Fibtree, 0, fb.nodes)

	for curr := fb.minheap; len(toVisit) == 0 || toVisit[0] != curr; curr = curr.next {
		toVisit = append(toVisit, curr)
	}
	for _, curr := range toVisit {
		for {
			for curr.Degree >= len(treeSlice) {
				treeSlice = append(treeSlice, nil)
			}
			if treeSlice[curr.Degree] == nil {
				treeSlice[curr.Degree] = curr
				break
			}

			other := treeSlice[curr.Degree]
			treeSlice[curr.Degree] = nil
			var minT, maxT *Fibtree
			if other.key < curr.key {
				minT = other
				maxT = curr
			} else {
				minT = curr
				maxT = other
			}

			//rmeove from root
			maxT.next.prev = maxT.prev
			maxT.prev.next = maxT.next

			maxT.next = maxT
			maxT.prev = maxT
			maxT.Parent = minT
			maxT.Mark = false
			minT.Degree += 1
			minT.Child = maxT
			curr = minT
		}
		if curr.key <= fb.minheap.key {
			fb.minheap = curr
		}
	}
	return min.key
}
func (fb1 *FibHeap) Merge(fb2 *FibHeap) *FibHeap {
	if fb1 == nil || fb2 == nil {
		return nil
	}
	newsixe := fb1.nodes + fb2.nodes
	newheap := mergeintolist(fb1.minheap, fb2.minheap)
	fb1.minheap = nil
	fb2.minheap = nil
	fb1.nodes = 0
	fb2.nodes = 0
	nn := CreateNewHeap()
	nn.minheap = newheap
	nn.nodes = newsixe
	return nn
}
func (fb *FibHeap) PrintMin() {
	fmt.Println("MIN IS ", fb.minheap)
}
