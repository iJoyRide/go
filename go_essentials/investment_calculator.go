package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000 //explicitly declare variables as float64
	var expectedReturnRate = 5.5        // or use .0 for literals
	years := 10.0                       // You used the := shorthand which is concise and idiomatic in Go

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}
