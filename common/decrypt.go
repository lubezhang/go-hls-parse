package common

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesDecrypt(cryptoData []byte, key string, ivs ...string) ([]byte, error) {
	bKey := []byte(key)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	var iv []byte
	if len(ivs) == 0 {
		iv = bKey
	} else {
		iv = []byte(ivs[0])
	}
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(cryptoData))
	blockMode.CryptBlocks(origData, cryptoData)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
