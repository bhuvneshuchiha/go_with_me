package fileOps

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bhuvneshuchiha/todo_app/models"
)

func WriteData(data models.Todo) {
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
