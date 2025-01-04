package tasks

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func bruh6() {
	var k, x int
	fmt.Scanf("%d %d\n", &k, &x)

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')

	line = strings.TrimSpace(line)

	strArray := strings.Split(line, " ")

	var arr []int
	for _, str := range strArray {
		num, _ := strconv.Atoi(str)
		arr = append(arr, num)
	}

	sort.Slice(arr, func(i, j int)bool {
		diffI := int(math.Abs(float64(arr[i] - x)))
		diffJ :=  int(math.Abs(float64(arr[j] - x)))

		if diffI == diffJ {
			return i < j
		}

		return diffI < diffJ
	})

	newArr := arr[:k]

	sort.Ints(newArr)

	for _, item := range newArr {
		fmt.Printf("%v ", item)
	}
}
