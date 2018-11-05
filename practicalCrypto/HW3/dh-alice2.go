/*
 * Written by Eric Rothman
 * This program recieves the name of a text file as input with an AES encyrpted message.
 * It calls decrypt-test to test the output of encrypt-auth on entries.
 * It uses that information to execute a padding oracle attack on the encrypted message, decrypting all of it.
 * It then outputs it to standard out.
 */


package main

import (
		"fmt"
		"os"
        "strings"
		"io/ioutil"
        "math/big"
	   )

func modular_exponentiation(c, d, N *big.Int) *big.Int{

    result := big.NewInt(1)
    cprime := c

    for i := 0; i <= d.BitLen(); i++ {
        if d.Bit(i) == 1 {
            result = big.NewInt(0).Mod(big.NewInt(0).Mul(result, cprime), N)
        }

        cprime = big.NewInt(0).Mod(big.NewInt(0).Mul(cprime, cprime), N)
    }

    return result
}



func get_inputs(filename1, filename2 string) (*big.Int, *big.Int, *big.Int) {
    input, err := ioutil.ReadFile(filename1)
    if err != nil {
        fmt.Println(err)
    }
    stringInput := string(input)

    stringSplit := strings.Split(stringInput, ",")

    p, ok1 := big.NewInt(0).SetString(stringSplit[0][1:len(stringSplit[0])],10)
    a, ok2 := big.NewInt(0).SetString(stringSplit[2][0:len(stringSplit[2])-1],10)
    if !ok1 || !ok2 {
        fmt.Println("ERROR")
    }

    input2, err2 := ioutil.ReadFile(filename2)
    if err != nil {
        fmt.Println(err2)
    }

    stringInput2 := string(input2)

    gb, ok := big.NewInt(0).SetString(stringInput2[1:len(stringInput2)-1],10)
    if !ok  {
        fmt.Println("ERROR")
    }
    return p, a, gb
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {

    args := os.Args[1:]
    if (len(args) != 2) {
        fmt.Println("Wrong Arguments: There need to be two command line arguments.")
        fmt.Println("The first one is the key for the encryption that must consist of all uppercase letters.")
        fmt.Println("The second argument is the file name for encryption or decrpytion.")
        os.Exit(1)
    }

    inputBobFileName := args[0]
    inputSecretFileName := args[1]

    p, a, gb := get_inputs(inputSecretFileName, inputBobFileName)
    fmt.Println(modular_exponentiation(gb, a, p))
}
