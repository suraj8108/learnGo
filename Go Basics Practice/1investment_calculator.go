package main

import (
	"fmt"
	"math"
)

func investmentCalculator() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/ 100), years)
	var futureRealValue = futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	formattedStrings := fmt.Sprint("Futute Value is %v", futureValue);
	fmt.Print(formattedStrings);

	//Multiline
	multipleLines := `My Name is suraj
										And I am a Software Engineer`
	fmt.Println(multipleLines)
}