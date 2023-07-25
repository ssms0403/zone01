package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("descend-path.txt")
	slice := strings.Split(string(content), "\r\n")
	lenSlice := len(slice)
	for i := 0; i < lenSlice/2; i++ {
		slice[i], slice[lenSlice-i-1] = slice[lenSlice-i-1], slice[i]
	}
	file, _ := os.Create("all-paths.txt")
	for i := 0; i < lenSlice; i++ {
		file.WriteString(strconv.Itoa(i+1) + " + " + slice[i] + "\n")
	}
	file.Sync()
}
