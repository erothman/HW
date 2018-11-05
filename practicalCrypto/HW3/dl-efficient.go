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

func x_con(x, alph, beta, p *big.Int) *big.Int {
    if big.NewInt(1).Cmp(big.NewInt(0).Mod(x, big.NewInt(3))) == 0 {
        return big.NewInt(0).Mod(big.NewInt(0).Mul(beta, x), p)
    } else if big.NewInt(0).Cmp(big.NewInt(0).Mod(x, big.NewInt(3))) == 0 {
        return big.NewInt(0).Mod(big.NewInt(0).Mul(x,x), p)
    } else {
        return big.NewInt(0).Mod(big.NewInt(0).Mul(alph, x), p)
    }
}

func a_con(a, x, p *big.Int) *big.Int {
    if big.NewInt(1).Cmp(big.NewInt(0).Mod(x, big.NewInt(3))) == 0 {
        return a
    } else if big.NewInt(0).Cmp(big.NewInt(0).Mod(x, big.NewInt(3)))== 0 {
        return big.NewInt(0).Mod(big.NewInt(0).Mul(a, big.NewInt(2)), big.NewInt(0).Sub(p, big.NewInt(1)))
    } else {
        return big.NewInt(0).Mod(big.NewInt(0).Add(a, big.NewInt(1)), big.NewInt(0).Sub(p, big.NewInt(1)))
    }
}

func b_con(b, x, p *big.Int) *big.Int {
    if big.NewInt(1).Cmp(big.NewInt(0).Mod(x, big.NewInt(3))) == 0 {
        return big.NewInt(0).Mod(big.NewInt(0).Add(b, big.NewInt(1)), big.NewInt(0).Sub(p, big.NewInt(1)))
    } else if big.NewInt(0).Cmp(big.NewInt(0).Mod(x, big.NewInt(3)))== 0 {
        return big.NewInt(0).Mod(big.NewInt(0).Mul(b, big.NewInt(2)), big.NewInt(0).Sub(p, big.NewInt(1)))
    } else {
        return b
    }
}

func pollard_rho(p, g, h *big.Int) *big.Int {
    fmt.Println(p)
    fmt.Println(g)
    fmt.Println(h)
    var m1 map[string]*big.Int
    var m2 map[string]*big.Int
    var m3 map[string]*big.Int
    m1 = make(map[string]*big.Int)
    m2 = make(map[string]*big.Int)
    m3 = make(map[string]*big.Int)
    fmt.Println(m1)
    m1[big.NewInt(0).String()] = big.NewInt(1)
    m2[big.NewInt(0).String()] = big.NewInt(0)
    m3[big.NewInt(0).String()] = big.NewInt(0)
    fmt.Println(m1)
    fmt.Println(m1[big.NewInt(0).String()])
    i := big.NewInt(1)
    i2 := big.NewInt(2)
    m1[big.NewInt(1).String()] = x_con(m1[big.NewInt(0).String()], g, h, p)
    m2[big.NewInt(1).String()] = a_con(m2[big.NewInt(0).String()], m1[big.NewInt(0).String()], p)
    m3[big.NewInt(1).String()] = b_con(m3[big.NewInt(0).String()], m1[big.NewInt(0).String()], p)
    m1[big.NewInt(2).String()] = x_con(m1[big.NewInt(1).String()], g, h, p)
    m2[big.NewInt(2).String()] = a_con(m2[big.NewInt(1).String()], m1[big.NewInt(1).String()], p)
    m3[big.NewInt(2).String()] = b_con(m3[big.NewInt(1).String()], m1[big.NewInt(1).String()], p)
    fmt.Println(m1)
    for m1[i.String()].Cmp(m1[i2.String()]) != 0 {
        m1[big.NewInt(0).Add(i2, big.NewInt(1)).String()] = x_con(m1[i2.String()], g, h, p)
        m2[big.NewInt(0).Add(i2, big.NewInt(1)).String()] = a_con(m2[i2.String()], m1[i2.String()], p)
        m3[big.NewInt(0).Add(i2, big.NewInt(1)).String()] = b_con(m3[i2.String()], m1[i2.String()], p)
        i2 = big.NewInt(0).Add(i2, big.NewInt(1))
        m1[big.NewInt(0).Add(i2, big.NewInt(1)).String()] = x_con(m1[i2.String()], g, h, p)
        m2[big.NewInt(0).Add(i2, big.NewInt(1)).String()] = a_con(m2[i2.String()], m1[i2.String()], p)
        m3[big.NewInt(0).Add(i2, big.NewInt(1)).String()] = b_con(m3[i2.String()], m1[i2.String()], p)
        i2 = big.NewInt(0).Add(i2, big.NewInt(1))
        i = big.NewInt(0).Add(i, big.NewInt(1))
    }
    r := big.NewInt(0).Mod(big.NewInt(0).Sub(m3[i.String()], m3[i2.String()]), big.NewInt(0).Sub(p, big.NewInt(1)))
    fmt.Println(big.NewInt(0).Sub(p, big.NewInt(1)))
    x := big.NewInt(0).Mod(big.NewInt(0).Div(big.NewInt(0).Sub(m2[i2.String()], m2[i.String()]), r), big.NewInt(0).Sub(p, big.NewInt(1)))
    fmt.Println(m1["51"])
    fmt.Println(m2["51"])
    fmt.Println(m3["51"])
    fmt.Println(m1["102"])
    fmt.Println(m2["102"])
    fmt.Println(m3["102"])
    fmt.Println()
    fmt.Println(i)
    fmt.Println(m2[i.String()])
    fmt.Println(m2[i2.String()])
    fmt.Println(m3[i.String()])
    fmt.Println(m3[i2.String()])
    fmt.Println(big.NewInt(0).Sub(m2[i2.String()], m2[i.String()]))
    fmt.Println(r)
    fmt.Println(x)
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
