package main

import (
	"encoding/base64"
	"crypto/rc4"
	"fmt"
)

func main() {

	param := "mEMz/JFFnA=="
	fmt.Println(string(DescryptRC4Base64(param,"simple_bg")))
}
func IsAllNum(str string) bool{
	bytes := []byte(str)
	for _, item := range bytes {
		if item < '0' || item > '9' {
			return false
		}
	}
	return true
}


func DescryptRC4Base64(p, keystr string) []byte {
	key := []byte(keystr)
	str, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return nil
	}
	data := []byte(str)
	ct, err := rc4.NewCipher(key)
	if err != nil {
		return nil
	}
	dst := make([]byte, len(data))
	ct.XORKeyStream(dst, data)
	return dst
}
func EncryptRC4Base64(p []byte, key string) string {
	k := []byte(key)
	cl, _ := rc4.NewCipher(k)
	dst := make([]byte, len(p))
	cl.XORKeyStream(dst, p)
	str := base64.StdEncoding.EncodeToString(dst)
	return str
}

