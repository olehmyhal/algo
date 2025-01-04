package tasks

import (
	"fmt"
	"strings"
)

func bruh7(){
	var path string

	fmt.Scanln(&path)

	stack := ArrayStack2{}
	stack.Init()

	stack.Push("/")

	pathSlice := strings.Split(path, "/")

	Loop:
	for _, pathPart := range pathSlice {
		if pathPart == "" {
			continue Loop
		}

		if pathPart == ".." {
			stack.Pop()
			continue Loop
		}

		if pathPart == "." {
			continue Loop
		}

		stack.Push(pathPart)
	}

	joinedPath := stack.stack[0] + strings.Join(stack.stack[1:stack.size], "/")

	fmt.Printf("%v", joinedPath)
}

type ArrayStack2 struct {
	stack []string
	size int
	capacity int
}

func (a *ArrayStack2) Init(){
	a.capacity = 4
	a.stack = make([]string, 0, a.capacity)
}

func (a *ArrayStack2) GetLast()string {
	return a.stack[a.size - 1]
}

func (a *ArrayStack2) Push(x string) {
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	a.stack = append(a.stack, x)
	a.size += 1
}

func (a *ArrayStack2) Pop() {
	if a.size == 1 {
		return
	}

	if a.size < a.capacity / 4 {
		a.resize(a.capacity / 2)
	}

	a.stack = a.stack[:a.size - 1]
	a.size -= 1

}

func (a *ArrayStack2) resize(cap int) {
	newStack := make([]string, a.size, cap)

	for i, item := range a.stack {
		newStack[i] = item
	}

	a.stack = newStack
	a.capacity = cap
} 