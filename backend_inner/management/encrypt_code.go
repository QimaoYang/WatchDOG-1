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

func EncryptCode(w http.ResponseWriter, r *http.Request) {
	var ipAddr string
	var encryp_str string

	if !IpFilter(r) {
		ipAddr = "true"
		r.ParseForm()
		exPath, _ := os.Getwd()
		fmt.Println("expath: ", exPath)
		file, err := os.Open(exPath + "/management/key.txt")
		// file, err := os.Open(exPath + "/" + "key.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		key, err := ioutil.ReadAll(file)

		// seat_number := "02128"
		seat_number := r.PostFormValue("seat_number")
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		cipher_text := []byte(seat_number + currentTime)
		fmt.Println(string(cipher_text))
		encryp_text := encryptDES(cipher_text, key)
		fmt.Println(encryp_text)

		encryp_str = hex.EncodeToString(encryp_text)
		fmt.Println("byte convert str:", encryp_str)
	} else {
		ipAddr = "false"
	}

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ipAddr":      ipAddr,
		"encryp_text": encryp_str,
	})
	w.Write(resp)
}
