package remove

import (
	"fmt"
	"log"
	"os/exec"
)


func RemoveDir(s string) {
    cmd := exec.Command("rm -rf %d", s)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error : ", err)
        return
    }

    log.Println(output)
}
