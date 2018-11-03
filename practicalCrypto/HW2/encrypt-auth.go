/*
 * Written by Eric Rothman
 * This program recieves a variable to decide the mode, a hexadecimal 32 byte key and the name of a text file as input and output.
 * It then either encrypts or decrypts a AES block cipher using the first 16 bytes of they key as the MAC key and the second as the encrypt key.
 * It writes the output to the output file.
 */


package main

import (
		"fmt"
		"os"
		"io/ioutil"
        "crypto/sha256"
        "crypto/aes"
        "crypto/rand"
        "encoding/hex"
	   )

func xor_bytes(byteString1, byteString2 [16]byte) []byte {
    output := make([]byte, len(byteString1))
    for i := 0; i < len(byteString1); i++ {
        output[i] = byteString1[i] ^ byteString2[i]
    }
    return output
}

func HMAC_SHA_256(key [16]byte, message string) [32]byte{

    var ipad =  [16]byte{}
    var opad = [16]byte{}
    for i := 0; i < 16; i++ {
        ipad[i] = 0x36
        opad[i] = 0x5C
    }

    ipadXor := xor_bytes(ipad, key)
    opadXor := xor_bytes(opad, key)
    firstHash := ""
    secondHash := ""
    firstHash = hex.EncodeToString(ipadXor) + message
    secondHash = hex.EncodeToString(opadXor)

    hashed1 := sha256.Sum256([]byte(firstHash))
    secondHash = secondHash + hex.EncodeToString(hashed1[:])

    HMAC_SHA256_Value := sha256.Sum256([]byte(secondHash))

    return HMAC_SHA256_Value
}

func AES_CBC_ENC(key [16]byte, IV [16]byte, message string) []byte {
    messageBytes := []byte(message)
    initiation := IV
    aesCipher, err := aes.NewCipher(key[:])
    if (err != nil) {
        fmt.Println(err)
    }
    ciphertext := make([]byte, len(messageBytes))
    for i := 0; i < len(messageBytes); i += 16 {
      cipherHold := make([]byte, 16)
      bytes := [16]byte{}
      for j := 0; j < 16; j++ {
          bytes[j] = messageBytes[i + j]
      }
      xored := xor_bytes(initiation, bytes)
      aesCipher.Encrypt(cipherHold, xored)
      for k := 0; k < 16; k++ {
          initiation[k] = cipherHold[k]
          ciphertext[i + k] = cipherHold[k]
      }
    }
    return ciphertext
}

func AES_CBC_DEC(key [16]byte, IV [16]byte, ciphertext []byte) []byte {
    vector := IV
    aesCipher, err := aes.NewCipher(key[:])
    if (err != nil) {
        fmt.Println(err)
    }
    plaintext := make([]byte, len(ciphertext))
    for i := 0; i < len(ciphertext); i += 16 {
        vectHold := make([]byte, 16)
        bytes := make([]byte, 16)
        for j := 0; j < 16; j++ {
            bytes[j] = ciphertext[i + j]
        }
        aesCipher.Decrypt(vectHold, bytes)
        vectHold2 := [16]byte{}
        for index, element := range vectHold {
            vectHold2[index] = element
        }
        plaintextHold := xor_bytes(vectHold2, vector)
        for k := 0; k < 16; k++ {
            vector[k] = bytes[k]
            plaintext[i + k] = plaintextHold[k]
        }
    }
    return plaintext
}

func encryption(encryptKey, MACKey [16]byte, plaintext string) string {
    MAC := HMAC_SHA_256(MACKey, plaintext)
    MAC2 := make([]byte, 32)
    for index, value := range MAC {
        MAC2[index] = value
    }
    m1 := plaintext + string(MAC2)
    n := len(m1) % 16
    ps := make([]byte, 16-n)
    for i := 0; i < 16-n; i++ {
        ps[i] = byte(16-n)
    }
    m2 := m1 + string(ps)
    randomVect := make([]byte, 16)
    _, err := rand.Read(randomVect)
    if (err != nil) {
        panic(err)
    }
    IV := [16]byte{}
    for j := 0; j < 16; j++{
        IV[j] = randomVect[j]
    }
    EncryptedMessage := AES_CBC_ENC(encryptKey, IV, m2)
    return string(randomVect) + string(EncryptedMessage)
}

