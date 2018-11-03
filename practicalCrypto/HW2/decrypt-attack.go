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
		"io/ioutil"
        "os/exec"
        "strings"
        "flag"
	   )

func xorBytes(byteString1, byteString2 []byte) []byte {
    output := make([]byte, len(byteString1))
    for i := 0; i < len(byteString1); i++ {
        output[i] = byteString1[i] ^ byteString2[i]
    }
    return output
}

func getFileInput(fileName string) []byte {
    input, err := ioutil.ReadFile(fileName)
    if (err != nil) {
        panic(err)
    }
    return input
}

func printOutput(fileName, output string) {
    err := ioutil.WriteFile(fileName, []byte(output), 0644)
    if (err != nil) {
        panic(err)
    }
}


func testValidity(testInput []byte) string {

    input_file_name := "test.txt"
    printOutput(input_file_name, string(testInput))

    execCmd := exec.Command("./decrypt-test", "-i", input_file_name)
    execOut, err := execCmd.CombinedOutput()
    if err != nil {
        panic(execOut)
    }
    os.Remove(input_file_name)
    out := strings.TrimSpace(string(execOut))
    return out
}

func decryptLastBlock(ciphertext []byte) [16]byte {
    decryptedOutput := [16]byte{}
    current_message_length := len(ciphertext)
    for i := 0; i < 16; i++ {
       j := 1
       flag := false
       for j < 256 && !flag {
           delta := make([]byte, 16)
           p_star := make([]byte, 16)
           for a := 0; a < 16; a++ {
               delta[a] = 0
               p_star[a] = ciphertext[current_message_length - 32 + a]
           }
           for b := 0; b < i; b++ {
               delta[15 - b] = decryptedOutput[15 - b] ^ byte(i+1)
           }
           delta[15 - i] = byte(j)
           fmt.Println(delta)
           message_copy := make([]byte, current_message_length)
           copy(message_copy, ciphertext)
           changed_p_star := xorBytes(delta, p_star)
           for k := 0; k < 16; k++ {
               message_copy[current_message_length - 32 + k] = changed_p_star[k]
           }
           test_value := testValidity(message_copy)
           if (test_value != "INVALID PADDING") {
               if (i == 0) {
                   delta[14] = byte(1)
                   message_copy[current_message_length - 18] = xorBytes(delta, p_star)[14]
                   test_value2 := testValidity(message_copy)
                   if (test_value2 != "INVALID PADDING") {
                        flag = true
                        decryptedOutput[15-i] = (byte(1) ^ byte(j))
                   }
               } else {
                   flag = true
                   decryptedOutput[15-i] = (byte(i+1) ^ byte(j))
               }
           }
           j++
       }
    }
    return decryptedOutput
}

func attackMessage(ciphertext []byte) string {
    current_message_length := len(ciphertext)
    full_output := make([]byte, current_message_length - 16 - 32)

    flag := false
    i := byte(2)
    padding_length := 0
    for i <= 17 && !flag {
        delta := make([]byte, 16)
        p_star := make([]byte, 16)
        for j := 0; j < 16; j++ {
            delta[j] = 0
            p_star[j] = ciphertext[current_message_length - 32 + j]
        }
        delta[15] = i
        message_copy := make([]byte, current_message_length)
        copy(message_copy, ciphertext)
        changed_p_star := xorBytes(delta, p_star)
        for k := 0; k < 16; k++ {
            message_copy[current_message_length - 32 + k] = changed_p_star[k]
        }
        test_value := testValidity(message_copy)
        if (test_value != "INVALID PADDING") {
            flag = true
            padding_length = int(byte(1) ^ i)
        }
        i++
    }
    if (!flag) {
        padding_length = 1
    }
    current_message_length = current_message_length - 32
    for current_message_length > 16 {
        current_message := make([]byte, current_message_length)
        for b := 0; b < current_message_length; b++ {
            current_message[b] = ciphertext[b]
        }
        decrypted_current_last_block := decryptLastBlock(current_message)
        for index, value := range decrypted_current_last_block {
            full_output[current_message_length - 32 + index] = value
        }
        current_message_length = current_message_length - 16
    }
    output := make([]byte, len(full_output) - padding_length)
    for i := 0; i < len(full_output) - padding_length; i++ {
        output[i] = full_output[i]
    }
    return string(output)
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {


   input_file_name := flag.String("i", "", "file name of input file" )

   flag.Parse()

   input := getFileInput(*input_file_name)
   plaintext := attackMessage([]byte(input))
   fmt.Println([]byte(plaintext))
   fmt.Println(plaintext)
}
