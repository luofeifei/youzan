package tools

import (
	"hash/fnv"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// 判断来源是否合法
func IsPlatform(plat string) (err bool) {
	return plat == "wx" || plat == "h5" || plat == "app" || plat == "web" || plat == "ali"
}

func IsDigital(input string )(res bool){

	getInt,err:=strconv.Atoi(input)
	if getInt>0{
		res = true
	}
	if err!=nil{
		res = false
	}
	return res
}
func IsIdCard(input string )(res bool){
	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	res=true
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`,input); !m {
		 res =false
	}
	return res
}
func IsRuleTime(input string )(res bool){

	const layoutISO = "2006-01-02"
	t, _ := time.ParseInLocation(layoutISO, input,time.Local)
	return input==t.Format(layoutISO)
}
func IsEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func IsMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

