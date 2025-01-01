package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bruh1(){
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	parts := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(parts[0])
	s, _ := strconv.Atoi(parts[1])

	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimSpace(secondLine)
	weights := strings.Split(secondLine, " ")

	thirdLine, _ := reader.ReadString('\n')
	thirdLine = strings.TrimSpace(thirdLine)
	costs := strings.Split(thirdLine, " ")

	w := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		w[i], _ = strconv.Atoi(weights[i])
		c[i], _ = strconv.Atoi(costs[i])
	}

	j := 0

	currSum, currWeight, maxSum := 0, 0, 0

	for i := 0; i < n; i++ {
		currWeight += w[i]
		currSum += c[i]

		for currWeight > s {
			currSum -= c[j]
			currWeight -= w[j]

			j++
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	fmt.Printf("%v", maxSum)
}