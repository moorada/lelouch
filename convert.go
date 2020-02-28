package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/moorada/lelouch/dictionary"
)

func convert() {
	for {
		fmt.Println(Converter)
		var response string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			response = scanner.Text()
		}
		re := regexp.MustCompile("[^a-zA-Z0-9_]+")
		response = re.ReplaceAllString(response, "")
		if _, err := strconv.Atoi(response); err == nil {
			convertNW(response)
		} else {
			convertWN(response)
		}
	}
}

func convertNW(s string) {
	showTableWordSet(s)
}

func convertWN(word string) {
	fmt.Println(dictionary.WordToNumber(word))
}

