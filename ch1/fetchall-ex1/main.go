package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	logLine := fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)

	file, err := os.OpenFile("./result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		ch <- fmt.Sprintf("ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(logLine)
	if err != nil {
		ch <- fmt.Sprintf("ошибка записи в файл: %v", err)
		return
	}

	ch <- logLine
}
