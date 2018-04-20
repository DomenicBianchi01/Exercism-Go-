package isogram

import "strings"
import "unicode"

func IsIsogram(word string) bool {
	var foundCharacters []string

	for _, character := range word {
		if !unicode.IsLetter(character) {
			continue
		}
		if !CharacterExists(string(character), foundCharacters) {
			foundCharacters = append(foundCharacters, strings.ToLower(string(character)))
		} else {
			return false
		}
	}

	return true
}

func CharacterExists(searchCharacter string, foundCharacters []string) bool {
	for _, character := range foundCharacters {
		if strings.ToLower(searchCharacter) == character {
			return true
		}
	}
	return false
}
