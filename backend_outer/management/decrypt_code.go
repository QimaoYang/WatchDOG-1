package management

import (
	"crypto/cipher"
	"crypto/des"
	"net/http"

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

func DecrptCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	exPath, _ := os.Getwd()
	fmt.Println("expath: ", exPath)
	file, err := os.Open(exPath + "\\" + "management\\key.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	key, err := ioutil.ReadAll(file)

	// encryp_text := "cd931ef9439e966addb0fce74b28b1ebcdd57088f9a85be37b3beaadd845249f"
	encryp_text := r.PostFormValue("encryp_text")
	decodedStr, err := hex.DecodeString(encryp_text)
	fmt.Println(decodedStr)
	decrypt_text := decryptDES(decodedStr, key)
	fmt.Println(decrypt_text)
	seat_number := decrypt_text[:10]
	time := decrypt_text[10:]
	fmt.Println(seat_number)
	fmt.Println(time)
	fmt.Fprintln(w, "seat_number = "+seat_number)
	fmt.Fprintln(w, "time = "+time)
}
