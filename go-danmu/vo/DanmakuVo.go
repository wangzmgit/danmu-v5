package vo

type DanmakuVo struct {
	Time  uint   `json:"time"`
	Type  int    `json:"type"`
	Color string `json:"color"`
	Text  string `json:"text"`
}
