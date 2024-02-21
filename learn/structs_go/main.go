package main

import "fmt"

func main() {
	dera := User{Name: "Chdiera", Email: "deracvi@gmail.com", Password: "o282974892", Phone: 12343423}

	fmt.Printf("User details %+v\n", dera)
	fmt.Printf("User's name is %v. User's email is %v\n", dera.Name, dera.Email)

}

type User struct {
	Name     string
	Email    string
	Password string
	Phone    int
}
