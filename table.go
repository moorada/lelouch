package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/evilsocket/islazy/log"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
)

func showTableWordSet(number string) {

	phrases := numberToWordSet(number)
	var rows [][]string
	for _, p := range phrases {
		newp := p
		for i, _ := range newp {
			newLine := 35
			for len(newp[i]) > newLine {
				newp[i] = newp[i][:newLine] + "\n" + newp[i][newLine:]
				newLine += newLine
			}
		}
		rows = append(rows, newp)
	}
	maxLen := 0
	for _, r := range rows {
		if len(r) > maxLen {
			maxLen = len(r)
		}
	}

	for i, _ := range rows {
		for len(rows[i]) < maxLen {
			rows[i] = append(rows[i], "--")
		}
	}

	partialRows := rows
	if len(rows) > 5 {
		partialRows = rows[:5]
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(partialRows)
	table.Render()

	prompt := promptui.Select{
		Label: "Seleziona",
		Items: []string{MostraTabellaCompleta, Continua, MenuPrincipale},
	}

	if len(rows) == len(partialRows) {
		prompt = promptui.Select{
			Label: "Seleziona",
			Items: []string{Continua, MenuPrincipale},
		}
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error: %s", err)
		return
	}
	switch result {
	case MostraTabellaCompleta:
		wordsCompatible(number)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.AppendBulk(rows)
		table.Render()
	case MenuPrincipale:
		start()
	default:
		return
	}

}

func wordsCompatible(number string)  {

	wordsCompatible := "Singole parole : "
	words := completeDictionary[number]
	if len(words) != 0 {
		for i, w := range words {
			wordsCompatible = wordsCompatible + w
			if i < len(words)-1 {
				wordsCompatible = wordsCompatible + ", "
			}
		}
		fmt.Println(wordsCompatible)
	}

}

func numberToWordSet(number string) [][]string {

	SS := getAllCombinations(number)
	var SSfiltered [][]string

	for i := 0; i < len(SS); i++ {

		ok := true
		var wordSet []string

		for j := 0; j < len(SS[i]) && ok; j++ {

			if commonDictionary[SS[i][j]] != nil {
				words := commonDictionary[SS[i][j]][0]
				for k := 1; k < len(commonDictionary[SS[i][j]]); k++ {
					words = words + ", " + commonDictionary[SS[i][j]][k]
				}
				wordSet = append(wordSet, words)
			} else {
				ok = false
			}
		}
		if ok {
			SSfiltered = append(SSfiltered, wordSet)
		}
	}

	return SSfiltered

}

func getAllCombinations(s string) [][]string {
	SS := [][]string{{s}}
	SS = append(SS, getSplit(s)...)
	sort.Slice(SS, func(i, j int) bool {
		return len(SS[i]) < len(SS[j])
	})
	return SS
}

func getSplit(s string) [][]string {
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

}
