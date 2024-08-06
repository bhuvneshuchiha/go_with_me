package processkill

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func KillProcess() {
	fmt.Println("Let me show you the current processes that are running")

	cmd := exec.Command("ps")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(reflect.TypeOf(output))
	fmt.Println(string(output))

	fmt.Println("Enter the PID you wish to kill")
	var pid string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		pid = scanner.Text()
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	com := exec.Command("kill", "-9", pid)
	_, er := com.CombinedOutput()
	if er == nil {
		log.Fatal(err)
		return
	}
	fmt.Println("The process has been killed")
}
