package lec13heap

import "fmt"

func TestHeap() {
	h := Heap{items: []float64{12, 20, 15, 29, 27, 17, 22}}
	h.size = 7
	fmt.Printf("1. %v\n", h.items)

	h.ExtractMin()
	fmt.Printf("2. %v\n", h.items)

	h.Insert(1)
	fmt.Printf("3. %v\n", h.items)

	h.Insert(145)
	fmt.Printf("4. %v\n", h.items)

	h.ExtractMin()
	fmt.Printf("5. %v\n", h.items)
}

type Heap struct {
	items []float64
	size  int
}

func (h *Heap) ExtractMin() float64 {
	if h.size == 0 {
		panic("Heap is empty")
	}

	min := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.items = h.items[:h.size]

	h.siftDown(0)

	return min
}

func (h *Heap) Insert(n float64) {
	h.items = append(h.items, n)
	h.size++
	h.siftUp(h.size - 1)
}

func (h *Heap) siftUp(index int) {
	for index > 0 && h.items[index] < h.items[(index-1)/2] {
		h.items[index], h.items[(index-1)/2] = h.items[(index-1)/2], h.items[index]
		index = (index - 1) / 2
	}
}

func (h *Heap) siftDown(index int) {
	for 2*index+1 < h.size {
		j := 2*index + 1 

		if j+1 < h.size && h.items[j+1] < h.items[j] {
			j = j + 1 
		}

		if h.items[index] <= h.items[j] {
			break
		}

		h.items[index], h.items[j] = h.items[j], h.items[index]
		index = j
	}
}