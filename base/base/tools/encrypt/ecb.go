package encrypt

import (
	"crypto/cipher"
)

//ECBEncrypt
func ECBEncrypt(block cipher.Block, src []byte, padding PaddingMode) ([]byte, error) {
	blockSize := block.BlockSize()
	src = Padding(padding, src, blockSize)
	dst := make([]byte, len(src))
	tmpData := make([]byte, blockSize)
	for index := 0; index < len(src); index += blockSize {
		block.Encrypt(tmpData, src[index:index+blockSize])
		copy(dst[index:], tmpData)
	}
	return dst, nil
}

//ECBDecrypt
func ECBDecrypt(block cipher.Block, src []byte, padding PaddingMode) ([]byte, error) {
	dst := make([]byte, len(src))
	blockSize := block.BlockSize()
	tmpData := make([]byte, blockSize)
	for index := 0; index < len(src); index += blockSize {
		block.Decrypt(tmpData, src[index:index+blockSize])
		copy(dst[index:], tmpData)
	}
	dst = UnPadding(padding, dst)
	return dst, nil
}