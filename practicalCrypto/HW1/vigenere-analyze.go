/*
 * Written by Eric Rothman
 * This program accomplishes parts 2 and 3 of the assignmen.
 * So it recieves as an argument the name of a text file that includes text encrypted using a vigenere cipher.
 * It then outputs the best guess for the length of the key of the cipher, from 1-20, and the best guess for the key itself.
 */

package main

import (
		"fmt"
		"os"
		"strings"
		"math"
		"io/ioutil"
	   )

var NUMBERCHECKED = 20

var englishAlphFreq = [26]float64{8.04, 1.54, 3.06, 3.99, 12.51, 2.3, 1.96, 5.49, 7.26, 0.16, 0.67, 4.14, 2.53, 7.09, 7.60, 2.0, 0.11, 6.12, 6.54, 9.25, 2.71, 0.99, 1.92, 0.19, 1.73, 0.09}

/*
 * This function returns the input from the file whose name was given as an argument
 */
func getInput(fileName string) string {
	input, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		panic(err)
	}
	return string(input)
}

/*
 * This function converts upper and lower case letters to be between 0 and 25 while also sanatizing to remove non-letter charactors.
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
 * This function receives the text from the inputed file and then does an index of coincidence attack to figure out the most likely key length between 1 and 20.
 */
func indexOfCoincidence(fileInput string) int {
	fileSize := len(fileInput)
	var placement [20]int
	for i := 3; i <= NUMBERCHECKED; i++ {
		coincidences := 0
		for j := 0; j < len(fileInput) - i; j++ {
			if (fileInput[j] == fileInput[j + i]) {
				coincidences++
			}
		}
		placement[i - 1] = coincidences
	}
	for b:=0; b < NUMBERCHECKED; b++ {
		if(float64(placement[b]) > 0.066*float64(fileSize) - 10) {
			return b + 1
		}
	}
	return 1
}

/*
 * This function returns the cosine similarity between the vector of the percentage use of the english alphabet letters shifted by offset positions and the values in the freqArray.
 */
func similarity(offset int, freqArray [26]int) float64 {
	dotProduct := 0.0
	alphabetSize := 0.0
	freqArraySize := 0.0
	for alphSpot := 0; alphSpot < 26; alphSpot++ {
		alphabetSize += englishAlphFreq[alphSpot] * englishAlphFreq[alphSpot]
		freqArraySize += float64(freqArray[alphSpot]) * float64(freqArray[alphSpot])
		dotProduct += englishAlphFreq[(26 + alphSpot - offset) % 26] * float64(freqArray[alphSpot])
	}
	return dotProduct/(math.Sqrt(alphabetSize)*math.Sqrt(freqArraySize))
}

/*
 * This function returns the string consisting of every keyLength spots in the inputText starting at the startSpot spot.
 */
func makeString(startSpot, keyLength int, inputText string) string {
	output:= ""
	for i := startSpot; i < len(inputText); i += keyLength {
		output += string(inputText[i])
	}
	return output
}

/*
 * This function returns an array where each spot corresponds to the frequency that letter appears in the input text.
 */
func getStringFrequencies(inputString string) [26]int {
	output := [26]int{}
	for i := 0; i < len(inputString); i++ {
		spot := convertToAlph(inputString[i])
		if (spot != 100) {
			output[spot]++
		}
	}
	return output
}

/*
 * This function returns the spot in an array of floats which is the largest.
 */
func findMax(numbers []float64) int {
	maxSpot := 0
	maxValue := 0.0
	for i := 0; i < len(numbers); i++ {
		if (numbers[i] > maxValue) {
			maxValue = numbers[i]
			maxSpot = i
		}
	}
	return maxSpot
}

/*
 * This function goes through each spot in the key and using cosine similarity with every single possible shift of the english alphabet calculates the most likely letter for each spot of the key.
 * It returns a list of up to 4 guesses for the key.
 */
func guessKey(keyLength int, inputText string) [4]string {
	outputOptions := make([][]float64, keyLength)
	var output [4]string
	secondLetters := ""
	secondLikelyPos := 0
	secondLikelyValue := 0.0
	thirdLikelyPos := 0
	thirdLikelyValue := 0.0
	for i := 0; i < keyLength; i++ {
		outputOptions[i] = make([]float64, 26)
	}
	for keySpot := 0; keySpot < keyLength; keySpot++ {
		sameAlph := makeString(keySpot, keyLength, inputText)
		frequency := getStringFrequencies(sameAlph)
		for shift := 0; shift < 26; shift++ {
			outputOptions[keySpot][shift] = similarity(shift, frequency)
		}
		max := findMax(outputOptions[keySpot])
		maxValue := outputOptions[keySpot][max]
		outputOptions[keySpot][max] = 0
		max2 := findMax(outputOptions[keySpot])
		output[0] += string(max + 65)
		if (maxValue - outputOptions[keySpot][max2] < 0.001) {
			secondLetters += string(max2 + 65)
			if (secondLikelyValue < outputOptions[keySpot][max2]) {
				thirdLikelyPos = secondLikelyPos
				thirdLikelyValue = secondLikelyValue
				secondLikelyPos = len(secondLetters)
				secondLikelyValue = outputOptions[keySpot][max2]
			} else if (thirdLikelyValue < outputOptions[keySpot][max2]) {
				thirdLikelyPos = len(secondLetters)
				thirdLikelyValue = outputOptions[keySpot][max2]
			}
		} else {
			secondLetters += string(max + 65)
		}
	}

		hold := strings.Split(output[0], "")
		hold2 := strings.Split(secondLetters, "")
	if (thirdLikelyPos != 0) {
		hold[secondLikelyPos] = hold2[secondLikelyPos]
		output[1] = strings.Join(hold, "")
		hold = strings.Split(output[0], "")
		hold[thirdLikelyPos] = hold2[thirdLikelyPos]
		output[2] = strings.Join(hold, "")
		hold[secondLikelyPos] = hold2[secondLikelyPos]
		output[3] = strings.Join(hold, "")
	} else if (secondLikelyPos != 0) {
		hold[secondLikelyPos] = hold2[secondLikelyPos]
		output[1] = strings.Join(hold, "")
		output[2] = ""
		output[3] = ""
	} else {
		output[1] = ""
		output[2] = ""
		output[3] = ""
	}

	return output
}

func main() {
	args := os.Args[1:]
	if (len(args) != 1) {
		fmt.Println("The name of a file with encrypted text in it needs to be provided as the only argument")
		os.Exit(1)
	}
	fileName := args[0]
	fileInput := getInput(fileName)
	keyLength := indexOfCoincidence(fileInput)
	fmt.Println(keyLength)
    keyLength = 11
	keyGuess := guessKey(keyLength, fileInput)
	for i := 0; i < 4; i++ {
		if (keyGuess[i] != "") {
			fmt.Println(keyGuess[i])
		}
	}
}
