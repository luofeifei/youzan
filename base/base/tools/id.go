package tools

import "strconv"

//ChaosID 混淆一个int64类型的ID
func ChaosID(id int64, key string) string {
	base := len(key)
	//将seed转换为36进制字符串
	baseInt := strconv.FormatInt(id, base)
	var num string
	for _, char := range baseInt {
		//按照混淆字符串进行映射
		i, _ := strconv.ParseInt(string(char), base, 64)
		num = num + string(key[i])
	}
	return num
}

//RestoreID 还原被混淆过的int64类型ID
func RestoreID(messed string, key string) int64 {
	var baseInt string
	base := len(key)
	//按照混淆字符串进行逆映射
	for _, char := range messed {
		for i, c := range key {
			if char == c {
				s := strconv.FormatInt(int64(i), base)
				baseInt = baseInt + s
			}
		}
	}
	//将36进制字符串转化为10进制
	num, err := strconv.ParseInt(baseInt, base, 64)
	if err != nil {
		return 0
	}
	return num
}
