package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		regEx := os.Args[1]
		text := os.Args[2]
		if validateRegEx(regEx) {
			searchWords := Split(regEx[1:len(regEx)-1], []rune{'|'})
			textWords := Split(text, []rune{' ', ',', '\''})
			foundIndex := 1
			for _, word := range textWords {
				for _, search := range searchWords {
					if index(word, search) != -1 {
						fmt.Println(foundIndex, ": ", word)
						foundIndex++
						break
					}
				}
			}
		}
	}
}

func validateRegEx(regEx string) bool {
	if regEx[0] != '(' && regEx[len(regEx)-1] != ')' {
		return false
	}
	return true
}

func index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		subs := s[i : i+len(toFind)]
		if subs == toFind {
			return i
		}
	}
	return -1
}

func Split(s string, seps []rune) []string {
	arr := []string(nil)
	lastWhiteSpaceIndex := 0
	for i := 0; i < len(s); i++ {
		isSeparator := false
		for _, sep := range seps {
			if rune(s[i]) == sep {
				isSeparator = true
				break
			}
		}
		if isSeparator {
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
