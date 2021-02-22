package tools

import (
	"fmt"
)

func MustSubstring(str string, start int, end int) string {
	res, err := Substring(str, start, end)
	if err != nil {
		panic(err)
	}
	return res
}

//Substring 截取字符串
func Substring(str string, start int, end int) (string, error) {
	if start < 0 || start >= len(str) {
		return "", fmt.Errorf("start (%d) is out of range", start)
	}
	if end != 0 && end <= start {
		return "", fmt.Errorf("end (%d) cannot be equal to or smaller than start (%d)", end, start)
	}
	if end > len(str) {
		return "", fmt.Errorf("end (%d) is out of range", end)
	}
	var startByte = -1
	var runeIndex int
	for i := range str {
		if runeIndex == start {
			startByte = i
			if end == 0 {
				return str[startByte:], nil
			}
		}
		if end != 0 && runeIndex == end {
			return str[startByte:i], nil
		}
		runeIndex++
	}
	if startByte < 0 {
		return "", fmt.Errorf("start (%d) is out of range (%d)", start, runeIndex)
	}
	if end == runeIndex {
		return str[startByte:], nil
	}
	return "", fmt.Errorf("end (%d) is out of range (%d)", end, runeIndex)
}
