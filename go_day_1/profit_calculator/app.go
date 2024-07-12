package main

import (
    "fmt"
)


func main() {

    var revenue float64
    var expenses float64
    var taxRate float64

    fmt.Print("Enter your revenue: ")
    fmt.Scan(&revenue)

    fmt.Print("Enter the expenses: ")
    fmt.Scan(&expenses)

    fmt.Print("Enter the tax rate: ")
    fmt.Scan(&taxRate)

    earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

    fmt.Printf("Here is your EBT : %.2f\n", earningsBeforeTax)
    fmt.Printf("Here is your profit : %.2f\n", profit)
    fmt.Printf("Here is your ratio : %.2f\n", ratio)

}


func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {

    earningsBeforeTax := revenue - expenses
    profit := earningsBeforeTax - (earningsBeforeTax * taxRate / 100)
    ratio := earningsBeforeTax / profit

    return earningsBeforeTax, profit, ratio

}
