package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int, 6) // 6 - емкость мапы, Это не означает, что карта будет ограничена ровно 6 элементами,
	// просто это предоставляет Go информацию о начальном размере карты, чтобы избежать динамического изменения её размера, если это возможно.
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type: %T\n", counts)
	fmt.Printf("Type: %T\n", input)
	for i := 0; i < 6 && input.Scan(); i++ {
		counts[input.Text()]++
	}

	//игнорим потенциальные ошибки из input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
