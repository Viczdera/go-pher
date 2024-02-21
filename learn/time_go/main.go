package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Lets do time")

	present := time.Now()
	createdDate := time.Date(2022, time.December, 20, 23, 5, 0, 0, time.Local)
	fmt.Println(present)
	fmt.Println("Created date:", createdDate)
	fmt.Println("Formated date:", createdDate.Format("01-02-2006 Monday"))
}
