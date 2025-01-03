package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bruh5(){
	reader := bufio.NewReader(os.Stdin)
	
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	stack := ArrayStack{}
	stack.Init()

	inputSlice := strings.Split(input, " ")

	Loop:
	for _, item := range inputSlice {
		num, err := strconv.Atoi(item)

		if err == nil {
			stack.Push(num)
			continue Loop
		}

		firstInt := stack.Pop()
		secondInt := stack.Pop()

		if item == "*" {
			stack.Push(firstInt * secondInt)
		} else if item == "/" {
			stack.Push(secondInt / firstInt)
		} else if item == "+" {
			stack.Push(firstInt + secondInt)
		} else if item == "-" {
			stack.Push(secondInt - firstInt)
		}
	}

	res := stack.Pop()
	fmt.Printf("%v", res)
}

type ArrayStack struct {
	stack []int
	size int
	capacity int
}

func (a *ArrayStack) Init(){
	a.capacity = 4
	a.stack = make([]int, 0, a.capacity)
}

func (a *ArrayStack) GetLast()int {
	return a.stack[a.size - 1]
}

func (a *ArrayStack) Push(x int) {
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	a.stack = append(a.stack, x)
	a.size += 1
}

func (a *ArrayStack) Pop()int {
	if a.size < a.capacity / 4 {
		a.resize(a.capacity / 2)
	}

	deleted := a.stack[a.size - 1]
	a.stack = a.stack[:a.size - 1]
	a.size -= 1

	return deleted
}

func (a *ArrayStack) resize(cap int) {
	newStack := make([]int, a.size, cap)

	for i, item := range a.stack {
		newStack[i] = item
	}

	a.stack = newStack
	a.capacity = cap
} 