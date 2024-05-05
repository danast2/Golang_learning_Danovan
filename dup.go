package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < 6 && input.Scan(); i++ {
		counts[input.Text()]++
	}

	//игнорим потенциальные ошибки из input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
