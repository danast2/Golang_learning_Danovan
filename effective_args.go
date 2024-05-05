package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Command: ", os.Args[0], "Arguments", strings.Join(os.Args[1:], " "))
	//fmt.Println("Command:", os.Args[0], "Arguments:", strings.Join(os.Args[1:], " "))
}
