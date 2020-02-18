package main

type Figure struct {
	char      rune
	phonetics []string
}

type Adjustament struct {
	original string
	variants []string
}

type MS [10]Figure

type AD []Adjustament

//TODO Risolvere il problema della "gli", gutturale solo se all'inizio della parola

var ItalianAdjustments = AD{
	{"a", []string{"á", "à",}},
	{"e", []string{"è", "é"}},
	{"i", []string{"ì", "í", "y", "j"}},
	{"o", []string{"ò", "ó"}},
	{"u", []string{"ù", "ú"}},
	{"ks", []string{"x"}},
}

var ItalianMS = MS{
	{'0', []string{"ss", "zz", "sci", "sce", "s", "z"}},
	{'1', []string{"t", "d"}},
	{'2', []string{"nn", "gn", "n"}},
	{'3', []string{"mm", "m"}},
	{'4', []string{"rr", "r"}},
	{'5', []string{"gli", "ll", "l"}},
	{'6', []string{"cci", "ci", "cce", "ce", "ggi", "gi", "gge", "ge"}},
	{'7', []string{"cc", "gg", "cq", "kk", "q", "k", "c", "g",}},
	{'8', []string{"ff", "vv", "f", "v"}},
	{'9', []string{"pp", "bb", "p", "b"}},
}
