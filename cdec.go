package cdec

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func ScanFreqsFromReader(reader io.Reader) (FreqCount, error) {
	wordCount := make(map[string]float64)
	numWords := 0
	var err error
	scannedReader := bufio.NewScanner(reader)
	if reader == nil {
		err = errors.New("file is empty")
		freqErr := FreqCount{wordCount, 0}
		return freqErr, err
	}

	for scannedReader.Scan() {

		words := strings.Fields(scannedReader.Text())

		//fmt.Println(words)
		c := 0

		for index, vals := range words {

			count := wordCount[vals]
			wordCount[words[index]] = count + 1
			numWords = numWords + 1
			c++
		}

	}
	freqToRet := FreqCount{wordCount, numWords}
	return freqToRet, err
}
func FreqOfWord(word string, freqMap FreqCount) float64 {
	//	 var freq float64 = float64(wordCount[word]) / float64(numWords) * 100
	freq := (freqMap.Freqs[word] / float64(freqMap.numWords)) * 100
	return freq
}

type FreqCount struct {
	Freqs    map[string]float64
	numWords int
}
