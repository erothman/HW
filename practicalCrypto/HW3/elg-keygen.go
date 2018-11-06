/*a
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

func big_exp(c, d *big.Int) *big.Int{

    result := big.NewInt(1)
    cprime := c

    for i := 0; i <= d.BitLen(); i++ {
        if d.Bit(i) == 1 {
            result = big.NewInt(0).Mul(result, cprime)
        }

        cprime = big.NewInt(0).Mul(cprime, cprime)
    }

    return result
}

func factor_power_2(n *big.Int) (int, *big.Int) {
    r := 0
    d := n
    two := big.NewInt(2)
    for new(big.Int).Mod(d, two).Cmp(big.NewInt(0)) == 0 {
        d = new(big.Int).Div(d, two)
        r++
    }
    return r, d
}

func Miller_Rabin_check_prime(n *big.Int, k int) bool {
    r,d := factor_power_2(new(big.Int).Sub(n, big.NewInt(1)))
    two := big.NewInt(2)
    for i:=0; i < k; i++ {
        flag := false
        a, err := rand.Int(rand.Reader, new(big.Int).Sub(n, two))
        if err != nil {
            fmt.Println(err)
        }
        for a.Cmp(two) == -1 {
            a, err = rand.Int(rand.Reader, new(big.Int).Sub(n, two))
            if err != nil {
                fmt.Println(err)
            }
        }
        x := modular_exponentiation(a, d, n)
        if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
            continue
        }
        for j:=0; j < r-1; j++ {
            x = modular_exponentiation(x, two, n)
            if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
                flag = true
                j = r-1
            }
        }
        if !flag {
            return false
        }
    }
    return true
}

func generate_primes() (*big.Int, *big.Int) {
    one := big.NewInt(1)
    upper_bound := big.NewInt(0).Sub(big_exp(big.NewInt(2), big.NewInt(20)), one)
    q, err := rand.Int(rand.Reader, upper_bound)
    if err != nil {
        fmt.Println(err)
    }
    for q.Cmp(big.NewInt(2048)) < 0 || !Miller_Rabin_check_prime(q, 4) {
        q2, err := rand.Int(rand.Reader, upper_bound)
        q = q2
        if err != nil {
            fmt.Println(err)
        }
    }
    i := int64(2)
    p := new(big.Int).Add(big.NewInt(0).Mul(q, big.NewInt(i)), one)
    for !Miller_Rabin_check_prime(p, 4) && p.Cmp(big_exp(big.NewInt(2), big.NewInt(20))) == -1 {
        i++
        p2 := new(big.Int).Add(new(big.Int).Mul(q, big.NewInt(i)), one)
        p = p2
    }

    if p.Cmp(big_exp(big.NewInt(2), big.NewInt(20))) != -1 {
        return generate_primes()
    }
    return p, q
}

func prime_factor(a *big.Int) map[string]*big.Int {
    m := make(map[string]*big.Int)

    if big.NewInt(0).Mod(a, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
        m[big.NewInt(2).String()] = big.NewInt(2)
        for big.NewInt(0).Mod(a, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
            a = big.NewInt(0).Div(a, big.NewInt(2))
        }
    }
    i := big.NewInt(3)
    for a.Cmp(big.NewInt(1)) == 1 {
        if big.NewInt(0).Mod(a, i).Cmp(big.NewInt(0)) == 0 {
            m[i.String()] = i
            a = big.NewInt(0).Div(a, i)
        } else {
            i = big.NewInt(0).Add(i, big.NewInt(1))
        }
    }
    return m
}

func test_gen(g, p, group_size *big.Int, factors map[string]*big.Int) bool{

    for _, value := range factors {
        gen := modular_exponentiation(g, big.NewInt(0).Div(group_size, value), p)
        if gen.Cmp(big.NewInt(1)) == 0 {
            return false
        }
    }
    return true
}

func get_generator(p *big.Int, q *big.Int) *big.Int {
    group_size := new(big.Int).Sub(p, big.NewInt(1))
    j := new(big.Int).Div(group_size, q)
    fmt.Println(j)
    factors := prime_factor(j)
    factors[q.String()] = q
    i := big.NewInt(2)
    for group_size.Cmp(i) > 0 {
        if test_gen(i, p, group_size, factors) {
            return i
        }
        i = new(big.Int).Add(i, big.NewInt(1))
    }
    fmt.Println("ERROR")
    return group_size
}

func gen_a(p *big.Int) *big.Int {
    a, err := rand.Int(rand.Reader, big.NewInt(0).Sub(p, big.NewInt(1)))
    if err != nil {
        fmt.Println(err)
    }
    return a
}

func write_Bob_output(filename string, p, g, ga *big.Int) {

    pstring := p.String()
    gstring := g.String()
    gastring := ga.String()

    var strBuilder strings.Builder
    strBuilder.WriteString("(")
    strBuilder.WriteString(pstring)
    strBuilder.WriteString(",")
    strBuilder.WriteString(gstring)
    strBuilder.WriteString(",")
    strBuilder.WriteString(gastring)
    strBuilder.WriteString("}")

//    output := "(" + pstring + "," + gstring + "," + gastring + ")"

    err := ioutil.WriteFile(filename, []byte(strBuilder.String()), 0644)
    if err != nil {
        fmt.Println(err)
    }
}

func write_secret_output(filename string, p, g, a *big.Int) {
    pstring := p.String()
    gstring := g.String()
    astring := a.String()

    var strBuilder strings.Builder
    strBuilder.WriteString("(")
    strBuilder.WriteString(pstring)
    strBuilder.WriteString(",")
    strBuilder.WriteString(gstring)
    strBuilder.WriteString(",")
    strBuilder.WriteString(astring)
    strBuilder.WriteString(")")

//    output := "(" + pstring + "," + gstring + "," + gastring + ")"

    err := ioutil.WriteFile(filename, []byte(strBuilder.String()), 0644)
    if err != nil {
        fmt.Println(err)
    }

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

    bobFileName := args[0]
    secretFileName := args[1]

    p, q := generate_primes()
    g := get_generator(p, q)
    a := gen_a(p)
    ga := modular_exponentiation(g, a, p)
    write_Bob_output(bobFileName, p, g, ga)
    write_secret_output(secretFileName, p, g, a)

}
