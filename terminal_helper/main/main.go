package main

import (
	"bufio"
	"log"
	"os"

	"github.com/bhuvneshuchiha/term_help/choices"
	"github.com/bhuvneshuchiha/term_help/deleteFile"
	dircreate "github.com/bhuvneshuchiha/term_help/dirCreate"
	filecreate "github.com/bhuvneshuchiha/term_help/fileCreate"
	movedir "github.com/bhuvneshuchiha/term_help/moveDir"
	movefile "github.com/bhuvneshuchiha/term_help/moveFile"
	processkill "github.com/bhuvneshuchiha/term_help/processKill"
	"github.com/bhuvneshuchiha/term_help/remove"
)

func main() {

	choices.CreateChoices()

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Command 1 -> Create Directory
	if input == "1" {
		dircreate.CreateDir()
	}

	//Command 2 -> To remove a directory
	if input == "2" {
		remove.RemoveDir()
	}

	//Command 3 -> Create a file
	if input == "3" {
		filecreate.CreateFile()
	}

	//Command 4 -> Delete a file
	if input == "4" {
		deleteFile.DeleteFile()
	}

	//Command 5 -> Move a directory
	if input == "5" {
		movedir.MoveDir()
	}

	//Command 6 -> Move a file
	if input == "6" {
		movefile.MoveFile()
	}

	//Command 7 -> Check for PIDs and then kill the processes based on PIDs
	if input == "7" {
		processkill.KillProcess()
	}

}
