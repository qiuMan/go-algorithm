/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-23 11:13:45
 * @LastEditTime: 2019-11-24 13:04:56
 */
package test

import (
	"fmt"
	"testing"
	"go-algorithm/structure"
)

func TestInsertHeap(t *testing.T) {
	array := []int{90, 12, 34, 333, 1111, 1, 3, 9, 2, 8, 110, 34, 53, 97, 89, 200}
	var heap []int
	heap = append(heap, -1)
	arrayLen := len(array)
	for i := 0; i < arrayLen; i++ {
		function.InsertHeap(&heap, array[i])
	}
	t.Log(heap)
}

func TestDeleteHeap(t *testing.T) {
	heap := []int{-1, 1, 2, 3, 9, 8, 34, 12, 200, 90, 1111, 110, 34, 53, 97, 89, 333}
	hLen := len(heap)
	for i := 0; i < hLen-1; i++ {
		frist := function.DeleteHeap(&heap)
		fmt.Println(fmt.Sprintf("%d", frist))
	}

}

func TestSortHeap(t *testing.T) {
	heap := []int{-1, 1, 2, 3, 9, 8, 34, 12, 200, 90, 1111, 110, 34, 53, 97, 89, 333}
	r := function.SortHeap(heap)
	t.Log(r)
}