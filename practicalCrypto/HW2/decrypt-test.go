/*
 * Written by Eric Rothman
 * This program recieves a key of capital letters and the name of a text file as input and outputs the text in the text file encrypted using the Vigenere cipher and the key.
 */


package main

import (
		"fmt"
//        "crypto/sha256"
//        "crypto/aes"
//        "crypto/cipher"
//        "crypto/rand"
//        "encoding/hex"
        "flag"
        "os/exec"
//        "bytes"
	   )

var KEY = "0000000000000001000000000000000000000000000000000000000000000000"

func decrypt(input_file_name string) string {
//    binary, lookErr := exec.LookPath("ls")
//    if lookErr != nil {
//        panic(lookErr)
//    }
//    args := []string{"./encrypt-auth", "decrypt", "-k", KEY, "-i", input_file_name, "-o", "out.txt"
    execCmd := exec.Command("./encrypt-auth", "decrypt", "-k", KEY, "-i", input_file_name, "-o", "out2.txt")
    execOut, err := execCmd.CombinedOutput()
    if err != nil {
        fmt.Println(execOut)
        return string(execOut)
//        panic(err)
    }
    if (string(execOut) != "INVALID MAC\n" && string(execOut) != "INVALID PADDING\n") {
        return "SUCCESS"
    }
    return string(execOut)
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {

//	args := os.Args[1:]
//	if (len(args) != 7) {
//		os.Exit(1)
//	}

//    mode := flag.String("t", "", "The mode to run in")

//    fmt.Println(*mode)

   input_file_name := flag.String("i", "", "file name of input file" )

   flag.Parse()

//   fmt.Println(*input_file_name)

   fmt.Println(decrypt(*input_file_name))
}
