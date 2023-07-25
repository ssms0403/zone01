package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		min, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		max, _err := strconv.Atoi(os.Args[2])
		if _err != nil {
			fmt.Println(_err)
			return
		}
		for i := min; ; {
			fmt.Print(i)
			if i != max {
				fmt.Print(" ")
			} else {
				fmt.Println()
				break
			}
			if min > max {
				i--
			} else {
				i++
			}
		}
	}
}
