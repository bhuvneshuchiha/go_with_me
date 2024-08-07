package User

import (
    "fmt"
    "time"
    "errors"
)


type User struct {
    firstName string
    lastName string
    birthDate string
    createdAt time.Time
}

func NewUser(firstName, lastName, birthDate string) (*User, error){
    if firstName == "" || lastName == "" || birthDate == "" {
        return nil, errors.New("Fields should not be empty")
    }
    return &User{
        firstName: firstName,
        lastName: lastName,
        birthDate: birthDate,
    }, nil
}
//Attaching this to a struct, its a method now.
func (u User) OutputUserDetails() {
    fmt.Println(u.firstName, u.lastName, u.birthDate)
}

//Another method with the receiver argument
func (u *User) ClearUserName() {
    u.firstName = ""
    u.lastName = ""
}
