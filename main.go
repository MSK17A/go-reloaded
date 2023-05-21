package main

import (
	"fmt"
	"os"
)

func main() {
	file_name := os.Args[1:]
	dat,_ := os.ReadFile(string(file_name[0]))
	fmt.Print(dat)
}
