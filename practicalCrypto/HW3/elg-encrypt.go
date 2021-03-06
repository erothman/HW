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
        "io"
		"io/ioutil"
        "crypto/rand"
        "math/big"
        "crypto/sha256"
        "crypto/cipher"
        "crypto/aes"
        "encoding/hex"
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

func write_output(filename string, gb *big.Int, c, IV []byte) {

    gbstring := gb.String()

    var str strings.Builder
    str.WriteString(hex.EncodeToString(IV))
    str.WriteString(hex.EncodeToString(c))

    var strBuilder strings.Builder
    strBuilder.WriteString("(")
    strBuilder.WriteString(gbstring)
    strBuilder.WriteString(",")
    strBuilder.WriteString(str.String())
    strBuilder.WriteString(")")

    err := ioutil.WriteFile(filename, []byte(strBuilder.String()), 0644)
    if err != nil {
        fmt.Println("err")
        os.Exit(0)
    }
}

func get_inputs(filename string) (*big.Int, *big.Int, *big.Int) {
    input, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("error")
        os.Exit(0)
    }
    stringInput := string(input)

    stringSplit := strings.Split(stringInput, ",")

    p, ok1 := big.NewInt(0).SetString(stringSplit[0][1:len(stringSplit[0])],10)
    g, ok2 := big.NewInt(0).SetString(stringSplit[1][0:len(stringSplit[1])],10)
    ga, ok3 := big.NewInt(0).SetString(stringSplit[2][0:len(stringSplit[2])-1],10)
    if !ok1 || !ok2 || !ok3 {
        fmt.Println("Error")
        os.Exit(0)
    }

    return p, g, ga
}

func computeK(ga, gb, gab *big.Int) []byte {


    var strBuilder strings.Builder
    strBuilder.WriteString(ga.String())
    strBuilder.WriteString(gb.String())
    strBuilder.WriteString(gab.String())

    h := sha256.New()
    h.Write([]byte(strBuilder.String()))
    k := h.Sum(nil)
    return k
}

func encryptMessage(messageText string, k []byte) ([]byte, []byte) {
    messageTextBytes := []byte(messageText)

	aesCipher, err := aes.NewCipher(k)
	if err != nil {
		panic(err.Error())
	}

	IV := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, IV); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(aesCipher, 16)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, IV, messageTextBytes, nil)

	return ciphertext, IV
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {

    args := os.Args[1:]
    if (len(args) != 3) {
        fmt.Println("Wrong Arguments: There need to be two command line arguments.")
        fmt.Println("The first one is the key for the encryption that must consist of all uppercase letters.")
        fmt.Println("The second argument is the file name for encryption or decrpytion.")
        os.Exit(1)
    }

    messageText := args[0]
    publicKeyFileName := args[1]
    ciphertextFileName := args[2]

    p, g, ga := get_inputs(publicKeyFileName)

    b := gen_b(p)
    gb := modular_exponentiation(g, b, p)
    gab := modular_exponentiation(ga, b, p)
    k := computeK(ga, gb, gab)
    c, IV := encryptMessage(messageText, k)
    write_output(ciphertextFileName, gb, c, IV)
}
