package vo

type TotalData struct {
	User    int64 `json:"user"`
	Video   int64 `json:"video"`
	Review  int64 `json:"review"`
	Message int64 `json:"message"`
}

type OneDayData struct {
	User  int64  `json:"user"`
	Video int64  `json:"video"`
	Date  string `json:"date"`
}

type ServerInfoVo struct {
	Web     FEInfoVo `json:"web"`     //前端
	Version string   `json:"version"` //后端
	Redis   string   `json:"redis"`
	FFmpeg  string   `json:"ffmpeg"`
	MySQl   string   `json:"mysql"`
}

type FEInfoVo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Cover   string `json:"cover"`
	Author  string `json:"author"`
	Desc    string `json:"desc"`
}
