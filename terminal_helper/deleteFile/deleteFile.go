package deleteFile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bhuvneshuchiha/term_help/ls"
)

func DeleteFile() {

    var choice string
    fmt.Println("Enter the name of the file you want to delete.")
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        choice = scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    cmd := exec.Command("rm", "-rf", choice)
    output, err := cmd.CombinedOutput()

    if err != nil {
        log.Fatal(err)
    }
    log.Println(output)
    ls.LsCommand()
}
