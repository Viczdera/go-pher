package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	fname string //struct field should be caps
	lname string
}

func main() {
	// content, err := os.ReadFile("names.txt")
	var path string
	fmt.Print("Enter file path: ")
	_, scanErr := fmt.Scan(&path)
	if scanErr != nil {
		print(scanErr)
		return
	}
	//open and read
	file, err := os.Open(path)
	if err != nil {
		print(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	readNames := []User{}

	for scanner.Scan() {
		fullname := strings.Split(scanner.Text(), " ")
		readNames = append(readNames, User{
			fname: fullname[0],
			lname: fullname[1],
		})
	}

	for _, name := range readNames {
		fmt.Printf("Firstname is %s and lastname is %s\n", name.fname, name.lname)
	}

}
