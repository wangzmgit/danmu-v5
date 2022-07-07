package util

import "strconv"

//字符串转int
func StringToInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return res
}

//字符串转uint
func StringToUint(v string) uint {
	res, err := strconv.Atoi(v)
	if err != nil {
		return uint(0)
	}
	return uint(res)
}
