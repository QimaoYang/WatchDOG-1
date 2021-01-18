package management

import (
	"crypto/cipher"
	"crypto/des"

	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}

func decryptDES(src []byte, key []byte) string {
	block, _ := des.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return string(src)
}

func getKey() []byte {
	exPath, _ := os.Getwd()
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

func DecryptCode(encryp_text string) string {
	// get the key
	key := getKey()

	decodedStr, _ := hex.DecodeString(encryp_text)
	fmt.Println(decodedStr)
	log.Println("start decrypting the seat_number and date...")
	decrypt_text := decryptDES(decodedStr, key)
	log.Println("the decrypt_text is: " + decrypt_text)

	seat_number := decrypt_text[:5]
	time := decrypt_text[5:]
	log.Println("the seat_number is: " + seat_number)
	log.Println("the time is: " + time)
	return seat_number
}
