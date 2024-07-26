package shortenurl

import (
    "fmt"
    "math/rand"
)

func Shortenurl (url string) string{

    s := ""
    for i := 0; i < 6; i++ {
        s += string(rand.Intn(25) + 97)
    }

    shortenedUrl := fmt.Sprintf("http://localhost:8080/%s", s)
    return shortenedUrl
}
