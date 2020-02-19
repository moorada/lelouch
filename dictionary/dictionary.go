package dictionary

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/evilsocket/islazy/log"
)

const pathCompleteDictionaryJSON = "dizionari/completeDictionary.json"
const pathCommonDictionaryJSON = "dizionari/commonDictionary.json"
const pathSimpleDictionaryJSON = "dizionari/simpleDictionary.json"

const pathCommonDictionary = "dizionari/parolecomuni.txt"
const pathCompleteDictionary = "dizionari/tutteleparole.txt"
const pathSimpleDictionary = "dizionari/parole1-99.txt"


type Words []string
type AS map[string]Words

func MakeDictionaries() {
	makeDictionary(pathCommonDictionary, pathCommonDictionaryJSON)
	makeDictionary(pathCompleteDictionary, pathCompleteDictionaryJSON)
	makeDictionary(pathSimpleDictionary, pathSimpleDictionaryJSON)

}

func makeDictionary(path string, jsonPath string) {

	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	var words []string

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		morewords := strings.Split(string(byteValue), "\n")
		words = append(words, morewords...)
		file.Close()
	}

	m := make(AS)

	for _, w := range words {
		w = strings.ToLower(w)
		index := WordToNumber(w)

		elem, ok := m[index]
		if !ok {
			m[index] = Words{w}
		} else {
			if !elem.contain(w) {
				m[index] = append(elem, w)
			}
		}
	}

	f, err := os.Create(jsonPath)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	defer f.Close()
	_, err = f.WriteString(m.toJsonString())
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}

func getDictionary(path string) (error, AS) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err, nil
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err, nil
	}

	var as AS
	err = json.Unmarshal(byteValue, &as)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	return nil, as
}

func GetCommonDictionary() (error, AS) {
	return getDictionary(pathCommonDictionaryJSON)
}

func GetSimpleDictionary() (error, AS) {
	return getDictionary(pathSimpleDictionaryJSON)
}

func GetCompleteDictionary() (error, AS) {
	return getDictionary(pathCompleteDictionaryJSON)
}

func (ws Words) contain(word string) bool {
	for _, w := range ws {
		if word == w {
			return true
		}
	}
	return false
}

func (as AS) toJsonString() string {

	// Marshal the map into a JSON string.
	asData, err := json.Marshal(as)
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	jsonStr := string(asData)
	return jsonStr
}

func (as AS) makeByString(s string) {

	err := json.Unmarshal([]byte(s), &as)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}
