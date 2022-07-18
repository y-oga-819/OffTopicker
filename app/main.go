package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

var unused int // ERROR "`unused` is unused"

func nouse() {
	x := math.MinInt8
	for {
		_ = x
		x = 0 // ERROR "ineffectual assignment to x"
		x = 0
	}
}

func main() {
	fmt.Println(quote.Hello())
}
