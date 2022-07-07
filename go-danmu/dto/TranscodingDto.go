package dto

type ResDto struct {
	Res360   string
	Res480   string
	Res720   string
	Res1080  string
	Original string
}

type VideoInfoData struct {
	Stream []Streams `json:"streams"`
	Format Format    `json:"format"`
}

type Streams struct {
	CodecName string `json:"codec_name"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	PixFmt    string `json:"pix_fmt,omitempty"`
	Duration  string `json:"duration"`
}

type Format struct {
	BitRate string `json:"bit_rate"`
}
