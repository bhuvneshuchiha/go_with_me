package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
    InputFilePath string
    OutputFilePath string
}

func ReadLines(path string) ([]string, error) {

    file, err := os.Open(path)

    if err != nil {
        return nil, errors.New("Failed to open file")
    }

    scanner := bufio.NewScanner(file)
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    err = scanner.Err()

    if err != nil {
        return nil, errors.New("Failed to read line")
    }
    file.Close()
    return lines, nil

}


func(fm FileManager) WriteResult(data interface{})(error) {
    file, err := os.Create(fm.OutputFilePath)

    if err != nil {
        return errors.New("Failed to create this file")
    }
    encoder := json.NewEncoder(file)
    err = encoder.Encode(data)

    if err != nil {
        file.Close()
        return errors.New("Failed to convert the data to json")
    }

    file.Close()
    return nil
}

func New(inputPath, outputPath string) FileManager {
    return FileManager {
        InputFilePath: inputPath,
        OutputFilePath: outputPath,
    }
}
