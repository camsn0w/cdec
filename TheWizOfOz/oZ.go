package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

)



func fileRead(fileName string) *os.File {

	print("Filename: " + fileName +"\n")
	inFile , err := os.Open(fileName)

	//T , err := os.Open("testFile.txt")
	if err != nil {
		print("ERROR during fileRead ")
		print(err.Error())

		return nil
	}
	//defer inFile.Close()
	return inFile
}
func fileProcess(theFile *os.File) {

}
func userInput() (string,error) {
	inpt := bufio.NewReader(os.Stdin)
	userString, userStringError := inpt.ReadString('\n')
	if userStringError != nil {
		fmt.Println(userStringError)
	}
	userString = strings.Replace(userString, "\n","",-1)
	return userString,userStringError
}



func main() {
	//TODO test new userInput() method
	flag.Parse()
	thing := 5

	//Take filename as input from user:
	//inpt := bufio.NewReader(os.Stdin)
	fmt.Print("Enter filename: ")
	//fileName, nameError := inpt.ReadString('\n')
	/*if nameError != nil {
		fmt.Println(nameError)
	}*/
	fileName, _ := userInput()
	/*fileName = strings.Replace(fileName, "\n","",-1)*/

	//
	fileName = "WizardOfOz.txt"
	//

	//Take word input from user
	fmt.Print("Word: ")
	/*wordInput := bufio.NewReader(os.Stdin)
	word, err := wordInput.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	word = strings.Replace(word, "\n","",-1)*/
	word, _:= userInput()

	//
	word = "Scarecrow"
	//
	
	//fileName := flag.Arg(0) //old, this was to use the program args for filename
	inputFile := fileRead(fileName)
	var reader = bufio.NewReader(inputFile)
	scannedReader := bufio.NewScanner(reader)
	wordCount := make(map[string]int64)
	numWords := 0

	for scannedReader.Scan(){

		words := strings.Fields(scannedReader.Text())

		//fmt.Println(words)
		c := 0

		for index, vals := range words {

			count := wordCount[vals]
			wordCount[words[index] ]= count+1
			numWords = numWords +1
			c++
		}

	}
	//var floatedWordCount float64 = float64(wordCount[word])
	//wordFreq ,float64,_ = wordCount[word]
	//var wordFreqFloated = float64(wordFreq)

	println("Numwords: ")
	fmt.Println(numWords)
	 var freq float64 = float64(wordCount[word]) / float64(numWords) * 100
	//s := strconv.FormatInt(wordCount[word],10)
	//var freq, _ (float64,bool) = wordCount[word];
	fmt.Println("Number of times the word " + word + " was found: ")
	fmt.Println(wordCount[word])
	fmt.Println("Frequency of " + word + " in the text: ")
	 fmt.Print(freq)
	fmt.Print("%")

	//fmt.Printf("%F", float64(float64(wordCount[word])/float64(numWords))*100)
	 //fmt.Printf("%F", float64(float64(wordCount[word])/float64(numWords))*100)
	//print(freq*100)

}
