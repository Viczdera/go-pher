package main

import (
	"fmt"
	"os"
)

type User struct {
	fname string //struct field should be caps
	lname string
}

func main() {
	content, err := os.ReadFile("names.txt")
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(content)
}
