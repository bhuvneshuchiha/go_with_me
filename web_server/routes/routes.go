package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Router() {
    resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=8f2a34503d3a4b77973114019241108&q=london")
    if err != nil {
        log.Println("error while request sending")
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    body, er := io.ReadAll(resp.Body)
    if er != nil {
        fmt.Println(er)
        return
    }
    fmt.Println(string(body))

    // @Unmarshalling into a map for structured handling of the response
    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        fmt.Println("Error in Unmarshalling ", err)
        return
    }

    //Marshall it back for printing
    content, err := json.MarshalIndent(result, "", " ")
    if err != nil {
        fmt.Println("Error in Marshalling ", err)
        return
    }
    fmt.Println(string(content))
}














