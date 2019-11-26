package cdec

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func ScanFreqsFromReader(reader io.Reader) (FreqCount, error) {
	err := error(nil)
	if reader == nil {
		err := errors.New("file is empty")
		freqErr := FreqCount{nil, 0}
		return freqErr, err

	}
	wordCount := make(map[string]float64)
	numWords := 0
	swoleBuf := make([]byte, 4069)
	inpt := make([]byte, 0)
	oneAhead := make([]byte, 1)

	for {
		_, err := reader.Read(swoleBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("%s", err)
			os.Exit(0)
		}
		inpt = append(inpt, swoleBuf...)
		if string(swoleBuf[4068]) != " " {
			for {
				_, err2 := reader.Read(oneAhead)
				if err2 != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("%s", err)
					os.Exit(0)
				}
				inpt = append(inpt, oneAhead...)
				if string(oneAhead) == " " {
					break
				}
			}

		}

	}

	toWords := strings.Fields(string(inpt))
	for index, vals := range toWords {

		count := wordCount[vals]
		wordCount[toWords[index]] = count + 1
		numWords++

	}

	freqToRet := FreqCount{wordCount, numWords}
	return freqToRet, err
}

func FreqOfWord(word string, freqMap FreqCount) float64 {
	freq := (freqMap.Freqs[word] / float64(freqMap.numWords)) * 100
	fmt.Printf("%f", freqMap.Freqs[word])
	return freq
}

type FreqCount struct {
	Freqs    map[string]float64
	numWords int
}
