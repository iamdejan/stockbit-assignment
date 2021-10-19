// Q4 answer
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func createreqTable(s string) [26]int {
	var arr [26]int
	for _, value := range strings.ToLower(s) {
		idx := int(value) - 'a'
		arr[idx] += 1
	}
	return arr
}

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anagramMap := make(map[[26]int][]string)
	for _, value := range input {
		freqTable := createreqTable(value)
		anagramMap[freqTable] = append(anagramMap[freqTable], value)
	}
	var output [][]string
	for _, value := range anagramMap {
		output = append(output, value)
	}
	result, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("Error when parsing to JSON: %s", err.Error())
		return
	}
	fmt.Println(string(result))
}
