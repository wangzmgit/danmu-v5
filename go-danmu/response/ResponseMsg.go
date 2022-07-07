package response

const (
	OK                     = "ok"
	ParameterError         = "参数错误"
	RequestError           = "请求错误"
	SystemError            = "服务器出错了"
	OperationTooFrequently = "操作过于频繁"
	SendFail               = "发送失败"
	VerificationFail       = "验证失败"
	CreateFail             = "上传失败"
	ModifyFail             = "修改失败"
	DeleteFail             = "删除失败"
	UploadFail             = "上传失败"
	PageOrSizeError        = "页码或请求数量错误"
	TooManyRequests        = "请求数量过多"

	//用户相关
	NickCheck             = "昵称不能为空"
	NameExist             = "用户名已存在"
	EmailRegistered       = "该邮箱已经被注册了"
	EmailFormatCheck      = "邮箱格式有误"
	PasswordCheck         = "密码不能少于6位"
	VerificationCodeError = "验证码错误"
	GenderCheck           = "性别选择有误"
	UserNotExist          = "用户不存在"
	NameOrPasswordError   = "用户名和密码错误"
	BirthdayFormatCheck   = "请输入正确的出生日期"
	InsufficientAuthority = "权限不足"
	RoleNotExist          = "角色不存在"

	//分区相关
	PartitionNotExist       = "分区不存在"
	ParentPartitionNotExist = "所属分区不存在"
	PartitionContentCheck   = "分区名称不能为空"

	//视频相关
	TranscodingError  = "视频处理出现错误"
	TitleCheck        = "标题不能为空"
	CoverCheck        = "请上传封面图片"
	PartitionCheck    = "未选择分区"
	VideoNotExist     = "视频不存在"
	UpdateStatusFail  = "视频状态更新失败"
	NotExistResources = "视频资源不存在"

	//点赞、收藏、关注
	IsCollect              = "已经收藏过了"
	IsLike                 = "已经点过赞了"
	NotCollect             = "还没有收藏"
	NotLike                = "还没有点赞"
	CantFollowYourself     = "不能关注自己"
	CreateCollectionFail   = "创建收藏夹失败"
	CollectionNotExist     = "收藏夹不存在"
	CollectionNameNotExist = "收藏夹名不能为空"

	//评论回复
	CommentContentCheck = "评论内容不能为空"

	//消息公告
	AnnounceContentCheck = "公告内容不能为空"
	MessageNotExist      = "消息不存在"
	CantSendYourself     = "不能发送给自己"
	UserMsgContentCheck  = "不能发送空消息"
	SendMsgFail          = "消息发送失败"

	//弹幕
	DanmakuCheck = "不能发送空弹幕"

	//轮播图
	ImgNotExist      = "图片不存在"
	CarouselNotExist = "轮播图不存在"

	//文件相关
	FileUploadFail    = "文件上传失败"
	FileSaveFail      = "文件保存失败"
	FileCheckFail     = "文件不符合要求"
	FileSizeCheckFail = "文件大小不符合要求"
)
