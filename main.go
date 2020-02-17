package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/moorada/log"
)

func main() {

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"English", "Italian"},
	}
	_, result, err := prompt.Run()

	switch result {
	case "English":
		english()
	case "Italian":
		italian()
	default:
		fmt.Printf("English")
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func english() {
	fmt.Println("TODO")
}

func italian() {
	for {
		min := 10
		max := 99
		randomNumber := rand.Intn(max-min) + min

		fmt.Println("Write a correct word:", randomNumber)
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		isRight := rightAssociation(response, rune(randomNumber), ItalianMS)
		if isRight {
			fmt.Println("Correct answer!")
		} else {
			fmt.Println("Wrong!!")
		}
	}
}

func rightAssociation(word string, number rune, ms MS) bool {
	//per ogni lettera della parola
	for _, c := range word {
		if IsVowel(c) {
			continue
		}
		compatibleFigures := compatibleFigures(c, ms)
		//Se trova solo una figura compatibile è quella giusta
		if len(compatibleFigures) == 1 {
			if number == compatibleFigures[0].char{
				return true
			}
		} else {
			//altrimenti bisogna vedere la seconda lettera
		}

	}
	return false
}

func compatibleFigures(c rune, ms MS) []Figure {
	var compatibleFigures []Figure
	//se è una consonante seleziona le figures compatibili
	for _, f := range ms {
		for _, ph := range f.phonetics {
			if strings.ContainsRune(ph, c) {
				compatibleFigures = append(compatibleFigures, f)
				break
			}
		}
	}
	return compatibleFigures
}