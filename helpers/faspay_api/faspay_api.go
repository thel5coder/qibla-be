package faspay_api

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func getSignature(compose string) (res string) {
	var md5 = md5.New()
	var sha1 = sha1.New()
	md5.Write([]byte(compose))
	md5EncryptedStr := md5.Sum(nil)
	md5Str := fmt.Sprintf("%x", md5EncryptedStr)

	sha1.Write([]byte(md5Str))
	sha1EncryptedStr := sha1.Sum(nil)
	res = fmt.Sprintf("%x", sha1EncryptedStr)

	return res
}