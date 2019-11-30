/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-24 13:04:34
 * @LastEditTime: 2019-11-30 12:05:34
 */
package structure

type Heap struct {
	heap []int
	children []int
}

func NewHeap() *Heap {
	h := new(Heap)
	h.heap = append(h.heap, 0)
	h.children = h.heap
	return h
}

func (h *Heap) Insert(x int) {
	h.heap = append(h.heap, x)
	size := h.GetSize()
	halfOffset := int(size/2)
	j := size-1
	for i := halfOffset; i >= 1; i = int(i/2) {
		if h.heap[j] < h.heap[i] {
			tmep := h.heap[i]
			h.heap[i] = h.heap[j]
			h.heap[j] = tmep
		}
		j = i
	} 
}

func (h *Heap) GetSize() int {
	return len(h.heap)
}

func (h *Heap) GetHeap() []int {
	return h.heap
}

func (h *Heap) Delete() int {
	size := h.GetSize()
	//空时 返回-1
	if size == 0 {
		return -1
	}

	min := h.heap[1]//弹出第一个
	last := h.heap[size-1] //最后一个
	child := 0;
	
	//判断哪个儿子最小，补上第一个
	for i := 1; i*2 < size; i = child {
		child = i *2 
		if child + 1 < size && h.heap[child] > h.heap[child + 1] {
			child++
		}

		h.heap[i] = h.heap[child]
	}

	h.heap[child] = last //把最后一个元素， 放到最后挪动的那个节点
	h.heap = h.heap[0:size-1] //出堆
	return min
}

func (h *Heap) sort() {
	size := h.GetSize()
	for i := 1; i < size-1; i++ {
		min := h.Delete()
		h.heap = append(h.heap, min)
	}
}
