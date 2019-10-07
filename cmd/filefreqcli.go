package main

import (
	"bufio"
	"cdec"
	"flag"
	"fmt"
	"os"
	"strings"
)

func fileRead(fileName string) *os.File {

	print("Filename: " + fileName + "\n")
	inFile, err := os.Open(fileName)
	if err != nil {
		print("ERROR during fileRead ")
		print(err.Error())

		return nil
	}
	return inFile
}
func userInput() (string, error) {
	fmt.Printf("%s", "Enter a word to search for: ")
	inpt := bufio.NewReader(os.Stdin)
	userString, userStringError := inpt.ReadString('\n')
	if userStringError != nil {
		errorWithInput(userStringError)
	}
	userString = strings.Replace(userString, "\n", "", -1)
	return userString, userStringError
}
func errorWithInput(userStringError error) string {
	for {
		fmt.Printf("Error: %s; %s\n", userStringError, "; Type 'y' to try again or 'q' to quit")
		inpt, err := userInput()
		if err != nil {
			print("Input must be either 'y' or 'q'\n")
			continue
		} else {
			if inpt == "y" {
				fmt.Printf("%s", "Enter a word to search for: ")
				word, err := userInput()
				if err != nil {
					fmt.Printf("Error: %s", err)
					continue
				} else {
					return word
				}
			}
		}

	}
}
func main() {
	word, _ := userInput()
	flag.Parse()
	fileName := flag.Arg(0)
	inputFile := fileRead(fileName)
	var reader = bufio.NewReader(inputFile)
	freqData, err := cdec.ScanFreqsFromReader(reader)
	if err != nil {

		fmt.Printf("%s", err)
		inputFile.Close()
		os.Exit(1)

	}
	relativeFreq := cdec.FreqOfWord(word, freqData)
	fmt.Printf("Frequency of %s is %f", word, relativeFreq)

}
