package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bhuvneshuchiha/todo_app/todo_struct"
)

func main() {


    for {
        fmt.Println("Enter the title for your todo")
        scanner := bufio.NewScanner(os.Stdin)

        var title string
        var content string

        if scanner.Scan() {
            title = scanner.Text()
        }
        if err := scanner.Err(); err != nil {
            panic(err)
        }

        fmt.Println("Enter the content for your todo")
        sc := bufio.NewScanner(os.Stdin)

        if sc.Scan() {
            content = sc.Text()
        }
        if err := sc.Err(); err != nil {
            panic(err)
        }

        if title == "" || content == "" {
            fmt.Println("No field can be empty")
            return
        }
        todo_struct.CreateTodo(title, content)

        fmt.Println("If you are done adding todos press x to exit or c to continue")
        var choice string
        s := bufio.NewScanner(os.Stdin)
        if s.Scan() {
            choice = s.Text()
        }
        if err := s.Err(); err != nil {
            fmt.Println("Something went wrong")
            return
        }
        if choice == "c" {
            continue
        }else {
            break
        }
    }
    todo_struct.UpdateStatus(1, "Check", "eeemboom")
    todo_struct.RemoveTodo(1)
}
