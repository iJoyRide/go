package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64 = 1000 //explicitly declare variables as float64
	expectedReturnRate := 5.5           // or use .0 for literals
	years := 10.0                       // You used the := shorthand which is concise and idiomatic in Go

	fmt.Print("Enter the investment amount: ")

	fmt.Scan(&investmentAmount)
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
