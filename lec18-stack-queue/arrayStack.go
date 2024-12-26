package lec18stackqueue

type ArrayStack struct {
	stack []int
	size int
	capacity int
}

func (a *ArrayStack) Init(){
	a.capacity = 4
	a.stack = make([]int, a.capacity, a.capacity)
}

func (a *ArrayStack) Push(x int) {
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	a.stack[a.size] = x
	a.size += 1
}

func (a *ArrayStack) Pop() {
	if a.size < a.capacity / 4 {
		a.resize(a.capacity / 2)
	}

	a.stack = a.stack[:a.size - 1]
	a.size -= 1
}

func (a *ArrayStack) resize(cap int) {
	newStack := make([]int, cap, cap)

	for i, item := range a.stack {
		newStack[i] = item
	}

	a.stack = newStack
	a.capacity = cap
} 