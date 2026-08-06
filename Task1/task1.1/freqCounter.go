package main
import (
	"fmt"
	"strings"
	"unicode"
)

// make a function to check if a string is all alpha or not
func checkAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, ch := range s {
		if !unicode.IsLetter(ch) {
			return false
		}
	}
	return true
}


func counter(word string) map[string]int {
	// declare the map
	freq := make(map[string]int)

	// split the word
	word_list := strings.Split(word, " ")

	for i := 0; i < len(word_list); i++ {
		// leave out any punctuation
		word_list[i] = strings.Trim(word_list[i], ".,!?;:\"'()[]{}")

		// make them case insensitive
		word_list[i] = strings.ToLower(word_list[i])

		// ignore it if it is punctuation
		if checkAlpha(word_list[i]) {
			freq[word_list[i]] += 1
		}
	}

	return freq
}

func main() {
	word := "Hello I am BayMax"

	fmt.Println(counter(word))
}