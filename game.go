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
	l := chooseLevelGame()
	for {
		number := stringWithCharset(l, charsetNumber)
		SS := numberToWordSet(number)
		randomWordSet := rand.Intn(len(SS) - 1)
		wordSet := SS[randomWordSet]
		phrase := ""
		for _, ws := range wordSet {
			wordSlice := strings.Split(ws, ",")

			randomWord := rand.Intn(len(wordSlice) - 1)
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

	}

}

func gamePhraseNW() {
	l := chooseLevelGame()
	for {
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
			Items: []string{MostraSoluzioni, Continua, MenuPrincipale},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal("Error: %s", err)
			return
		}
		switch result {
		case MostraSoluzioni:
			showTableWordSet(dictionary.WordToNumber(response))
		case MenuPrincipale:
			start()
		default:
			return
		}

	}

}
