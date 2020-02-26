package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/evilsocket/islazy/log"
	"github.com/manifoldco/promptui"

	"github.com/moorada/lelouch/dictionary"
)

const UpdateDictionaries = "Aggiorna i dizionari"
const GameWordToNumber = "Gioco: Converti in numero"
const GameNumberToWord = "Gioco: Converti in parola"
const GameMix = "Gioco: Misto"
const Converter = "Convertitore (singola parola)"
const ConverterPhrase = "Convertitore (frase)"
const simpleLevel = "Sono all'inizio"
const mediumLevel = "Sono pratico con la conversione fonetica"
const extremeLevel = "Sono un convertitore fonetico vivente"
const Stats = "Statistiche"
const creazioneDeiDizionariInCorso = "Creazione dei dizionari in corso..."
const charsetNumber = "0123456789"

var completeDictionary dictionary.AS
var commonDictionary1000 dictionary.AS
var simpleDictionary dictionary.AS
var commonDictionary7000 dictionary.AS

func main() {

	var err1, err2, err3 error
	err1, completeDictionary = dictionary.GetCompleteDictionary()
	err2, commonDictionary1000 = dictionary.GetCommonDictionary()
	err3, simpleDictionary = dictionary.GetSimpleDictionary()

	if err1 != nil || err2 != nil || err3 != nil {
		dictionary.MakeDictionaries()
		showLoadingDictionaries()
		err1, completeDictionary = dictionary.GetCompleteDictionary()
		err2, commonDictionary1000 = dictionary.GetCommonDictionary()
		err3, simpleDictionary = dictionary.GetSimpleDictionary()
		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatal("Error to get dictionaries ")
		}
	}

	/*SS := numberToPhrases(dictionary.WordToNumber("32524"), 1)

	for _, s := range SS {

		for _, x := range s {
			fmt.Print(x + "||")
		}
		fmt.Println("")
	}*/

	/*var j int
	for j = 0; j < len(SS); j++ {
		ok := true
		sw := ""
		sn := ""
		for i := 0; i < len(SS[j]); i++ {
			sn = sn + SS[j][i] + "|"
			if commonDictionary[SS[j][i]] != nil {
				sw = sw + commonDictionary[SS[j][i]][0] + "|"
			} else {
				ok = false
			}
		}
		if ok {
			fmt.Println(sn + " --> " + sw)
		}
	}*/
	/*fmt.Println("lunghezza: ", j)*/


	start()

}

func start() {
	prompt := promptui.Select{
		Label: "Seleziona modalità",
		Items: []string{UpdateDictionaries, Converter, ConverterPhrase, GameWordToNumber, GameNumberToWord, GameMix, Stats},
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
	case ConverterPhrase:
		convertPhrase()
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

func convertPhrase() {
	//TODO
}

func showLoadingDictionaries() {
	fmt.Println(creazioneDeiDizionariInCorso)
}

func gameMix() {
	//TODO
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
	for _, ws := range commonDictionary1000 {
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
	commonWords := commonDictionary1000[number]
	if len(commonWords) == 0 {
		fmt.Println("Nessuna parola comune compatibile")
	} else {
		fmt.Println("Possibili parole comuni: ")
		for i, w := range commonWords {
			fmt.Print(w)
			if i < len(commonWords)-1 {
				fmt.Print(", ")
			} else {
				fmt.Println(".")
			}
		}
	}
	if len(words) == 0 {
		fmt.Println("Nessuna parola compatibile")
	} else {
		fmt.Println("Altre parole: ")
		for i, w := range words {
			fmt.Print(w)
			if i < len(words)-1 {
				fmt.Print(", ")
			} else {
				fmt.Println(".")
			}
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
		return commonDictionary1000
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
		convertNW(numbers[randomNumber])
	}
}

/*func gamePhraseNW() {

	var numbers []string

	l := chooseLevelGamePhrase()

	numberPhrase := stringWithCharset(l, charsetNumber)

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
		convertNW(numbers[randomNumber])
	}
}
*/

func numberToPhrases(number string, lenMin int) [][]string {

	SS := getAllCombinations(number, lenMin)

	log.Info("SS len %v", len(SS))

	var SSfiltered [][]string

	for i := 0; i < len(SS); i++ { // per tutte le combinazioni

		ok := true
		var phrase []string

		for j := 0; j < len(SS[i]); j++ { //per tutti i chank di ogni combinazione

			if commonDictionary1000[SS[i][j]] != nil {
				words := commonDictionary1000[SS[i][j]][0]
				for k := 1; k < len(commonDictionary1000[SS[i][j]]); k++ {
					words = words + "," + commonDictionary1000[SS[i][j]][k]
				}
				phrase = append(phrase, words)
			} else {
				ok = false
			}
		}

		if ok {
			SSfiltered = append(SSfiltered, phrase)
		}
	}

	return SSfiltered

}

func getAllCombinations(s string, lenMin int) [][]string {
	SS := [][]string{{s}}
	SS = append(SS, getSplit(s, lenMin)...)
	return SS
}

/*func getSplit(s string) [][]string {
	var SS [][]string
	if len(s) > 1 {
		for i := 1; i < len(s); i++ {
			s1 := s[:i]
			s2 := s[i:]
			SS = append(SS, []string{s1, s2})
			SS2 := getSplit(s2)
			if SS2 != nil {
				for j := 0; j < len(SS2); j++ {
					SS2[j] = append([]string{s1}, SS2[j]...)
				}
				SS = append(SS, SS2...)
			}
		}
		return SS
	} else {
		return nil
	}

}*/

func getSplit(s string, lenMin int) [][]string {

	var SS [][]string
	if len(s) > lenMin {
		for i := lenMin; i < len(s)-(lenMin-1); i = i + lenMin {
			s1 := s[:i]
			s2 := s[i:]
			SS = append(SS, []string{s1, s2})
			SS2 := getSplit(s2, lenMin)
			if SS2 != nil {
				for j := 0; j < len(SS2)-1; j = j + 2 {
					SS2[j] = append([]string{s1}, SS2[j]...)
				}
				SS = append(SS, SS2...)
			}
		}
		return SS
	} else {
		return nil
	}

}

func chooseLevelGamePhrase() int {
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
		randomNumber := rand.Intn(5)
		return randomNumber
	case extremeLevel:
		randomNumber := rand.Intn(5) + 5
		return randomNumber
	default:
		randomNumber := rand.Intn(10) + 10
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

/*func makeDictionaryy() {

	data, err := ioutil.ReadFile("parole.txt")
	if err != nil {
		fmt.Println(err)
	}

	dataString := string(data)
	dataSlice := strings.Split(dataString,",")
	fmt.Println(dataSlice)
	var nwDataSlice []string
	for _,s := range dataSlice {
		s = strings.ReplaceAll(s, " ","")
		s = strings.ReplaceAll(s, "\n","")
		if s != ""{
			nwDataSlice = append(nwDataSlice, s)
		}
	}

	str := ""
	for _,s := range nwDataSlice {
		str= str+s+"\n"
	}

	mydata := []byte(str)

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("parolecomuni.txt", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

}
*/
