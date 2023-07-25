package main

import (
	"fmt"
	"os"
	"strconv"
)

func SplitWhiteSpaces(s string) []string {
	arr := []string(nil)
	lastWhiteSpaceIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if lastWhiteSpaceIndex != i {
				arr = append(arr, s[lastWhiteSpaceIndex:i])
			}
			lastWhiteSpaceIndex = i + 1
		} else if i == len(s)-1 {
			arr = append(arr, s[lastWhiteSpaceIndex:i+1])
			lastWhiteSpaceIndex = i + 1
		}
	}
	return arr
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	var values []int
	expression := SplitWhiteSpaces(os.Args[1])

	for _, operator := range expression {
		val, err := strconv.Atoi(operator)
		if err == nil {
			values = append(values, val)
			continue
		}

		n := len(values)
		if n < 2 {
			fmt.Println("Error")
			return
		}

		switch operator {
		case "+":
			values[n-2] += values[n-1]
			values = values[:n-1]
		case "-":
			values[n-2] -= values[n-1]
			values = values[:n-1]
		case "*":
			values[n-2] *= values[n-1]
			values = values[:n-1]
		case "/":
			values[n-2] /= values[n-1]
			values = values[:n-1]
		case "%":
			values[n-2] %= values[n-1]
			values = values[:n-1]
		default:
			fmt.Println("Error")
			return
		}
	}
	if len(values) == 1 {
		fmt.Println(values[0])
	} else {
		fmt.Println("Error")
	}
}
