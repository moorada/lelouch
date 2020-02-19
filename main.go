package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/evilsocket/islazy/log"
	"github.com/manifoldco/promptui"

	"github.com/moorada/lelouch/dictionary"
)

const UpdateDictionaries = "Aggiorna i dizionari"
const GameWordToNumber = "Gioco: Converti in numero"
const GameNumberToWord = "Gioco: Converti in parola"
const GameMix = "Gioco: Misto"
const Converter = "Convertitore"
const simpleLevel = "Sono all'inizio"
const mediumLevel = "Sono pratico con la conversione fonetica"
const extremeLevel = "Sono un convertitore fonetico vivente"
const Stats = "Statistiche"
const creazioneDeiDizionariInCorso = "Creazione dei dizionari in corso..."

var completeDictionary dictionary.AS
var commonDictionary dictionary.AS
var simpleDictionary dictionary.AS

func main() {
	var err1, err2, err3 error
	err1, completeDictionary = dictionary.GetCompleteDictionary()
	err2, commonDictionary = dictionary.GetCommonDictionary()
	err3, simpleDictionary = dictionary.GetSimpleDictionary()

	if err1 != nil || err2 != nil || err3 != nil {
		dictionary.MakeDictionaries()
		showLoadingDictionaries()
	}

	prompt := promptui.Select{
		Label: "Seleziona modalità",
		Items: []string{UpdateDictionaries, Converter, GameWordToNumber, GameNumberToWord, GameMix, Stats},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error: %s", err)
		return
	}

	switch result {
	case UpdateDictionaries:
		dictionary.MakeDictionaries()
	case Converter:
		convert()
	case GameWordToNumber:
		gameWN()
	case GameNumberToWord:
		gameNW()
	case GameMix:
		gameMix()
	case Stats:
		stats()
	}

}

func showLoadingDictionaries() {
	fmt.Println(creazioneDeiDizionariInCorso)
}

func gameMix() {
	log.Info("TODO")
}
func convert() {
	for {
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		if _, err := strconv.Atoi(response); err == nil {
			convertNW(response)
		} else {
			convertWN(response)
		}
	}
}

func stats() {
	var numberOfWords int
	for _, ws := range completeDictionary {
		numberOfWords += len(ws)
	}
	fmt.Println("Nel dizionario di tutte le parole ci sono", numberOfWords, "parole")
	numberOfWords = 0
	for _, ws := range commonDictionary {
		numberOfWords += len(ws)
	}
	fmt.Println("Nel dizionario delle parole comuni ci sono", numberOfWords, "parole")
	numberOfWords = 0
	for _, ws := range simpleDictionary {
		numberOfWords += len(ws)
	}
	fmt.Println("Nel dizionario semplice ci sono", numberOfWords, "parole")

}

func convertWN(word string) {
	fmt.Println(dictionary.WordToNumber(word))
}

func convertNW(number string) {

	words := completeDictionary[number]
	if len(words) == 0 {
		fmt.Println("Nessuna parola compatibile")
	} else {
		fmt.Println("Possibili parole: ")
	}
	for i, w := range words {
		fmt.Print(w)
		if i < len(words)-1 {
			fmt.Print(", ")
		} else {
			fmt.Println(".")
		}
	}
}

func gameWN() {

	d := chooseLevelGame()
	var numbers []string

	for k, _ := range d {
		numbers = append(numbers, k)
	}

	for {
		var ws []string
		randomIndexNumber := rand.Intn(len(numbers) - 1)
		ws = d[numbers[randomIndexNumber]]

		randomIndexWord := 0
		if len(ws) > 1 {
			randomIndexWord = rand.Intn(len(ws) - 1)
		}

		fmt.Println("Converti la parola:", ws[randomIndexWord])
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		if response == numbers[randomIndexNumber] {
			fmt.Println("Corretto!")
		} else {
			fmt.Println("Sbagliato!!")
		}
	}
}

func chooseLevelGame() dictionary.AS {
	prompt := promptui.Select{
		Label: "Seleziona il livello",
		Items: []string{simpleLevel, mediumLevel, extremeLevel},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	switch result {
	case mediumLevel:
		return commonDictionary
	case extremeLevel:
		return completeDictionary
	default:
		return simpleDictionary
	}
}

func gameNW() {

	var numbers []string

	d := chooseLevelGame()

	for k, _ := range d {
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
		if dictionary.WordToNumber(response) == numbers[randomNumber] {
			fmt.Println("Corretto!")
		} else {
			fmt.Println("Sbagliato!!")
		}
	}
}
