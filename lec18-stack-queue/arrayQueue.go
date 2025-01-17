package lec18stackqueue

type ArrayQueue struct {
	queue []int
	n int
	capacity int 
	head int
	tail int
}

func (a *ArrayQueue) Init(){
	a.capacity = 8
	a.queue = make([]int, a.capacity, a.capacity)

	a.head, a.tail, a.n = 0, 0, 0
}

func (a *ArrayQueue) Enqueue(x int){
	if a.n == a.capacity {
		a.Reserve(2 * a.capacity)
	}
	a.queue[a.tail] = x
	a.tail = (a.tail + 1) % a.capacity
	a.n += 1
}

func (a *ArrayQueue) Dequeue()int{
	val := a.queue[a.head]

	a.head = (a.head + 1) % a.capacity
	a.n -= 1

	if a.n > 0 && a.n <= a.capacity/4 {
		a.Reserve(a.capacity / 2)
	}

	return val
}

func (a *ArrayQueue) Size()int{
	if a.head > a.tail {
		return a.capacity - a.head + a.tail
	}

	if a.head == a.tail {
		return a.capacity
	}

	return a.tail - a.head
}

func (a *ArrayQueue) Reserve(cap int){
	newQueue := make([]int, cap, cap)

	for i := 0; i < a.n; i++ {
		newQueue[i] = a.queue[(a.head + i) % a.capacity]
	}

	a.queue = newQueue
	a.capacity = cap
	a.head = 0
	a.tail = a.n
}