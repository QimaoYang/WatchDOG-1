package management

import (
	"crypto/cipher"
	"crypto/des"

	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
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

// check QR code whether is not expired
func checkQRcode(decrypt_time string) string {
	var errMsg string
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", decrypt_time, loc)
	timestamp := tmp.Unix()
	log.Println("the decrypt time is: ", timestamp)

	now := time.Now().Unix()
	dt := now - timestamp
	duration, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(dt)/float64(3600)), 64)
	log.Println("the duration is: ", duration)

	if duration > 1 {
		errMsg := "Sorry, the QR code has expired!"
		log.Println(errMsg)
	}
	return errMsg
}

func DecryptCode(encryp_text string) (string, bool) {
	var seat_number string
	expired := false
	key := getKey()

	decodedStr, _ := hex.DecodeString(encryp_text)
	log.Println("start decrypting the seat_number and date...")
	decrypt_text := decryptDES(decodedStr, key)
	log.Println("the decrypt_text is: ", decrypt_text)

	decrypt_time := decrypt_text[5:]
	errMsg := checkQRcode(decrypt_time)
	if errMsg != "" {
		expired = true
	}
	seat_number = decrypt_text[:5]
	log.Println("the seat_number is: ", seat_number)
	return seat_number, expired
}
