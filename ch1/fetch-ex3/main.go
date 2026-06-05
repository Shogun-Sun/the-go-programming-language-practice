package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var args []string = os.Args[1:]

	for i := 0; i < len(args); i++ {
		if args[i] == "--help" {
			fmt.Printf("Доступные аргументы:\n--help   - Помощь\n--status - Статус страницы\n--body   - Содержимое страницы\nПоследним аргументом всегда идёт URL\n")
			return
		}
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Ошибка: не указан URL. Используйте --help для справки.")
		os.Exit(1)
	}

	lastArg := args[len(args)-1]
	if strings.HasPrefix(lastArg, "-") {
		fmt.Fprintf(os.Stderr, "Ошибка: указан флаг %s, но не указан URL-адрес страницы.\n", lastArg)
		os.Exit(1)
	}

	showStatus := false
	showBody := false

	for i := 0; i < len(args)-1; i++ {
		switch args[i] {
		case "--status":
			showStatus = true
		case "--body":
			showBody = true
		default:
			fmt.Fprintf(os.Stderr, "Неизвестный флаг: %s\n", args[i])
			os.Exit(1)
		}
	}

	if !showStatus && !showBody {
		showBody = true
	}

	url := getUrl(args)
	resp := sendRequest(url)
	defer resp.Body.Close()

	if showStatus {
		var status string = resp.Status
		const (
			Reset = "\x1b[0m"
			Red   = "\x1b[31m"
			Green = "\x1b[32m"
		)

		if strings.HasPrefix(status, "200") {
			fmt.Printf("\nСтатус страницы: %s%s%s\n\n", Green, resp.Status, Reset)
		} else {
			fmt.Printf("\nСтатус страницы: %s%s%s\n\n", Red, resp.Status, Reset)
		}
	}

	if showBody {
		io.Copy(os.Stdout, resp.Body)
	}
}

func getUrl(args []string) string {
	url := args[len(args)-1]

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

func sendRequest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "curl: %v\n", err)
		os.Exit(1)
	}

	return resp
}
