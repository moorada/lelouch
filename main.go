package main

import (
	"fmt"
	"math/rand"

	"github.com/evilsocket/islazy/log"
	"github.com/manifoldco/promptui"
)

const MakeDizionario = "Crea un dizionario"
const GameDaParolaANumero = "GAME: Da parola a numero"
const GameDaNumeroAParola = "GAME: Da numero a parola"
const ConvertiParolaANumero = "CONVERTI: Da parola a numero"
const ConvertiNumeroAParola = "CONVERTI: Da numero a parola"

var indexedDictionary AS

func main() {

	indexedDictionary = getIndexedDictionary()

	prompt := promptui.Select{
		Label: "Seleziona gioco",
		Items: []string{MakeDizionario, ConvertiParolaANumero, ConvertiNumeroAParola, GameDaParolaANumero, GameDaNumeroAParola},
	}
	_, result, err := prompt.Run()

	switch result {
	case ConvertiParolaANumero:
		convertiPN()
	case ConvertiNumeroAParola:
		convertiNP()
	case GameDaParolaANumero:
		gamePN()
	case GameDaNumeroAParola:
		gameNP()
	case MakeDizionario:
		makeDizionario()
	}

	if err != nil {
		log.Fatal("Error: %s", err)
		return
	}
}

func convertiPN() {

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	fmt.Println(wordToNumber(response))
}

func convertiNP() {

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	words := indexedDictionary[response]
	fmt.Print("Possibili parole: ")
	for _, w := range words {
		fmt.Print(w + ", ")
	}
}

func gamePN() {
	var numbers []string

	for k, _ := range indexedDictionary {
		numbers = append(numbers, k)
	}

	for {
		var ws []string

		index := 0
		if len(numbers) > 1 {
			index = rand.Intn(len(numbers) - 1)
		}
		ws = indexedDictionary[numbers[index]]

		randomIndex := rand.Intn(len(ws) - 1)
		fmt.Println("Converti la parola:", ws[randomIndex])
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		if response == numbers[index] {
			fmt.Println("Corretto!")
		} else {
			fmt.Println("Sbagliato!!")
		}
	}
}

func gameNP() {

	var numbers []string

	for k, _ := range indexedDictionary {
		numbers = append(numbers, k)
	}

	for {
		randomNumber := rand.Intn(len(numbers) - 1)
		fmt.Println("Converti il numero:", numbers[randomNumber])
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		if wordToNumber(response) == numbers[randomNumber] {
			fmt.Println("Corretto!")
		} else {
			fmt.Println("Sbagliato!!")
		}
	}
}
