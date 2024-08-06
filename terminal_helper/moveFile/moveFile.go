package movefile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bhuvneshuchiha/term_help/ls"
)


func MoveFile() {

    fmt.Println("Enter the name of the file you want to move.")
    var sourceFile string

    fileScanner := bufio.NewScanner(os.Stdin)
    if fileScanner.Scan() {
       sourceFile = fileScanner.Text()
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
        return
    }

    if sourceFile == "" {
        fmt.Println("File name received empty!!")
        return
    }

    //Get absolute path
    absPath, err := filepath.Abs(sourceFile)
    if err != nil {
        fmt.Println("Could not find the file in your path")
    }

    //Check if the file exists
    info, err := os.Stat(absPath)
    if err != nil {
        if os.IsNotExist(err){
            fmt.Println("File does not exists")
        } else {
            fmt.Println("Could not locate the file")
        }
        return
    }

    if info.IsDir(){
        fmt.Println("A directory is found at this path, please use the move directory command instead")
        return
    }

    //File found now move it to the destination

    fmt.Println("Enter the full path of the destination")
    var destinationPath string

    destScan := bufio.NewScanner(os.Stdin)
    if destScan.Scan() {
        destinationPath = destScan.Text()
    }

    if err := destScan.Err(); err != nil {
        log.Fatal(err)
    }
    if destinationPath == "" {
        fmt.Println("Received blank path")
    }

    //Get absolute path as the path received is not a file object its a string
    //filepath.Abs returns a string of absolute path

    absDestPath, err := filepath.Abs(destinationPath)
    if err != nil {
        fmt.Println("Could not determine the absolute path of the destination:", err)
        return
    }

    //os.Stat method returns an os.FileInfo object and accepts a string file path
    path, error := os.Stat(absDestPath)
    if error != nil {
        if os.IsNotExist(error){
            fmt.Println("Path is not valid")
        } else {
            fmt.Println("Directory where you want to move was not found")
        }
    }

    newFilePath := filepath.Join(absDestPath, filepath.Base(sourceFile))

    if path.IsDir() {
        cmd := exec.Command("mv", sourceFile, newFilePath)
        output, err := cmd.CombinedOutput()
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("File has been moved successfully")

        fmt.Println(output)
        ls.LsCommand()
    }
}
