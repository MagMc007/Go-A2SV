package main
import (
	"fmt"
)

// make sure char is a letter
func isLetter(c byte) bool {
	// caps and smalls
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}


// make it lowercase
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func isPalindrome(s string) bool {
	// two pointers
	i := 0
	j := len(s) - 1

	// iterate and check
	for i < j {
		// check if they are both alpha
		if !isLetter(s[i]) {
			i++
		}

		if !isLetter(s[j]) {
			j--
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i ++
		j --
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("abaa"))
}