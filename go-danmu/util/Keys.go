package util

import (
	"fmt"
)

//注册验证码
func RegCode(email string) string {
	return fmt.Sprintf("danmu:reg:%s", email)
}

//视频点击量限制
func VideoClicksLimitKey(vid int, ip string) string {
	return fmt.Sprintf("danmu:clicks:%d:%s", vid, ip)
}
