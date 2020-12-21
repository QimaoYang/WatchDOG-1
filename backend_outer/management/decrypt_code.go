package management

import (
	"crypto/cipher"
	"crypto/des"
	"net/http"

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

func DecrptCode(w http.ResponseWriter, r *http.Request) {
	exPath, _ := os.Getwd()
	// fmt.Println("expath", exPath)
	file, err := os.Open(exPath + "\\" + "management\\key.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	key, err := ioutil.ReadAll(file)

	encryp_text := []byte{205, 147, 30, 249, 67, 158, 150, 106, 221, 176, 252, 231, 75, 40, 177, 235, 159, 40, 106, 85, 36, 237, 205, 154, 38, 251, 151, 0, 110, 36, 60, 215}
	decrypt_text := decryptDES(encryp_text, key)
	fmt.Println(decrypt_text)
	seat_number := decrypt_text[:10]
	time := decrypt_text[10:]
	fmt.Println(seat_number)
	fmt.Println(time)
}
