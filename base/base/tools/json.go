package tools

import (
	json "github.com/json-iterator/go"
)

/**
 * @brief json编码
 * @param mixed data 需要json编码的数据
 * @return mixed
 */
func JsonEncode(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

/**
 * @brief json解码
 * @param string data 待解码的字符串
 * @return mixed
 */
func JsonDecode(encodeString string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(encodeString), &data)
	if err == nil {
		return data
	} else {
		return nil
	}
}

/**
 * @brief json解码
 * @param byte data 待解码的二进制
 * @param byte v interface 类型
 * @return mixed
 */
func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}
