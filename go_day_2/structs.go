package main

import (
	"fmt"
"example.com/day_3/User"
)


func main() {
    userFirstName := getUserData("Please enter your first name: ")
    userLastName := getUserData("Please enter your last name: ")
    userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    // ... do something awesome with that gathered data!

    fmt.Println(userFirstName, userLastName, userBirthdate)

    var appUser *User.User
    appUser, err := User.NewUser(userFirstName, userLastName, userBirthdate)
    if err != nil {
        fmt.Print(err)
    }

    appUser.OutputUserDetails()
    appUser.ClearUserName()
    appUser.OutputUserDetails()
}


func getUserData(promptText string) string {
    fmt.Print(promptText)
    var value string
    fmt.Scanln(&value)
    return value
}
