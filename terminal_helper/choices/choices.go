package choices

import "fmt"

func CreateChoices() {
	fmt.Println("You have few choices : ")
	fmt.Println("1- Create a directory.")
	fmt.Println("2- Delete a directory.")
	fmt.Println("3- Create a file.")
	fmt.Println("4- Delete a file.")
	fmt.Println("5- Move a directory.")
	fmt.Println("6- Move a file.")
	fmt.Println("7- Kill a process.")
	fmt.Println("8- Create a TMUX session. (For this you need to have tmux already installed)")
	fmt.Println("9- Kill a TMUX session. (For this you need to have tmux already installed)")
	fmt.Println("10- Attach to a tmux session detaching other clients. (For this you need to have tmux already installed)")
}
