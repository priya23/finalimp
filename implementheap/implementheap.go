package implementheap

import (
	"container/heap"
	"encoding/json"
	"fmt"
	//"github.com/priya23/finalpq"
)

type Item struct {
	Val      interface{}
	Priority int
	Index    int
}

type PriorityQueue []*Item

//for sort implementation
func (pq PriorityQueue) Len() int {
	//fmt.Println("inside len")
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	//fmt.Println("inside less")
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	//fmt.Println("inside swap")
	//fmt.Println("I J is", i, j, pq[i].Index, pq[j].Index)
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) DecreaseKey(valu int, index int) {
	item := (*pq)[index]
	item.Priority = valu
	heap.Fix(pq, item.Index)
}

//creating a empty default priority queue
func Create() PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return pq
}

//implementing default interface
func (pl *PriorityQueue) Push(x interface{}) {
	length := len(*pl)
	item := x.(*Item)
	item.Index = length
	*pl = append(*pl, item)
	//fmt.Println("added")

}
func (pq *PriorityQueue) Pop() interface{} {
	//fmt.Println("INSIDE pop")
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

/*func (pq *PriorityQueue) printAll() {
	for i := 0; i < pq.Len(); i++ {
		fmt.Println("val is %v %v %v", (*pq)[i].Priority, (*pq)[i].Val, (*pq)[i].Index)
	}
}*/

func (k *PriorityQueue) Give(val string, pro int) {
	i := Item{Val: val, Priority: pro}
	heap.Push(k, &i)
	//k.Push(&i)
}

func (k *PriorityQueue) Take() int {
	retval := heap.Pop(k).(*Item).Priority
	return retval
}
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}

func CreateItem(val interface{}, pro int) *Item {
	k := Item{Val: val, Priority: pro}
	return &k
}
func CreateHeap() *PriorityQueue {
	p := make(PriorityQueue, 0)
	heap.Init(&p)
	return &p
}
func Merge(pq *PriorityQueue, pl *PriorityQueue) *PriorityQueue {
	k := CreateHeap()
	for _, r := range *pq {
		*k = append(*k, r)
	}
	for _, rr := range *pl {
		*k = append(*k, rr)
	}
	heap.Init(k)
	return k
}
func (pq *PriorityQueue) PrintValue() {
	//fmt.Println("length is %v", pq.Len())
	length := pq.Len()
	for i := 0; i < length; i++ {
		vv := pq.Take()
		fmt.Printf("\n %v", vv)
	}
}
