package implementheap

import (
	//"fmt"
	"testing"
)

func TestPushandPop(t *testing.T) {
	h1 := CreateHeap()
	h1.Give("priya", 4)
	h1.Give("test", 9)
	h1.Give("jdsshc", 3)
	v1 := h1.Take()
	//fmt.Println("vw", v1)
	if v1 != 9 {
		t.Error("pop not correct")
	}
	v2 := h1.Take()
	if v2 != 4 {
		t.Error("pop not correct")
	}
	v3 := h1.Take()
	if v3 != 3 {
		t.Error("pop not correct")
	}
}

func TestMerge(t *testing.T) {
	h1 := CreateHeap()
	h1.Give("priya", 4)
	h1.Give("test", 9)
	h1.Give("jdsshc", 3)
	h2 := CreateHeap()
	h2.Give("priya", 8)
	h2.Give("test", 0)
	h2.Give("jdsshc", 37)
	h := Merge(h1, h2)
	v1 := h.Take()
	if v1 != 37 {
		t.Error("merge not correct")
	}
}
