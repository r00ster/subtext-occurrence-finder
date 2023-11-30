package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// findOccurrences finds and prints the positions of occurrences of subtext in textToSearch
func findOccurrences(textToSearch, subtext string, positions *string) {
	var textLower string = strings.ToLower(textToSearch)
	var subtextLower string = strings.ToLower(subtext)

	var positionsArr []string
	for i := 0; i < len(textLower); {
		var index int = strings.Index(textLower[i:], subtextLower)
		if index == -1 {
			break
		}

		var wordPos int = i + index + 1;
		positionsArr = append(positionsArr, strconv.Itoa(wordPos))

		i = i + index + 1
	}

	if len(positionsArr) > 0 {
		*positions = strings.Join(positionsArr, ", ")
	}
}

func main() {
	var positions string = ""

	scanner := bufio.NewScanner(os.Stdin)
	// Prompt for the first string
	fmt.Print("Enter the text to search: ")
	scanner.Scan()
	var textToSearch string = scanner.Text()

	// Prompt for the second string
	fmt.Print("Enter the subtext: ")
	scanner.Scan()
	var subtext string = strings.TrimSpace(scanner.Text())

	// Find and print occurrences
	findOccurrences(textToSearch, subtext, &positions)

	if (positions != "") {
		fmt.Printf("\"%s\"", positions)
	}
}
