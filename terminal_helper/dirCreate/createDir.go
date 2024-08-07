package dircreate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bhuvneshuchiha/term_help/choices"
	"github.com/bhuvneshuchiha/term_help/ls"
)



func CreateDir() {
    var choice string

    fmt.Println("Enter the directory name.")
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        choice = scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    cmd := exec.Command("mkdir", choice)

    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
        choices.CreateChoices()
    }
    ls.LsCommand()
    log.Print(output)
}
