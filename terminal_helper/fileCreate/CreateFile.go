package filecreate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bhuvneshuchiha/term_help/ls"
)



func CreateFile() {
    fmt.Println("Enter the name of the file you wish to create.")
    scanner := bufio.NewScanner(os.Stdin)
    var fileName string

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    if scanner.Scan() {
        fileName = scanner.Text()
    }

    cmd := exec.Command("touch", fileName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    ls.LsCommand()
    log.Println(output)


}
