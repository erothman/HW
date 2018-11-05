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
        "crypto/rand"
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

func gen_b(p *big.Int) *big.Int {
    a, err := rand.Int(rand.Reader, big.NewInt(0).Sub(p, big.NewInt(1)))
    if err != nil {
        fmt.Println(err)
    }
    return a
}

func write_Alice_output(filename string, gb *big.Int) {

    gbstring := gb.String()

    var strBuilder strings.Builder
    strBuilder.WriteString("(")
    strBuilder.WriteString(gbstring)
    strBuilder.WriteString(")")

    err := ioutil.WriteFile(filename, []byte(strBuilder.String()), 0644)
    if err != nil {
        fmt.Println(err)
    }
}

func get_inputs(filename string) (*big.Int, *big.Int, *big.Int) {
    input, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    }
    stringInput := string(input)

    stringSplit := strings.Split(stringInput, ",")

    p, ok1 := big.NewInt(0).SetString(stringSplit[0][1:len(stringSplit[0])],10)
    g, ok2 := big.NewInt(0).SetString(stringSplit[1][0:len(stringSplit[1])],10)
    h, ok3 := big.NewInt(0).SetString(stringSplit[2][0:len(stringSplit[2])-1],10)
    if !ok1 || !ok2 || !ok3 {
        fmt.Println("ERROR")
    }

    return p, g, h
}

func brute_force(p, g, h *big.Int) *big.Int {
    x := big.NewInt(0)
    for x.Cmp(p) == -1 {
        fmt.Println(x)
        if h.Cmp(modular_exponentiation(g, x, p)) == 0 {
            return x
        }
        x = big.NewInt(0).Add(x, big.NewInt(1))
    }
    return big.NewInt(0)
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {

    args := os.Args[1:]
    if (len(args) != 1) {
        fmt.Println("Wrong Arguments: There need to be two command line arguments.")
        fmt.Println("The first one is the key for the encryption that must consist of all uppercase letters.")
        fmt.Println("The second argument is the file name for encryption or decrpytion.")
        os.Exit(1)
    }

    inputFileName := args[0]

    p, g, h := get_inputs(inputFileName)

    x := brute_force(p,g,h)
    fmt.Println(x)
}
