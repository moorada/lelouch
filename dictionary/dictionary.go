package dictionary

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/evilsocket/islazy/log"
)

const pathCompleteDictionaryJSON = "/dizionari/completeDictionary.json"
const pathcommonDictionaryJSON = "/dizionari/commonDictionary.json"

const pathCommonDictionary = "/dizionari/parolecomuni.txt"
const pathCompleteDictionary = "/dizionari/tutteleparole.txt"

type Words []string
type AS map[string]Words

func MakeDictionaries() {
	makeDictionary(pathCommonDictionary, pathcommonDictionaryJSON)
	makeDictionary(pathCompleteDictionary, pathCompleteDictionaryJSON)
}

func makeDictionary(path string, jsonPath string) {

	var words []string

	path, err := filepath.Abs(filepath.Dir(os.Args[0]) + path)
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Open(path)
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

	jsonPath, err = filepath.Abs(filepath.Dir(os.Args[0]) + jsonPath)
	if err != nil {
		log.Fatal(err.Error())
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
	path, err := filepath.Abs(filepath.Dir(os.Args[0])+path)
	if err != nil {
		log.Fatal(err.Error())
	}

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
	return getDictionary(pathcommonDictionaryJSON)
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
