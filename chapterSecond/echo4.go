package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}