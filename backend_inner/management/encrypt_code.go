package management

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"net/http"
	"time"

	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func getKey() []byte {
	exPath, _ := os.Getwd()
	log.Println("expath: ", exPath)
	// file, err := os.Open(exPath + "/management/key.txt")
	file, err := os.Open(exPath + "/" + "key.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	key, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		log.Fatal(err)
	}
	return key
}

func EncryptCode(w http.ResponseWriter, r *http.Request) {
	var encrypt_str string
	var msg string
	ipAddr := GetIP(r)

	if !IpFilter(r) {
		msg = "true"

		// get the key
		key := getKey()

		// get seat_number
		body := json.NewDecoder(r.Body)
		var params map[string]string
		body.Decode(&params)
		seat_number := params["seat_number"]
		log.Println("seat_number: " + seat_number)

		currentTime := time.Now().Format("2006-01-02 15:04:05")
		cipher_text := []byte(seat_number + currentTime)
		fmt.Println(string(cipher_text))

		encrypt_text := encryptDES(cipher_text, key)
		encrypt_str = hex.EncodeToString(encrypt_text)
		log.Println("encrypt_text: ", encrypt_str)
	} else {
		msg = "false"
	}

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"encrypt_text": encrypt_str,
		"ipAddr":       ipAddr,
		"msg":          msg,
	})
	w.Write(resp)
}
