package processkill

import (
	"fmt"
	"log"
	"os/exec"
)

func KillProcess() {
	fmt.Println("Let me show you the current processes that are running")

	cmd := exec.Command("ps aux | grep", "all")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
