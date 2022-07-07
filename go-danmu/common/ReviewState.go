package common

//视频审核状态
const (
	CreateVideo       = 100  //成功创建视频
	VideoProcessing   = 200  //视频处理中
	WaitingReview     = 1000 //等待审核
	AuditApproved     = 2000 //审核通过
	WrongVideoInfo    = 3100 //视频信息存在问题
	WrongVideoContent = 3200 //视频内容存在问题
	ProcessingFail    = 3300 //视频处理失败
)
