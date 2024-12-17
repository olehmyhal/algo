package lec20recursion

import (
	"fmt"
	"strconv"
	"strings"
)

func CallCalc(){
	res := Calc("2+2*2")

	fmt.Printf("%v\n", res)
}

//Calc("2+2*2")
func Calc(str string) int {
	if strings.Contains(str, "+") {
		index := strings.Index(str, "+")

		return Calc(str[:index]) + Calc(str[index + 1:])
		
		
	} else if strings.Contains(str, "*") {
		index := strings.Index(str, "*")


		return Calc(str[:index]) * Calc(str[index + 1:])
	} else {
		num, _ := strconv.Atoi(str)

		return num
	}
}