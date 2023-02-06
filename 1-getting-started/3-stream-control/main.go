package main

import "fmt"
import "os"

func main() {
	showMenu()

	command := readCommand()

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Loading logs...")
	// } else if command == 0 {
	// 	fmt.Println("Exiting...")
	// } else {
	// 	fmt.Println("Command not recognized.")
	// }

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Loading logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Command not recognized.")
		os.Exit(-1)
	}
}

func showMenu() {
	fmt.Println("0 - Exit")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}