func getFileInput(fileName string) string {
    input, err := ioutil.ReadFile(fileName)
    if (err != nil) {
        panic(err)
    }
    return string(input)
}

func printOutput(fileName, output string) {
    err := ioutil.WriteFile(fileName, []byte(output), 0644)
    if (err != nil) {
        panic(err)
    }
}


func decryption(decryptKey, MACKey [16]byte, ciphertext string) string {
    ciphertextBytes := []byte(ciphertext)
    IV := [16]byte{}
    cipherMessage := make([]byte, (len(ciphertextBytes) - 16))
    for i := 0; i < 16; i++ {
        IV[i] = ciphertextBytes[i]
    }
    for j := 16; j < len(ciphertextBytes); j++ {
        cipherMessage[j - 16] = ciphertextBytes[j]
    }

    m2 := AES_CBC_DEC(decryptKey, IV, cipherMessage)

    length_of_padding := int(m2[len(m2)-1])
    if (length_of_padding >= 17 || length_of_padding <= 0) {
        fmt.Println("INVALID PADDING")
        os.Exit(0)
    }
    m1 := make([]byte, (len(m2) - length_of_padding))
    for k := 0; k < length_of_padding; k++ {
        if (m2[len(m2) - 1 - k] != byte(length_of_padding)) {
            fmt.Println("INVALID PADDING")
            os.Exit(0)
        }
    }
    for x := 0; x < (len(m2) - length_of_padding); x++ {
        m1[x] = m2[x]
    }
    if (len(m1) <= 32) {
        fmt.Println("INVALID MAC")
        os.Exit(0)
    }
    T := make([]byte, 32)
    M := make([]byte, len(m1)-32)
    for y := 0; y < 32; y++ {
        T[31 - y] = m1[len(m1) - y - 1]
    }
    for z := 0; z < len(m1) - 32; z++ {
        M[z] = m1[z]
    }
    T2 := HMAC_SHA_256(MACKey, string(M))
    for index, value := range T2 {
        if (T[index] != value) {
            fmt.Println("INVALID MAC")
            os.Exit(0)
        }
    }

    return string(M)
}

/*
 *This is the control function. Calls the other functions.
 */
func main() {

	args := os.Args[1:]
	if (len(args) != 7) {
		fmt.Println("HERE")
	}

    mode := args[0]
    keyBytes, err := hex.DecodeString(args[2])

    if err != nil {
        panic(err)
    }
    input_file_name := args[4]
    output_file_name := args[6]

    if (args[1] != "-k") {
        fmt.Println("The key needs a -k before being entered.")
        os.Exit(1)
    }
    if (args[3] != "-i") {
        fmt.Println("The input file needs a -i before being entered.")
        os.Exit(1)
    }
    if (args[5] != "-o") {
        fmt.Println("The output file needs a -o before being entered.")
        os.Exit(1)
    }
    if (len(keyBytes) != 32) {
        fmt.Println(len(keyBytes))
        fmt.Println(keyBytes)
        fmt.Println("The key needs to be 32 bytes, 16 for the encryption key and 16 for the MAC key.")
        os.Exit(1)
    }

    input := getFileInput(input_file_name)

    cipherKey := [16]byte{}
    MACKey := [16]byte{}
    for i := 0; i < 16; i++ {
        cipherKey[i] = keyBytes[i]
        MACKey[i] = keyBytes[16 + i]
    }
    output := ""
    if (mode == "encrypt") {
        output = encryption(cipherKey, MACKey, input)
     } else {
        output = decryption(cipherKey, MACKey, input)
    }

    printOutput(output_file_name, output)
}
