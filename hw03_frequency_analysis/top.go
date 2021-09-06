package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const count = 10

var rePunctuation = regexp.MustCompile(`\p{P}`)

type Word struct {
	Value string
	Count int
}

type wordList []Word

func (list *wordList) AppendList(text map[string]int) {
	for value, count := range text {
		*list = append(*list, Word{value, count})
	}
}

func (list *wordList) Sort() {
	sortedStruct := *list
	sort.Slice(sortedStruct, func(i, j int) bool {
		if sortedStruct[i].Count == sortedStruct[j].Count {
			return sortedStruct[i].Value < sortedStruct[j].Value
		}
		return sortedStruct[i].Count > sortedStruct[j].Count
	})
}

func clearText(text string) string {
	text = strings.ToLower(text)
	text = rePunctuation.ReplaceAllString(text, "")
	return text
}

func calcWorld(text string) map[string]int {
	temp := strings.Fields(text)
	words := make(map[string]int, len(temp))
	for _, val := range temp {
		words[val]++
	}
	return words
}

func getResult(list wordList, wordsLen int) []string {
	result := make([]string, 0, count)
	for i := 0; (i < count) && (i < wordsLen); i++ {
		result = append(result, list[i].Value)
	}
	return result
}

func sortStruct(words map[string]int) []string {
	wordsLen := len(words)
	list := make(wordList, 0, wordsLen)
	list.AppendList(words)
	list.Sort()

	return getResult(list, wordsLen)
}

func Top10(s string) []string {
	clearText := clearText(s)
	words := calcWorld(clearText)
	return sortStruct(words)
}
