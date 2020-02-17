package main

type Figure struct {
	char      rune
	phonetics []string
}

type MS [10]Figure
//TODO Risolvere il problema della "gl", a volte si legge palatale, a volte gutturale.

var ItalianMS = MS{
	{'0', []string{"s", "z", "sc","sc"}},
	{'1', []string{"t", "d"}},
	{'2', []string{"n", "gn"}},
	{'3', []string{"m"}},
	{'4', []string{"r"}},
	{'5', []string{"l", "gl"}},
	{'6', []string{"ce", "ci", "ge", "gi", "j"}},
	{'7', []string{"c","g","ca", "co", "cu", "ch", "ga", "go", "gu", "gh", "q","k"}},
	{'8', []string{"f", "v"}},
	{'9', []string{"p", "b"}},
}

var Vowels = []rune{'a', 'e', 'i', 'o', 'u', 'à', 'é', 'í', 'ó', 'ú', 'à', 'è', 'ì', 'ò', 'ù'}

func IsVowel(vowel rune) bool {
	for _, c := range Vowels {
		if vowel == c {
			return true
		}
	}
	return false
}