package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/briandowns/spinner"
	"github.com/evilsocket/islazy/log"
	"github.com/manifoldco/promptui"

	"github.com/moorada/lelouch/dictionary"
)

const Select = "Seleziona modalità"
const UpdateDictionaries = "Aggiorna i dizionari"
const GameWordToNumber = "Gioco: Converti in numero"
const GameNumberToWord = "Gioco: Converti in parole"
const GameMix = "Gioco: Misto"
const Converter = "Converti numeri o parole"
const simpleLevel = "Sono all'inizio"
const mediumLevel = "Sono pratico con la conversione fonetica"
const extremeLevel = "Sono un convertitore fonetico vivente"
const Stats = "Statistiche"
const MostraTabellaCompleta = "Visualizza tabella completa"
const Continua = "Continua"
const MenuPrincipale = "Torna al menu' principale"
const charsetNumber = "0123456789"
const MostraSoluzioni = "Mostra soluzioni"

const (
	simplelevelint  = 5
	mediumlevelint  = 10
	extremelevelint = 20
)

var completeDictionary dictionary.AS
var commonDictionary dictionary.AS
var levelGame = simpleLevel

func main() {
	rand.Seed(time.Now().UnixNano())
	var err1, err2 error
	err1, completeDictionary = dictionary.GetCompleteDictionary()
	err2, commonDictionary = dictionary.GetCommonDictionary()

	if err1 != nil || err2 != nil {
		loadingDictionaries()
		err1, completeDictionary = dictionary.GetCompleteDictionary()
		err2, commonDictionary = dictionary.GetCommonDictionary()
		if err1 != nil || err2 != nil {
			log.Fatal("Error to get dictionaries ")
		}
	}

	start()

}

func start() {
	for {
		prompt := promptui.Select{
			Label: Select,
			Items: []string{UpdateDictionaries, Converter, GameWordToNumber, GameNumberToWord, GameMix},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal("Error: %s", err)
			return
		}

		switch result {
		case UpdateDictionaries:
			loadingDictionaries()
		case Converter:
			convert()
		case GameWordToNumber:
			gamePhraseWN()
		case GameNumberToWord:
			gamePhraseNW()
		case GameMix:
			gameMix()
		}

	}
}

func loadingDictionaries() {
	s := spinner.New(spinner.CharSets[26], 100*time.Millisecond)
	s.Prefix = "Aggiornamento dizionari in corso "
	s.FinalMSG = "Aggiornamento completato"
	s.Start()
	dictionary.MakeDictionaries()
	s.Stop()
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
}

func chooseLevelGame() int {
	prompt := promptui.Select{
		Label: "Seleziona il livello",
		Items: []string{simpleLevel, mediumLevel, extremeLevel},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	return getLevelGame(result)

}

func getLevelGame(result string) int {
	switch result {
	case simpleLevel:
		randomNumber := rand.Intn(simplelevelint) + 1
		return randomNumber
	case mediumLevel:
		randomNumber := rand.Intn(mediumlevelint) + simplelevelint
		return randomNumber
	default:
		randomNumber := rand.Intn(extremelevelint) + mediumlevelint
		return randomNumber
	}
}

func stringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
