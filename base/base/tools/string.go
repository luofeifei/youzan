package tools

import (
	"strconv"
	"time"
)

func StringToInt64(e string) (res int64) {
	res, _ = strconv.ParseInt(e, 10, 64)
	return res
}

func StringFloatToInt(e string, prec int) (int, error) {
	s, _ := strconv.ParseFloat(e, 64)
	return strconv.Atoi(strconv.FormatFloat(s, 'f', prec, 64))
}

func StringToFloat64(e string) (float64, error) {
	return strconv.ParseFloat(e, 64)
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrntTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetCurrntTime() time.Time {
	return time.Now()
}