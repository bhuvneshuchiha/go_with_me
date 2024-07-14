package main

import "fmt"

func main() {

    // website := []string{"https://google.com", "https://amazon.com"}
    website := map[string]string{

        "Google" : "https://google.com",
        "AWS" : "https://aws.com",

    }
    fmt.Println(website)
    fmt.Println(website["Google"])


}
