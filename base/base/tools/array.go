package tools

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"strings"
)

//ArrayToString 数组转字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

//MergeDuplicateStringArray 合并两个字符串数组并去重
func MergeDuplicateIntArray(slice []int, elems []int) []int {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []int
	for i := range t.Iterator().C {
		result = append(result, i.(int))
	}
	return result
}

//MergeDuplicateStringArray 合并两个字符串数组并去重
func MergeDuplicateStringArray(slice []string, elems []string) []string {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []string
	for i := range t.Iterator().C {
		result = append(result, i.(string))
	}
	return result
}

// 查找2个int64数组返回之间的相同值
func IntersectInt(slice []int64, elems []int64) (res []int64) {
	t := mapset.NewSet()
	for _, i := range slice {
		t.Add(i)
	}
	e := mapset.NewSet()
	for _, i := range elems {
		e.Add(i)
	}
	for _, i := range t.Intersect(e).ToSlice() {
		res = append(res, i.(int64))
	}
	return res
}

// 查找2个int64数组返回之间的差异
func DifferenceInt(slice []int64, elems []int64) (res []int64) {
	t := mapset.NewSet()
	for _, i := range slice {
		t.Add(i)
	}
	e := mapset.NewSet()
	for _, i := range elems {
		e.Add(i)
	}
	if len(slice) > len(elems) {
		for _, i := range t.Difference(e).ToSlice() {
			res = append(res, i.(int64))
		}
	} else {
		for _, i := range e.Difference(t).ToSlice() {
			res = append(res, i.(int64))
		}
	}
	return res
}

//RemoveRepeatedElement 数组去重 string
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//RemoveRepeatedElementInt64 数组去重 int64
func RemoveRepeatedElementInt64(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
