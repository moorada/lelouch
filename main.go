package main

import (
	"fmt"
	"math/rand"

	"github.com/manifoldco/promptui"
	"github.com/moorada/log"
)

const DA_PAROLA_A_NUMERO = "Da parola a numero"
const DA_NUMERO_A_PAROLA = "Da numero a parola"

func main() {

	prompt := promptui.Select{
		Label: "Seleziona gioco",
		Items: []string{DA_PAROLA_A_NUMERO, DA_NUMERO_A_PAROLA},
	}
	_, result, err := prompt.Run()

	switch result {
	case DA_PAROLA_A_NUMERO:
		gameNP()
	case DA_NUMERO_A_PAROLA:
		gamePN()
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func gameNP(){
	//TODO
}

func gamePN(){
	for {
		min := 10
		max := 99
		randomNumber := rand.Intn(max-min) + min
		fmt.Println("Converti il numero:", randomNumber)
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		//TODO
	}

}

