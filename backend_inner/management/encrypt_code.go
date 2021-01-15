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
	var encryp_str string
	var msg string
	ipAddr := GetIP(r)

	if !IpFilter(r) {
		msg = "true"

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

		seat_number := r.PostFormValue("seat_number")
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		cipher_text := []byte(seat_number + currentTime)
		fmt.Println(string(cipher_text))
		encryp_text := encryptDES(cipher_text, key)
		// fmt.Println(encryp_text)

		encryp_str = hex.EncodeToString(encryp_text)
		fmt.Println("encryp_text:", encryp_str)
	} else {
		msg = "false"
	}

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"encryp_text": encryp_str,
		"ipAddr":      ipAddr,
		"msg":         msg,
	})
	w.Write(resp)
}
