package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bhuvneshuchiha/term_help/choices"
	"github.com/bhuvneshuchiha/term_help/ls"
	"github.com/bhuvneshuchiha/term_help/remove"
)

func main() {
    var input string


    for {
        //Command 1 -> Create Directory
        choices.CreateChoices()
        if input == "1" {
            var choice string

            fmt.Println("Enter the directory name.")
            scanner := bufio.NewScanner(os.Stdin)

            if scanner.Scan() {
                choice = scanner.Text()
            }

            if err := scanner.Err(); err != nil {
                log.Fatal(err)
            }
            answer := "mkdir " + choice
            cmd := exec.Command(answer)

            output, err := cmd.CombinedOutput()
            if err != nil {
                log.Fatal(err)
                choices.CreateChoices()
            }
            ls.LsCommand()
            log.Print(output)
        }

        //Command 2 -> To remove a directory

        fmt.Println("Do you want to perform more file operations ? Press 1 for yes")
        scanner := bufio.NewScanner(os.Stdin)

        var choice string
        if scanner.Scan() {
            if scanner.Text() == "1"  {
                choices.CreateChoices()
                scanner := bufio.NewScanner(os.Stdin)
                if err := scanner.Err(); err != nil {
                    log.Fatal(err)
                }
                choice = scanner.Text()
                if choice == "2" {
                    fmt.Println("Enter the directory name you wish to delete")
                    var dir string
                    scanner := bufio.NewScanner(os.Stdin)
                    if err := scanner.Err(); err!= nil {
                        log.Fatal(err)
                    }
                    remove.RemoveDir(dir)
                }
            }
        }

        //Command 3 -> Create a file
    }



}
