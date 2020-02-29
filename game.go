package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/evilsocket/islazy/log"
	"github.com/manifoldco/promptui"

	"github.com/moorada/lelouch/dictionary"
)

func gamePhraseWN() {
	chooseLevelGame()
	for {
		if wn() == -1 {
			return
		}

	}

}

func gameMix() {
	chooseLevelGame()
	for {
		random := rand.Intn(2)
		if random == 0 {
			wn()
		} else if random == 1 {
			nw()
		} else {
			log.Fatal("fatale!")
		}
	}

}

func wn() int {
	l := getLevelGame(levelGame)
	number := stringWithCharset(l, charsetNumber)
	SS := numberToWordSet(number)
	randomWordSet := rand.Intn(len(SS) - 1)
	wordSet := SS[randomWordSet]
	phrase := ""
	for _, ws := range wordSet {
		wordSlice := strings.Split(ws, ",")
		randomWord := 0
		if len(wordSlice) > 1 {
			randomWord = rand.Intn(len(wordSlice) - 1)
		}

		phrase = phrase + wordSlice[randomWord] + " "
	}
	fmt.Println("Converti le parole:", phrase)

	var response string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response = scanner.Text()
	}
	re := regexp.MustCompile("[^a-zA-Z0-9_]+")
	response = re.ReplaceAllString(response, "")

	if number == response {
		fmt.Println("Corretto!")
	} else {
		fmt.Println("Sbagliato!! \nLa soluzione è: " + number)
	}

	prompt := promptui.Select{
		Label: "Seleziona",
		Items: []string{Continua, MenuPrincipale},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	switch result {
	case MenuPrincipale:
		return -1
	default:
		return 0

	}
}

func gamePhraseNW() {
	chooseLevelGame()
	for {
		if nw() == -1 {
			return
		}

	}

}
func nw() int {

	l := getLevelGame(levelGame)
	number := stringWithCharset(l, charsetNumber)
	fmt.Println("Converti il numero:", number)

	var response string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response = scanner.Text()
	}

	if dictionary.WordToNumber(response) == number {
		fmt.Println("Corretto!")
	} else {
		fmt.Println("Sbagliato!!")
	}

	prompt := promptui.Select{
		Label: "Seleziona",
		Items: []string{Continua, MostraSoluzioni, MenuPrincipale},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	switch result {
	case MostraSoluzioni:
		showTableWordSet(dictionary.WordToNumber(response))
		return 0
	case MenuPrincipale:
		return -1
	default:
		return 0

	}
}
