package main

import "fmt"

func main() {
	word := "ABC뎶"
	fmt.Println(len(word))
	fmt.Println("---------------------------------------------")
	for _, elem := range word {
		fmt.Printf("Char : %s | Value : %v | Type: %T\n", string(elem), elem, elem)
	}
	fmt.Println("---------------------------------------------")
	for i := 0; i < len(word); i++ {
		elem := word[i]
		fmt.Printf("Char : %s | Value : %v | Type: %T\n", string(elem), elem, elem)
	}
}
