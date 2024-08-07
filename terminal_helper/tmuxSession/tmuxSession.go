package tmuxSession

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func CreateTmux() {
	fmt.Println("Enter the name for the session")
	var seshName string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		seshName = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
	cmd := exec.Command("tmux", "new", "-s", seshName)

	// @@ Use of external package to run something that requires opening a terminal
	//session
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ptmx.Close()

	sc := bufio.NewScanner(ptmx)
	if sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Session created bitch")

}
