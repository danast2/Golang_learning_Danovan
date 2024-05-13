package main

//выполняет параллельную выборку URL и сообщает
//о затраченном премени и размере ответа для каждого из них

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
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		//получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //Отправка в канал ch
		return
	}
	//ioutil.Discard - переменная из пакета io/ioutil (реализация интерфейса io.Writer), но отбрасывает все данные, которые в нее записывают
	//т.е. строчка ниже нужна просто для просмотра кол-ва байт занимаемых файлом и ошибок в нем
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

}
