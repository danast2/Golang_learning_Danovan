package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//counts := map[string]int{} можно было использовать такой вариант, тут Go автоматичеси инициализирует мапу
	//Когда вы используете make(), вы можете указать емкость (capacity) карты,
	//чтобы уменьшить необходимость в динамическом изменении размера, что может быть полезно для больших коллекций данных.

	counts := make(map[string]int, 6) // 6 - емкость мапы, Это не означает, что карта будет ограничена ровно 6 элементами,
	// просто это предоставляет Go информацию о начальном размере карты, чтобы избежать динамического изменения её размера, если это возможно.

	input := bufio.NewScanner(os.Stdin)
	//var input *bufio.Scanner = bufio.NewScanner(os.Stdin)

	fmt.Printf("Type: %T\n", counts) //%T формат вывода для типа данных переменной
	fmt.Printf("Type: %T\n", input)

	for i := 0; i < 6 && input.Scan(); i++ {
		counts[input.Text()]++
	}

	//игнорим потенциальные ошибки из input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
