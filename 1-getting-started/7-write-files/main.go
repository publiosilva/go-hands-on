package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
			showLogs()
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

func registrateLog(url string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	var statusStr string

	if status {
		statusStr = "UP"
	} else {
		statusStr = "DOWN"
	}

	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + fmt.Sprintf("%-30s", url) + " | " + statusStr + "\n")

	file.Close()
}

func testUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Println(url, "is up")
		registrateLog(url, true)
	} else {
		fmt.Println(url, "in down")
		registrateLog(url, false)
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

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
