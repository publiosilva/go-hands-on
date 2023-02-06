package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const attempts = 5
const delay = 5 * time.Second

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

func testUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Println(url, "is up")
	} else {
		fmt.Println(url, "in down")
	}
}

func readUrlsFile() []string {
	var urls []string

	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return urls
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	urls := readUrlsFile()

	for i := 0; i < attempts; i++ {
		for _, url := range urls {
			testUrl(url)
		}

		time.Sleep(delay)
	}
}
