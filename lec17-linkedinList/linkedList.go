package lec17linkedinlist

type Node struct {
	 next *Node
	 val string
}

type LinkedList struct {
	size int
	sentinel *Node
}

func(l *LinkedList) Init(){
	l.sentinel = &Node{
		val: "anything",
	}
}

func(l *LinkedList) InsertFirst(val string) {
	l.sentinel.next = &Node{
		val: val,
		next: l.sentinel.next,
	}
	l.size++
}

func(l *LinkedList) InsertLast(val string) {
	craw := l.sentinel 

	for craw.next != nil {
		craw = craw.next
	}
	craw.next = &Node{
		val: val,
	}
	l.size++
}