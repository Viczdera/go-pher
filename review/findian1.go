package main

import (
	"fmt"
	"strings"
)

func main() {
	// Richiede all'utente di inserire una stringa
	fmt.Print("Inserisci una stringa: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	// Converte la stringa in minuscolo per rendere il confronto non sensibile alle maiuscole
	lowercaseInput := strings.ToLower(input)

	// Verifica se la stringa inizia con 'i', termina con 'n' e contiene 'a'
	if strings.HasPrefix(lowercaseInput, "i") && strings.HasSuffix(lowercaseInput, "n") && strings.Contains(lowercaseInput, "a") {
		fmt.Println("Trovato!")
	} else {
		fmt.Println("Non trovato!")
	}
}
