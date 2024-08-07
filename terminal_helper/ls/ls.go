package ls


import (
    "os/exec"
    "fmt"
)

func LsCommand() {
    cmd := exec.Command("ls")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error : ", err)
        return
    }
    fmt.Println(string(output))
}

