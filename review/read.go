package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {

	s := make([]Person, 0, 0)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please provide your file path")
	name, err := reader.ReadString('\n')
	name = strings.Trim(name, " \n")

	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		i := strings.Index(line, " ")
		p := Person{
			firstName: line[:i],
			lastName:  line[i+1:],
		}

		s = append(s, p)
	}

	for _, p := range s {
		fmt.Printf("%s %s\n", p.firstName[:], p.lastName[:])
	}
}
