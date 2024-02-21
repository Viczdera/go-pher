package main

import "fmt"

func main() {
	var programList = []string{"Golang", "javasacript", "Python"}
	//appending
	programList = append(programList, "C++", "Java")

	programList = append(programList[1:4]) //last val range is non inclusive
	fmt.Println(programList)

}
