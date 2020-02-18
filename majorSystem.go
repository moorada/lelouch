package main

import (
	"regexp"
	"strings"
)

const ParoleSenzaFonetiche = "99999999"

type Figure struct {
	number    string
	phonetics []string
}

type Adjustament struct {
	original string
	variants []string
}

type Exception struct {
	number string
	words  []string
}

var Exceptions = []Exception{
	{"756", []string{"glig", "glic"}},
	{"05", []string{"siglio"}},
	{"405", []string{"rsigli"}},
	{"205", []string{"msigli"}},
	{"05", []string{"ssigli"}},
	{"075", []string{"sigli"}},
}

type MS [10]Figure

type AD []Adjustament

var Vowel = []string{
	"a",
	"e",
	"i",
	"o",
	"u",
}

var ItalianAdjustments = AD{
	{"a", []string{"á", "à",}},
	{"e", []string{"è", "é"}},
	{"i", []string{"ì", "í", "y", "j"}},
	{"o", []string{"ò", "ó"}},
	{"u", []string{"ù", "ú"}},
	{"ks", []string{"x"}},
}

var ItalianMS = MS{
	{"0", []string{"ss", "zz", "sci", "sce", "s", "z"}},
	{"1", []string{"tt", "dd", "t", "d"}},
	{"2", []string{"nn", "gn", "n"}},
	{"3", []string{"mm", "m"}},
	{"4", []string{"rr", "r"}},
	{"5", []string{"gli", "ll", "l"}},
	{"6", []string{"cci", "ci", "cce", "ce", "ggi", "gi", "gge", "ge"}},
	{"7", []string{"cc", "gg", "cq", "kk", "q", "k", "c", "g",}},
	{"8", []string{"ff", "vv", "f", "v"}},
	{"9", []string{"pp", "bb", "p", "b"}},
}

func wordToNumber(word string) string {

	word = strings.ToLower(word)
	for _, a := range ItalianAdjustments {
		for _, v := range a.variants {
			word = strings.ReplaceAll(word, v, a.original)
		}
	}

	for _, a := range Exceptions {
		for _, v := range a.words {
			word = strings.ReplaceAll(word, v, a.number)
		}
	}

	for _, a := range ItalianMS {
		for _, p := range a.phonetics {
			word = strings.ReplaceAll(word, p, a.number)
		}
	}
	for _, v := range Vowel {
		word = strings.ReplaceAll(word, v, "")
	}

	if word == "" {
		return ParoleSenzaFonetiche
	}

	re := regexp.MustCompile("[^0-9]+")
	word = re.ReplaceAllString(word, "")

	if word == "" {
		return ParoleSenzaFonetiche
	}

	return word
}
