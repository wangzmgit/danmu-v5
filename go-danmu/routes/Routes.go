package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api/v1")
	{
		GetUserRoutes(v1)        //用户路由
		GetPartitionRoutes(v1)   //分区路由
		GetVideoRoutes(v1)       //视频列表
		GetReviewRoutes(v1)      //审核接口
		GetArchiveRoutes(v1)     //点赞收藏
		GetFollowRoutes(v1)      //关注路由
		GetCommentRoutes(v1)     //评论回复
		GetMessageRoutes(v1)     //消息公告
		GetDanmakuRoutes(v1)     //弹幕接口
		GetCarouselRoutes(v1)    //轮播图接口
		GetWebsiteDataRoutes(v1) //网站数据接口
		GetConfigRoutes(v1)      //配置接口
		GetUploadRoutes(v1)      //文件上传接口
	}

	//获取静态文件
	r.StaticFS("/api/avatar", http.Dir("./file/avatar"))
	r.StaticFS("/api/cover", http.Dir("./file/cover"))
	r.StaticFS("/api/video", http.Dir("./file/video"))
	r.StaticFS("/api/carousel", http.Dir("./file/carousel"))

	return r
}
