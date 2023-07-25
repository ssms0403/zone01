package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

type roman struct {
	num        int
	romanDigit string
}

func main() {
	// Parse the command-line argument as an integer
	nbr, err := strconv.Atoi(os.Args[1])

	// Write the error message in case of an invalid number
	if err != nil || nbr >= 4000 || nbr == 0 {
		printStr("ERROR: cannot convert to roman digit")
		return
	}

	// Initialize an array of Roman numeral strings
	romanNumerals := []roman{
		{num: 1000, romanDigit: "M"},
		{num: 900, romanDigit: "CM"},
		{num: 500, romanDigit: "D"},
		{num: 400, romanDigit: "CD"},
		{num: 100, romanDigit: "C"},
		{num: 90, romanDigit: "XC"},
		{num: 50, romanDigit: "L"},
		{num: 40, romanDigit: "XL"},
		{num: 10, romanDigit: "X"},
		{num: 9, romanDigit: "IX"},
		{num: 5, romanDigit: "V"},
		{num: 4, romanDigit: "IV"},
		{num: 1, romanDigit: "I"},
	}

	// Initialize a slice to hold the Roman numeral characters
	chars := make([]string, 0)

	// Initialize a variable to store the remaining value to be converted
	remaining := nbr

	// Loop through the map of integer values to Roman numeral strings
	for _, r := range romanNumerals {
		// Divide the remaining value by the current integer value
		count := remaining / r.num
		if count == 0 {
			continue
		}

		// Concatenate the corresponding Roman numeral characters to the slice
		roman := r.romanDigit

		// Concatenate the corresponding Roman numeral characters to the slice
		for i := 0; i < count; i++ {
			chars = append(chars, roman)
		}

		// Update the remaining value
		remaining %= r.num

		// Break the loop if the remaining value is equal to 0
		if remaining == 0 {
			break
		}
	}

	// Join the slice of Roman numeral characters to create the Roman numeral string
	roman := Join(chars, "")

	// Join the slice of Roman numeral characters to create the Roman numeral string
	for i, digit := range chars {
		if len(digit) == 2 {
			chars[i] = "(" + string(digit[1]) + "-" + string(digit[0]) + ")"
		}
	}
	sumRoman := Join(chars, "+")

	// Print the Roman numeral and the calculation used to obtain it
	printStr(sumRoman + "\n" + roman)
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Join(arrString []string, sep string) string {
	result := ""
	for i, str := range arrString {
		result += str
		if i != len(arrString)-1 {
			result += sep
		}
	}
	return result
}
