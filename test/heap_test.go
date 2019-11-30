/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-23 11:13:45
 * @LastEditTime: 2019-11-30 12:06:44
 */
package test

import (
	// "fmt"
	"testing"
	"go-algorithm/structure"
)

func TestInsert(t *testing.T) {
	// array := []int{90, 12, 34, 333, 1111, 1, 3, 9, 2, 8, 110, 34, 53, 97, 89, 200}
	h := structure.NewHeap()
	h.Insert(90)
	t.Log(h.GetHeap())
	h.Insert(12)
	t.Log(h.GetHeap())
	h.Insert(34)
	t.Log(h.GetHeap())
	h.Insert(333)
	t.Log(h.GetHeap())
	h.Insert(1)
	t.Log(h.GetHeap())
	h.Insert(3)
	t.Log(h.GetHeap())
	h.Insert(9)
	t.Log(h.GetHeap())
	h.Insert(2)
	t.Log(h.GetHeap())
	h.Insert(110)
	t.Log(h.GetHeap())
	
	min := h.Delete()
	t.Log(min ,h.GetHeap())
	min = h.Delete()
	t.Log(min ,h.GetHeap())
	min = h.Delete()
	t.Log(min ,h.GetHeap())
}

