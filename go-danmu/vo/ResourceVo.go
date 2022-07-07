package vo

import (
	"kuukaa.fun/danmu-v5/model"
)

type ResourceVo struct {
	ID uint `json:"id"`
	//分P使用的标题
	Title string `json:"title"`
	//不同分辨率
	Res360   string `json:"res360"`
	Res480   string `json:"res480"`
	Res720   string `json:"res720"`
	Res1080  string `json:"res1080"`
	Original string `json:"original"`
	//时长
	Duration float64 `json:"duration"`
	//审核状态
	Review int `json:"review"`
}

func ToReviewResourceListVo(resources []model.Resource) []ResourceVo {
	length := len(resources)
	newResources := make([]ResourceVo, length)
	for i := 0; i < length; i++ {
		newResources[i].ID = resources[i].ID
		newResources[i].Title = resources[i].Title
		newResources[i].Res360 = resources[i].Res360
		newResources[i].Res480 = resources[i].Res480
		newResources[i].Res720 = resources[i].Res720
		newResources[i].Res1080 = resources[i].Res1080
		newResources[i].Original = resources[i].Original
		newResources[i].Duration = resources[i].Duration
		newResources[i].Review = resources[i].Review
	}
	return newResources
}
