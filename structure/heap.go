/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-24 13:04:34
 * @LastEditTime: 2019-11-24 13:04:44
 */
package structure

func InsertHeap(h *[]int, x int) {
	hLen := len(*h)
	halfHeapLen := int(hLen/2)
	if halfHeapLen > 0 {
		if x < (*h)[halfHeapLen] {
			tmp := (*h)[halfHeapLen]
			(*h)[halfHeapLen] = x
			x = tmp
		}
	}

	*h = append(*h, x)	
	s := (*h)[0:halfHeapLen]
	if len(s) > 0 {
		InsertHeap(&s, (*h)[halfHeapLen])
	}
}

func DeleteHeap(h *[]int) int {
	hLen := len(*h)
	var i, child int
	frist := (*h)[1]
	lastInt := (*h)[hLen-1]
	for i = 1; i * 2 < hLen; i = child {
		child = i * 2
		if child + 1 < hLen && (*h)[child] > (*h)[child + 1] {
			child++
		}

		if lastInt > (*h)[child] {
			(*h)[i] = (*h)[child]
		} else {
			break
		}
	}
	
	(*h)[i] = lastInt
	*h = (*h)[:hLen-1]

	return frist 
}

func SortHeap(h []int) []int {
	var sort []int
	l := len(h)
	for i := 0; i < l-1; i++ {
		e := DeleteHeap(&h)
		sort = append(sort, e)
	}
	return sort
} 