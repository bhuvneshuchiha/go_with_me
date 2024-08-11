package fileOps

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bhuvneshuchiha/todo_app/models"
)

func WriteData(data []models.Todo) {
	json_data, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error marshalling data :", err)
		return
	}

	// Open the file and write the data
	file, err := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the file :", err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info")
		return
	}

	if fileInfo.Size() > 0 {
		_, err := file.WriteString(",\n")
		if err != nil {
			fmt.Println("Error writing to the file :", err)
			return
		}
	}

	_, er := file.Write(json_data)
	if er != nil {
		fmt.Println("Couldnt write the stupid data")
		return
	}
	fmt.Println("Data written successfully")

}

// @@ Function to read file. Fixing bug here
func ReadFile() ([]models.Todo, error) {
    data, err := os.ReadFile("out.txt")
    if err != nil {
        return nil , err
    }
    var todos []models.Todo
    err = json.Unmarshal(data, &todos)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Raw data:", string(data))
        return nil, err
    }
    return todos, nil
}

func GetFileName() string {
	fmt.Println("Enter your name")
	var fileName string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fileName = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Cannot get the file name")
	}

	if fileName != "" {
		return fileName
	} else {
		return "Couldnt get the file name"
	}
}
