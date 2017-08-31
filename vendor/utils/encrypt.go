package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/**
*把字节转化为MD5字符串
 */
func MD5(b []byte) string {
	vCrypto := md5.New()
	vCrypto.Write(b)
	return hex.EncodeToString(vCrypto.Sum(nil))
}

/*
*加密用户密码
 */
func EncodeUserPwd(password string) string {
	return MD5([]byte(password))
}
