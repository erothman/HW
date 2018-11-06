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
    h, ok3 := big.NewInt(0).SetString(stringSplit[2][0:len(stringSplit[2])-2],10)
    fmt.Println(h)
    if !ok1 || !ok2 || !ok3 {
        fmt.Println("ERROR")
    }

    return p, g, h
}

func extend_euclid_alg(a, n *big.Int) *big.Int {
    t := big.NewInt(0)
    newt := big.NewInt(1)
    r := n
    newr := a
    for newr.Cmp(big.NewInt(0)) != 0 {
        quotient := big.NewInt(0).Div(r, newr)
        rh := newr
        th := newt
        newr = big.NewInt(0).Sub(r, big.NewInt(0).Mul(quotient, newr))
        newt = big.NewInt(0).Sub(t, big.NewInt(0).Mul(quotient, newt))
        r = rh
        t = th
    }
    return t

}

func pollard_rho(p, g, h *big.Int) *big.Int {
    var m map[string]*big.Int

    m = make(map[string]*big.Int)

    i := big.NewInt(0)
    for big.NewInt(0).Sub(p, big.NewInt(1)).Cmp(big.NewInt(0).Mul(i, i)) == 1 {
        alphi := modular_exponentiation(g,i,p)
        fmt.Println(alphi)
        m[alphi.String()] = i
        i = big.NewInt(0).Add(i, big.NewInt(1))
    }
    alphm := modular_exponentiation(g,i,p)
    alphminv := big.NewInt(0).Mod(extend_euclid_alg(alphm, p), p)
    gamma := h

    j := big.NewInt(0)
    for j.Cmp(i) == -1 {
        if m[gamma.String()] != nil {
            return big.NewInt(0).Add(big.NewInt(0).Mul(j, i), m[gamma.String()])
        }
        gamma = big.NewInt(0).Mod(big.NewInt(0).Mul(gamma, alphminv) , p)
        j = big.NewInt(0).Add(j, big.NewInt(1))
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

    x := pollard_rho(p,g,h)
    fmt.Println(x)
}
