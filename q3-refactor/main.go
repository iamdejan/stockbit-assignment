package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
}

// Q3 answer
func findRefactor(s string) string {
	var result []rune
	isRecording := false
	for _, c := range s {
		if c == '(' {
			isRecording = true
			continue
		}

		if c == ')' {
			isRecording = false
			break
		}

		if isRecording {
			result = append(result, c)
		}
	}
	return string(result[:len(result)-1])
}

func main() {
	s := "(lore)m ipsum dolor (sit) amet"
	fmt.Println(findFirstStringInBracket(s))
	fmt.Println(findRefactor(s))
}
