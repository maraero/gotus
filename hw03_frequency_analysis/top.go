package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const maxNumber = 10

func Top10(input string) []string {
	freqByWords := calcWords(input)
	wordFreqPairs := getWordFrequencyPairs(freqByWords)
	return getMostFrequentWordsByLimit(wordFreqPairs, maxNumber)
}

func calcWords(text string) map[string]int {
	freqByWords := make(map[string]int)
	words := strings.Fields(text)
	for _, w := range words {
		freqByWords[w]++
	}
	return freqByWords
}

type wordFreqPair struct {
	freq int
	word string
}
type frequencyComparator []wordFreqPair

func (a frequencyComparator) Len() int      { return len(a) }
func (a frequencyComparator) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a frequencyComparator) Less(i, j int) bool {
	if a[i].freq == a[j].freq {
		return a[i].word > a[j].word
	}
	return a[i].freq < a[j].freq
}

func getWordFrequencyPairs(freqByWords map[string]int) []wordFreqPair {
	wordFreqPairs := make([]wordFreqPair, 0, len(freqByWords))
	for word, freq := range freqByWords {
		wordFreqPairs = append(wordFreqPairs, wordFreqPair{freq: freq, word: word})
	}
	sort.Sort(frequencyComparator(wordFreqPairs))
	return wordFreqPairs
}

func getMostFrequentWordsByLimit(list []wordFreqPair, limit int) []string {
	result := make([]string, 0, limit)
	for i := len(list) - 1; i >= 0 && limit > 0; i-- {
		result = append(result, list[i].word)
		limit--
	}
	return result
}
