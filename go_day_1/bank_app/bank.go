package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func getBalanceFromFile() (float64, error){

    data, err := os.ReadFile(fileName)

    if err != nil {
        return 1000, errors.New("Failed to find the balance file.")
    }

    balanceText := string(data)
    balance, err := strconv.ParseFloat(balanceText, 64)

    if err != nil {
        return 1000, errors.New("Failed to convert the balance to float64")
    }

    return balance,nil
}

func writeBalanceToFile(balance float64) {

    balanceText := fmt.Sprint(balance)
    os.WriteFile(fileName, []byte(balanceText), 0644)

}

func main() {

    balance , error := getBalanceFromFile()
    if error != nil {
        fmt.Print("Something went wrong!\n")
        fmt.Println(error)
        panic("Sorry cant continue due to an error")
    }

    exit := false
    var choice int
    fmt.Print("Welcome to GO BANK\n")

    for exit == false{

        fmt.Print("What would you like to do?\n")
        fmt.Print("1.Check balance:- \n")
        fmt.Print("2.Deposit money:- \n")
        fmt.Print("3.Withdraw money:- \n")
        fmt.Print("4.Exit!\n")
        fmt.Print("Your choice: \n")
        fmt.Scan(&choice)
        fmt.Printf("Your choice -> %v\n", choice)

        if choice == 1 {
            fmt.Printf("Your balance is: %v\n", balance)
        } else if choice == 2 {
            fmt.Printf("How much do you want to deposit? \n")
            moneyDeposit := 0.0
            fmt.Scan(&moneyDeposit)
            if moneyDeposit <= 0{
                fmt.Print("Invalid amount entered ")
                continue
            }
            balance = balance + moneyDeposit
            writeBalanceToFile(balance)
        } else if choice == 3 {
            fmt.Print("How much you want to withdraw? \n")
            moneyToWithdraw := 0.0
            fmt.Scan(&moneyToWithdraw)

            if moneyToWithdraw > balance {
                fmt.Print("Not sufficient balance \n")
            } else{
                balance = balance - moneyToWithdraw
                writeBalanceToFile(balance)
                fmt.Print("Successfully withdrawn !\n")
            }
        }else if choice == 4 {
            fmt.Print("Goodbye!")
            exit = true
        }
    }



}
