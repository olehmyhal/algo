package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bruh8(){
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	var q, n int
	fmt.Sscanf(scanner.Text(), "%d %d", &q, &n)
	
	commands := make([]string, q)
	
	for i := 0; i < q; i++ {
		scanner.Scan()
		commands[i] = scanner.Text()
	}

	res := []int{}
	unreserved := []int{}

	for _, command := range commands {
		parts := strings.Split(command, " ")
		com := parts[0]

		if com == "reserve" {
			if len(unreserved) > 0 {
				seat := unreserved[len(unreserved) - 1]
				unreserved = unreserved[:len(unreserved) - 1]
				res = append(res, seat)
				continue
			}

			if len(res) < 1 {
				res = append(res, 1)
			} else {
				res = append(res, res[len(res) - 1] + 1)
			}
		} else if com == "unreserve" {
			seatNumber, _ := strconv.Atoi(parts[1])
			unreserved = insertElem(unreserved, seatNumber)
		}
	}

	for _, r := range res {
		fmt.Printf("%v\n", r)
	}
}

func insertElem(availableSeats []int, seatNumber int) []int {
	availableSeats = append(availableSeats, seatNumber)
	sliceSize := len(availableSeats)

	for j := sliceSize - 2; j >= 0; j-- {
		if availableSeats[j] < availableSeats[j+1] {
			availableSeats[j+1] = availableSeats[j]
		} else {
			availableSeats[j+1] = seatNumber
			break
		}
	}

	if availableSeats[0] < seatNumber {
		availableSeats[0] = seatNumber
	}

	return availableSeats
}