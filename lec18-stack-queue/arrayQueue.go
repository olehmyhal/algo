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
	a.queue[a.tail] = x
	a.tail = (a.tail + 1) % a.capacity
	a.n += 1
}

func (a *ArrayQueue) Dequeue()int{
	val := a.queue[a.head]

	a.head = (a.head + 1) % a.capacity
	a.n -= 1
	return val
}

func (a *ArrayQueue) Size()int{
	if a.head > a.tail {
		return a.capacity - a.head + a.tail
	}

	return a.tail - a.head
}

func (a *ArrayQueue) Reserve(cap int){
	newQueue := make([]int, a.capacity, a.capacity)

	for i := 0; i < a.n; i++ {
		newQueue[i] = a.queue[(a.head + i) % a.capacity]
	}

	a.queue = newQueue
	a.head = 0
	a.tail = a.n
}