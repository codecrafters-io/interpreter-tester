package internal

import (
	"math/rand"
	"strings"
)

func repeatSlice(slice []string, n int) []string {
	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, slice...)
	}
	return ret
}

func randomSelection(count int, characters []string, joinWith string) string {
	selectedCharacters := make([]string, count)
	for i := 0; i < count; i++ {
		selectedCharacters[i] = characters[rand.Intn(len(characters))]
	}
	return strings.Join(selectedCharacters, joinWith)
}

func randomStringFromCharacters(totalLength int, characters []string) string {
	characterList := make([]string, totalLength)
	for i := 0; i < totalLength; i++ {
		characterList[i] = characters[rand.Intn(len(characters))]
	}

	shuffledString := strings.Join(characterList, "")
	return shuffledString
}
