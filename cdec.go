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
	//aheadReader := reader
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
	/*for {
	// size of 4kb + 1 to see if the last byte is a space or not
	//bufAhead := make([]byte, 4070)
	inpt := make([]byte, 0)
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
			readAhead := reader
			oneAhead := make([]byte, 1)

			//print("Case (true = enter loop): " + " " + string(oneAhead[0]) != " "+"\n")

			if string(oneAhead[0]) != " " {
				_, err3 := readAhead.Read(oneAhead)
				if err3 != nil {
					if err3 == io.EOF {
						break
					}
					fmt.Printf("%s", err3)
					os.Exit(0)
				}

			}
			temp := swoleBuf
			inpt = append(inpt, oneAhead[0])
			reader = readAhead
			if string(oneAhead[0]) == " " {
				print("String that it is at: " + string(temp[len(temp)-10:]) + "\n")
				print("What is being added: " + string(inpt[len(inpt)-10:])+"\n")
				break
			}

		}
	}*/
	toWords := strings.Fields(string(inpt))
	for index, vals := range toWords {

		count := wordCount[vals]
		wordCount[toWords[index]] = count + 1
		numWords++

	}

	freqToRet := FreqCount{wordCount, numWords}
	return freqToRet, err
}

/*func ScanFreqsFromReader(reader io.Reader) (FreqCount, error) {
	wordCount := make(map[string]float64)
	numWords := 0
	var err error
	scannedReader := bufio.NewScanner(reader)
	//reader.Read()
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
}*/
func FreqOfWord(word string, freqMap FreqCount) float64 {
	freq := (freqMap.Freqs[word] / float64(freqMap.numWords)) * 100
	fmt.Printf("%f", freqMap.Freqs[word])
	return freq
}

type FreqCount struct {
	Freqs    map[string]float64
	numWords int
}
