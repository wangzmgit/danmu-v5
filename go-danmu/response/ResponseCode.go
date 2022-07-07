package response

import "github.com/gin-gonic/gin"

const (
	SuccessCode     = 2000
	FailCode        = 4000
	CheckFailCode   = 4022
	ServerErrorCode = 5000
)

type ResponseStruct struct {
	HttpStatus int    //http状态
	Code       int    //状态码
	Data       gin.H  //数据
	Msg        string //信息
}
