package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

const (
	keyword = "84a9c82328738185"
	ivword  = "ea33fafa60eed64e"
)

func MainEncrypt(endata string) string {

	key := []byte(keyword)
	result, err := AesEncrypt([]byte(endata), key)
	if err != nil {
		beego.Error(err)
	}

	var data = Base64encodeUrl(base64.StdEncoding.EncodeToString(result))
	return data
}

func MainDecrypt(dedata string) ([]byte, error) {
	// beego.Debug(dedata)
	key := []byte(keyword)
	result, err := base64.StdEncoding.DecodeString(Base64decodeUrl(dedata))
	origData, err := AesDecrypt(result, key)

	return origData, err
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	iv := []byte(ivword)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) (dec []byte, e error) {
	origData := make([]byte, len(crypted))

	defer func() {
		if err := recover(); err != nil {
			e = errors.New(fmt.Sprintln(err))
			dec = nil
		}
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// blockSize := block.BlockSize()
	iv := []byte(ivword)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/// <summary>
/// 从普通字符串转换为适用于URL的Base64编码字符串
/// </summary>
func Base64decodeUrl(base64String string) string {
	str := strings.Replace(base64String, "-", "+", -1)
	return str
}

/// <summary>
/// 从普通字符串转换为适用于URL的Base64编码字符串
/// </summary>
func Base64encodeUrl(base64String string) string {
	str := strings.Replace(base64String, "+", "-", -1)
	return str
}
