package vo

import "time"

type CarouselVo struct {
	ID        uint      `json:"id"`
	Img       string    `json:"img"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
