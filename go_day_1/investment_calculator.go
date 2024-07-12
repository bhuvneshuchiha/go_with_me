package main

import (
    "fmt"
    "math"
)

func main() {

    const inflationRate = 2.5
    var investmentAmount float64
    var expectedReturnRate float64
    var years float64

    //Ask for amount
    fmt.Print("Enter the investment amount please : ")
    fmt.Scan(&investmentAmount)

    //Ask for years
    fmt.Print("Enter the number of years : ")
    fmt.Scan(&years)

    //Ask for expectedReturnRate
    fmt.Print("Enter the expected return rate : ")
    fmt.Scan(&expectedReturnRate)

    futureValue := investmentAmount *
    math.Pow((1 + expectedReturnRate/100), years)

    futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

    fmt.Println(futureValue)
    fmt.Println(futureRealValue)

}
