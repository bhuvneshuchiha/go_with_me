package remove

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bhuvneshuchiha/term_help/ls"
)

func RemoveDir() {
	fmt.Println("Enter the directory name you wish to delete:")
	scanner := bufio.NewScanner(os.Stdin)
	var inputPath string

	if scanner.Scan() {
		inputPath = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if inputPath == "" {
		fmt.Println("No directory name entered. Exiting.")
		return
	}

	// Get the absolute path
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}

	//If directory exists, delete it.
	if info, err := os.Stat(absPath); err == nil && info.IsDir() {
		fmt.Println("The directory exists.")
		fmt.Printf("Absolute path: %s\n", absPath)

		cmd := exec.Command("rm", "-rf", absPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %v, Output: %s\n", err, output)
			return
		}

		fmt.Println("Directory deleted successfully.")
		log.Println(string(output))

		ls.LsCommand()
	} else if os.IsNotExist(err) {
		fmt.Println("The directory does not exist.")
	} else {
		fmt.Printf("Error checking directory: %v\n", err)
	}

}
