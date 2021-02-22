package encrypt

import (
	"bytes"
)

type PaddingMode string
const (
	// PADDING_PKCS5
	PADDING_PKCS5 PaddingMode = "PKCS5"
	//PADDING_PKCS7
	PADDING_PKCS7 PaddingMode = "PKCS7"
	//PADDING_ZEROS
	PADDING_ZEROS PaddingMode = "ZEROS"
)

func Padding(paddingMode PaddingMode, src []byte, blockSize int) []byte {
	switch paddingMode {
	case PADDING_PKCS5:
		src = PKCS5Padding(src, blockSize)
	case PADDING_PKCS7:
		src = PKCS7Padding(src, blockSize)
	case PADDING_ZEROS:
		src = ZerosPadding(src, blockSize)
	}
	return src
}

func UnPadding(paddingMode PaddingMode, src []byte) []byte {
	switch paddingMode {
	case PADDING_PKCS5:
		src = PKCS5Unpadding(src)
	case PADDING_PKCS7:
		src = PKCS7UnPadding(src)
	case PADDING_ZEROS:
		src = ZerosUnPadding(src)
	}
	return src
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	return PKCS7Padding(src, blockSize)
}

func PKCS5Unpadding(src []byte) []byte {
	return PKCS7UnPadding(src)
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func ZerosPadding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

func ZerosUnPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}