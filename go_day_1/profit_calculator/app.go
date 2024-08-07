package main

import (
    "errors"
    "fmt"
    "os"
)

func writeValuesToFile(ebt, profit, ratio float64) error{

    data := fmt.Sprintf("EBT : %.2f, Profit is %.2f, and the Ratio is %.2f", ebt, profit, ratio)
    err :=  os.WriteFile("profits.txt", []byte(data), 0644)

    if err != nil {
        fmt.Println("Error occured")
        fmt.Print(err)
    }
    return nil
}

func main() {

    var revenue float64
    var expenses float64
    var taxRate float64

    // fmt.Print("Enter your revenue: ")
    // fmt.Scan(&revenue)


    // fmt.Print("Enter the expenses: ")
    // fmt.Scan(&expenses)
    //
    // fmt.Print("Enter the tax rate: ")
    // fmt.Scan(&taxRate)

    inputRevenue(&revenue)
    inputExpenses(&expenses)
    inputTaxRate(&taxRate)

    earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

    fmt.Printf("Here is your EBT : %.2f\n", earningsBeforeTax)
    fmt.Printf("Here is your profit : %.2f\n", profit)
    fmt.Printf("Here is your ratio : %.2f\n", ratio)

    err := writeValuesToFile(earningsBeforeTax, profit, ratio)
    if err != nil{
        err := errors.New("Coudnt write the values to the file")
        fmt.Println(err)
    }

}

func inputRevenue(revenue *float64)error {

    fmt.Print("Enter the revenue: ")
    fmt.Scan(revenue)
    if *revenue < 0.0 {
        err := errors.New("Wrong input received")
        return err
    }
    return nil
}

func inputExpenses(expenses *float64) error {

    fmt.Print("Enter the expenses: ")
    fmt.Scan(expenses)
    if *expenses < 0.0 {
        err := errors.New("Invalid data")
        return err
    }
    return nil
}

func inputTaxRate(taxRate *float64) error {

    fmt.Print("Enter the tax rate: ")
    fmt.Scan(taxRate)
    if *taxRate <= 0.0 {
        err := errors.New("Invalid data")
        return err
    }
    return nil
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {

    earningsBeforeTax := revenue - expenses
    profit := earningsBeforeTax - (earningsBeforeTax * taxRate / 100)
    ratio := earningsBeforeTax / profit

    return earningsBeforeTax, profit, ratio

}
