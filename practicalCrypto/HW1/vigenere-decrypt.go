/*
 * Written by Eric Rothman
 * This program recieves a key of capital letters and the name of a text file as input and outputs the text in the text file decrypted using the Vigenere cipher and the key.
 */


package main

import (
		"fmt"
		"os"
		"io/ioutil"
	   )

/*
 *This function converts both capital and lower case letters to be expressed as the alphabet starting with a=0.
 *In addition it sanitizes the input to remove punctuation and whitespace so only the letters remain.
 */
func convertToAlph(char uint8) int {
	charValue := int(char)
    if (charValue >= 65 && charValue <= 90){
		return charValue - 65
	} else if (charValue >= 97 && charValue <= 122) {
		return charValue - 97
	} else {
		return 100
	}
}

/*
 *This function checks the input argument for the key and makes sure it is valid
 */
func checkArguments(key string) int {
	if (len(key) > 32) {
		fmt.Println("Wrong Argument: The key cannot be more than 32 charactors long.")
		return 1
	}
	for a := 0; a < len(key); a++ {
		if (key[a] < 65 || key[a] > 90) {
			fmt.Println("Wrong Argument: The first argument must be the key and must consist of only capital letters of the english alphabet.")
			return 1
		}
	}
	return 0
}

/*
 *This function returns the text from the file given as an argument.
 */
func getFileInput(fileName string) string {
	input, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		panic(err)
	}
	return string(input)
}

/*
 *This function prints out the decrypted input text using the vigenere cipher and the key given.
 */

func decryptMessage(key, input string) string {
	output := ""
	keyPlacer := 0
	for a:= 0; a < len(input); a++ {
		letter := convertToAlph(input[a])
		if (letter != 100) {
			newLetter := (((26 + letter - convertToAlph(key[keyPlacer])) % 26) + 65)
			output += string(newLetter)
		}
		keyPlacer = ((keyPlacer + 1) % len(key))
	}
	fmt.Println(output)
	return output
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {

	args := os.Args[1:]
	if (len(args) != 2) {
		fmt.Println("Wrong Arguments: There need to be three command line arguments.")
		fmt.Println("The first argument is the key for the encryption that must consist of all uppercase letters.")
		fmt.Println("The second argument is the file name for encryption or decrpytion.")
		os.Exit(1)
	}

	key := args[0]
	fileName := args[1]

	if (checkArguments(key) == 1) {
		os.Exit(1)
	}

	input := getFileInput(fileName)
	decryptMessage(key, input)
}
