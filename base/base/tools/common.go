package tools

import (
	"encoding/base64"
	"encoding/binary"
	"net"
	"strconv"
)

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) string {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return "0"
	}
	return strconv.FormatUint(uint64(binary.BigEndian.Uint32(ip.To4())), 10)
}

/**
 * @brief base64编码
 * @param string name 待编码的字符
 * @return mixed
 */
func Base64Encode(data string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	return encodeString
}

/**
 * @brief base64解码
 * @param string encodeString base64编码的字符
 * @return mixed
 */
func Base64Decode(encodeString string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err == nil {
		return string(decodeBytes)
	} else {
		return ""
	}
}
