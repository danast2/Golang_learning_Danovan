package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchFirstNew(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		//получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchFirstNew(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //Отправка в канал ch
		return
	}
	defer resp.Body.Close() //исключение утечки ресурсов

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	// Создание файла для записи HTML-кода
	file, err := os.Create(getFileName(url))
	if err != nil {
		ch <- fmt.Sprintf("while creating file for %s: %v", url, err)
		return
	}
	defer file.Close()

	// Запись HTML-кода в файл
	_, err = io.WriteString(file, string(htmlData))
	if err != nil {
		ch <- fmt.Sprintf("while writing to file for %s: %v", url, err)
		return
	}

	nbytes := len(htmlData)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func getFileName(url string) string {
	fileName := ""
	for i := len(url) - 1; i > 0; i-- {
		if url[i] == '/' {
			fileName = url[i+1:]
			break
		}
	}
	if fileName == "" {
		fileName = "index.html"
	}
	return fileName
}
