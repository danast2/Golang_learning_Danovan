package main

//выполняет параллельную выборку URL и сообщает
//о затраченном премени и размере ответа для каждого из них

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
	//ioutil.Discard - переменная из пакета io/ioutil (реализация интерфейса io.Writer), но отбрасывает все данные, которые в нее записывают
	//т.е. строчка ниже нужна просто для просмотра кол-ва байт занимаемых файлом и ошибок в нем
	file, err := os.Create(url)
	nbytes, err := io.Copy(file, resp.Body)
	defer resp.Body.Close() //исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
