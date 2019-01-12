package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"crypto/des"
	"encoding/base64"
	"game2/lib/encrypt/aescypto"
)

func PKCS5Padding(cipher []byte, blockSize int) []byte {
	padding := blockSize - len(cipher)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipher, pad...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func AESEncrypt(src, key string) string {
	byteKey := []byte(key)
	block, _ := aes.NewCipher(byteKey)
	origData := PKCS5Padding([]byte(src), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, byteKey)
	crypt := make([]byte, len(origData))
	blockMode.CryptBlocks(crypt, origData)
	return string(crypt)
}

func AESDecrypt(src, key string) string {
	block, _ := aes.NewCipher([]byte(key))
	blockMode := cipher.NewCBCDecrypter(block, []byte(key))
	origData := make([]byte, len(src))
	blockMode.CryptBlocks(origData, []byte(src))
	origData = PKCS5UnPadding(origData)
	return string(origData)
}

func MD5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func DESEncrypt(src, key string) string {
	block, _ := des.NewCipher([]byte(key))
	origData := PKCS5Padding([]byte(src), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(key))
	crypt := make([]byte, len(origData))
	blockMode.CryptBlocks(crypt, origData)
	return string(crypt)
}

func DESDecrypt(src, key string) string {
	block, _ := des.NewCipher([]byte(key))
	blockMode := cipher.NewCBCDecrypter(block, []byte(key))
	origData := make([]byte, len(src))
	blockMode.CryptBlocks(origData, []byte(src))
	origData = PKCS5UnPadding(origData)
	return string(origData)
}

func DesEcbPkc5Encrypt(data, key string) string {
	block, _ := des.NewCipher([]byte(key))
	bs := block.BlockSize()
	padData := PKCS5Padding([]byte(data), bs)
	if len(padData)%bs != 0 {
		return ""
	}
	out := make([]byte, len(padData))
	dst := out
	for len(padData) > 0 {
		block.Encrypt(dst, padData[:bs])
		padData = padData[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

func AesEcbEncrypt(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err.Error()
	}
	ecb := aescrypto.NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}
