package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {
	has := md5.New()
	has.Write([]byte(data))
	tmpstr := has.Sum(nil)
	return hex.EncodeToString(tmpstr)
}

//加盐
func Md5AddSalt(inputpassword, salt string) string {

	return Md5Encode(inputpassword + salt)

}

// 密码校验

func ValidPassword(inputpassword, salt, password string) bool {
	return Md5Encode(inputpassword+salt) == password
}
