package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["GO"] = "Golang"
	languages["PY"] = "Python"

	delete(languages, "P")
	fmt.Println("languages: ", languages)

	for key, value := range languages {
		fmt.Printf("language %v: %v\n", key, value)
	}

}
