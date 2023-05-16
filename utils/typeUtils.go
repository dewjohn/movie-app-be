package utils

import "strconv"

// 字符串转int
func StringToInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return res
}

// 字符串转uint
func StringToUint(v string) uint {
	res, err := strconv.Atoi(v)
	if err != nil {
		return uint(0)
	}
	return uint(res)
}

// 字符串数组去重
func StringArrayUnique(arr []string) []string {
	m := make(map[string]bool)
	result := []string{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			m[item] = true
			result = append(result, item)
		}
	}
	return result
}

// 数字数组去重
func IntArrayUnique(arr []int) []int {
	m := make(map[int]bool)
	result := []int{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			m[item] = true
			result = append(result, item)
		}
	}
	return result
}
