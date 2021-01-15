package management

import (
	"crypto/cipher"
	"crypto/des"

	"encoding/hex"
	"fmt"
	"io/ioutil"
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

func DecryptCode(encryp_text string) string {
	exPath, _ := os.Getwd()
	fmt.Println("expath: ", exPath)
	file, err := os.Open(exPath + "/management/key.txt")
	// file, err := os.Open(exPath + "/" + "key.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	key, err := ioutil.ReadAll(file)

	decodedStr, err := hex.DecodeString(encryp_text)
	fmt.Println(decodedStr)
	decrypt_text := decryptDES(decodedStr, key)
	fmt.Println(decrypt_text)
	seat_number := decrypt_text[:5]
	time := decrypt_text[5:]
	fmt.Println(seat_number)
	fmt.Println(time)
	return seat_number
}
