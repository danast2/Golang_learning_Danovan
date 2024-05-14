package main

import "fmt"

const boiling = 212.0

func main() {
	var f = boiling
	var c = (f - 32) * 5 / 9
	fmt.Printf("Температура кипения =%gF или %gC\n", f, c)
	//вывод
	//температура кипения = 212F или 100 град
}

//константа bolling представляет собой объявление уровня пакета, тогда как f и c являются
//локальными переменными для ф-ции main
