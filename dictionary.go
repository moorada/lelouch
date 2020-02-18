package main

//Questo modulo serve per creare un nuovo dizionario associato ai numeri.
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/evilsocket/islazy/log"
)

const IndexedDictionaryJSON = "indexedDictionary.json"

type Words []string
type AS map[string]Words

func makeDizionario() {

	files, err := filepath.Glob("paroleitaliane/*")
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	fmt.Println("Creando dizioniario da ", len(files), "file.")

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
		if len(w) > 26 {
			continue
		}
		index := wordToNumber(w)

		elem, ok := m[index]
		if !ok {
			m[index] = Words{w}
		} else {
			if !elem.contain(w) {
				m[index] = append(elem, w)
			}
		}
	}

	f, err := os.Create(IndexedDictionaryJSON)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	defer f.Close()
	_, err = f.WriteString(m.toJsonString())
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}

func getIndexedDictionary() AS {
	file, err := os.Open(IndexedDictionaryJSON)
	defer file.Close()
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	var as AS
	err = json.Unmarshal(byteValue, &as)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	return as
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
