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
        "crypto/sha256"
        "crypto/aes"
        "crypto/cipher"
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


func get_inputs(filename1, filename2 string) (*big.Int, *big.Int, *big.Int, *big.Int, []byte) {
    input, err := ioutil.ReadFile(filename1)
    if err != nil {
        fmt.Println(err)
    }
    stringInput := string(input)

    stringSplit := strings.Split(stringInput, ",")

    p, ok1 := big.NewInt(0).SetString(stringSplit[0][1:len(stringSplit[0])],10)
    g, ok2 := big.NewInt(0).SetString(stringSplit[1][0:len(stringSplit[1])],10)
    a, ok3 := big.NewInt(0).SetString(stringSplit[2][0:len(stringSplit[2])-1],10)
    if !ok1 || !ok2 || !ok3{
        fmt.Println("error")
        os.Exit(0)
    }

    input2, err := ioutil.ReadFile(filename2)
    if err != nil {
        fmt.Println("error")
        os.Exit(0)
    }

    stringInput2 := string(input2)

    stringSplit2 := strings.Split(stringInput2, ",")

    gb, ok := big.NewInt(0).SetString(stringSplit2[0][1:len(stringSplit2[0])],10)
    c, err2 := hex.DecodeString(stringSplit2[1][0:len(stringSplit2[1])-1])
    if !ok  {
        fmt.Println("error")
        os.Exit(0)
    }
    if err2 != nil {
        fmt.Println("error")
        os.Exit(0)
    }
    return p, a, g, gb, c
}

func decrypt_encryption(k, c []byte) {
    IV := make([]byte, 16)

    for i:=0; i < 16; i++ {
        IV[i] = c[i]
    }
    IVlessC := make([]byte, (len(c)-16))
    for j:=0; j<len(c)-16;j++ {
        IVlessC[j] = c[16+j]
    }
    aescipher, err := aes.NewCipher(k)
	if err != nil {
		fmt.Println("error")
        os.Exit(0)
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(aescipher, 16)
	if err != nil {
		fmt.Println("error")
        os.Exit(0)
	}

	plaintext, err := aesgcm.Open(nil, IV, IVlessC, nil)

	if err != nil {
		fmt.Println("error")
        os.Exit(0)
	}

	fmt.Printf("%s\n", string(plaintext))
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

    p, a, g, gb, c := get_inputs(inputSecretFileName, inputBobFileName)
    gab := modular_exponentiation(gb, a, p)
    ga := modular_exponentiation(g, a, p)
    k := computeK(ga, gb, gab)
    decrypt_encryption(k, c)
}
