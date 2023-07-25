package main

import "fmt"

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	var jack Pilot
	jack.Name = "Donnie"
	jack.Life = 100.0
	jack.Age = 24
	jack.Aircraft = AIRCRAFT1

	fmt.Println(jack)
}

const AIRCRAFT1 = 1
