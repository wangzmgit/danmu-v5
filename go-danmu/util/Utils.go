package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"time"
)

/********************************************************************************
* @功    能：生成随机数字
* @输入参数：随机数位数
* @返 回 值：随机数
* @日    期：2021/7/23
* @版    本：1.0
********************************************************************************/
func RandomCode(n int) string {
	var letters = []byte("1234567890")
	result := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

/********************************************************************************
* @功    能：邮箱格式匹配
* @输入参数：邮箱字符串
* @返 回 值：邮箱格式是否正确
* @日    期：2021/7/10
* @版    本：1.0
********************************************************************************/
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/********************************************************************************
* @功    能：隐藏邮箱
* @输入参数：邮箱字符串
* @返 回 值：隐藏后的邮箱
* @日    期：2021/10/21
* @版    本：1.0
********************************************************************************/
func HideEmail(email string) string {
	pattern := `(\w?)(\w+)(\w)(@\w+\.[a-z]+(\.[a-z]+)?)`
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(email, "$1****$3$4")
}

/********************************************************************************
* @功    能：计算字符MD5
* @输入参数：字符串
* @返 回 值：MD5字符串
* @日    期：2022年7月2日15:52:53
* @版    本：1.0
********************************************************************************/
func GetStringMD5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}
