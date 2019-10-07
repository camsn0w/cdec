package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

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
	
}
