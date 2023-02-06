package main

import "fmt"
import "os"
import "net/http"

func main() {
	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Monitoring...")

	url := "https://www.google.com"
	resp, _ := http.Get(url)

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Println(url, "is up")
	} else {
		fmt.Println(url, "in down")
	}
}