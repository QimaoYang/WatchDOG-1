package management

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"net/http"
	"time"

	"encoding/hex"
	"encoding/json"
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
	file, err := os.Open(exPath + "/management/key.txt")
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
	w.Header().Add("Content-Type", "application/json")
	var encrypt_str string
	var msg string
	ipFilter, ipAddr := IpFilter(r)

	if ipFilter {
		msg = "true"

		key := getKey()

		// get seat_number
		body := json.NewDecoder(r.Body)
		var params map[string]string
		body.Decode(&params)
		seat_number := params["seat_number"]
		log.Println("get seat_number from frontend: ", seat_number)

		if seat_number != "" {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			cipher_text := []byte(seat_number + currentTime)

			log.Println("start encrypting the seat_number and date...")
			encrypt_text := encryptDES(cipher_text, key)
			encrypt_str = hex.EncodeToString(encrypt_text)
			log.Println("the encrypt_text is: ", encrypt_str)
		} else {
			errMsg := "cannot get seat_number from frontend"
			log.Println(errMsg)
			resp, _ := json.Marshal(map[string]string{
				"errMsg": errMsg,
			})
			w.Write(resp)
			return
		}
	} else {
		msg = "false"
	}

	resp, _ := json.Marshal(map[string]string{
		"encrypt_text": encrypt_str,
		"ipAddr":       ipAddr,
		"msg":          msg,
	})
	w.Write(resp)
}
