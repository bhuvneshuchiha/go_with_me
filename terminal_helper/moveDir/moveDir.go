package movedir

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func MoveDir() {
	fmt.Println("Enter the name of the directory you want to move.")
	scanner := bufio.NewScanner(os.Stdin)

	var dirName string
	if scanner.Scan() {
		dirName = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if dirName == "" {
		fmt.Println("No directory name received")
		return
	}

	// Get absolute path
	absPath, err := filepath.Abs(dirName)
	if err != nil {
		fmt.Println("Error getting the directory's path")
		return
	}

	// Check if the directory exists
	info, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Directory does not exist")
		} else {
			fmt.Println("Error checking the directory:", err)
		}
		return
	}

	if !info.IsDir() {
		fmt.Println("The provided path is not a directory")
		return
	}

	fmt.Println("Directory exists")
	fmt.Printf("Path is: %s\n", absPath)

	fmt.Println("Enter the destination for this directory.")
	var destinationPath string
	destinationScanner := bufio.NewScanner(os.Stdin)

	if destinationScanner.Scan() {
		destinationPath = destinationScanner.Text()
	}

	if err := destinationScanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	if destinationPath == "" {
		fmt.Println("No destination path received")
		return
	}

	// Execute the move command
	cmd := exec.Command("mv", absPath, destinationPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error moving the directory:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Directory moved successfully")
}
