package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if num := 3; num <= 3 {
		fmt.Println("num is less than or equal to 3")
	} else {
		fmt.Println("greater than three")
	}
	fmt.Printf("\n")

	//switch

	//rand.Seed(time.Now().UnixNano())...this is depreciated
	diceNum := rand.Intn(7)

	switch diceNum {
	case 0:
		fmt.Println("Dice value is 0, roll again")
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("Got 6, can roll agian")

	}
	fmt.Printf("\n")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(days); i++ {
		fmt.Printf("With normal for statement %v\n", days[i])
	}
	for i := range days {
		fmt.Printf("With range %v\n", days[i])
	}

	var someValue int = 0
	for someValue < 10 {
		fmt.Println("Some value:", someValue)
		if someValue == 4 {
			goto endIt
		}
		someValue++
	}

endIt:
	fmt.Println("End the program")

}
