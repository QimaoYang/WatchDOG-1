package management

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"net/http"

	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func padding(src []byte, blocksize int) []byte {
	n := len(src)
	padnum := blocksize - n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	dst := append(src, pad...)
	return dst
}

func encryptDES(src []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func EncrptCode(w http.ResponseWriter, r *http.Request) {
	exPath, _ := os.Getwd()
	// fmt.Println("expath", exPath)
	file, err := os.Open(exPath + "\\" + "management\\key.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	key, err := ioutil.ReadAll(file)

	seat_number := "WS02.02128"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	cipher_text := []byte(seat_number + currentTime)
	fmt.Println(string(cipher_text))

	encryp_text := encryptDES(cipher_text, key)
	fmt.Println(encryp_text)
	// encodedStr := hex.EncodeToString(encryp_text)
	// fmt.Println("byte convert str:", encodedStr)
}
