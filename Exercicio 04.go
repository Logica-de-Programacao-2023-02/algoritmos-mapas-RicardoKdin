package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sort.StringSlice(r))
	return string(r)
}

func groupAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		if _, exists := anagramGroups[sortedWord]; !exists {
			anagramGroups[sortedWord] = []string{word}
		} else {
			anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
		}
	}

	return anagramGroups
}

func main() {
	words := []string{"listen", "silent", "heart", "earth", "listen", "foo", "bar", "ball"}

	anagramGroups := groupAnagrams(words)

	for _, group := range anagramGroups {
		if len(group) > 1 {
			fmt.Println(group)
		}
	}
}
