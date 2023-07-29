package routes

import (
	v1 "ginStudy/api/v1"
	"ginStudy/config"
	"ginStudy/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 上线时一定要改成 release 模式
	gin.SetMode(config.Config.AppMode)
	r := gin.New()
	// 使用自定义的中间件
	r.Use(config.GinLogger(), config.GinRecovery(true), middleware.Cors())
	// v1版本客户端业务接口
	auth := r.Group("/api/v1")

	// 获取短信验证码
	auth.GET("/user/sms_code", v1.SendSMSCodeHandler)
	// 用户通过手机号 + 密码注册
	auth.POST("/user/signup_mobile", v1.MobileSignUpHandler)
	// 用户通过手机号 + 密码登录
	auth.POST("/user/login_mobile", v1.MobileLoginHandler)
	// 用户通过手机号 + 验证码登录
	auth.POST("/user/login_sms_code", v1.SMSCodeLoginHandler)
	// 用户通过验证码 + 密码重置密码
	auth.PUT("/user/neglect/password", v1.FindPwdBySMSCodeHandler)
	auth.Use(middleware.JWTAuthMiddleware(), middleware.TimeoutMiddleware())
	{
		// 获取邮箱验证码
		auth.GET("/user/email_code", v1.SendEmailCodeHandler)
		// 用户退出登录
		auth.POST("/user/logout", v1.LogoutHandler)
		// 根据旧密码修改新密码
		auth.PUT("/user/modify_pwd_oldpwd", v1.ModifyPwdByOldPwdHandler)
		// 根据验证码修改新密码
		auth.PUT("/user/modify_pwd_smscode", v1.ModifyPwdBySMSCodeHandler)
		// 绑定邮箱
		auth.PUT("/user/improvement/email", v1.BindEmailHandler)
		// 获取用户信息
		auth.GET("/user/userInfo", v1.GetUserInfoHandler)
		// 获取用户词书等级
		auth.GET("/user/vocabulary", v1.GetVocabularyHandler)
		// 用户修改个人信息
		auth.PUT("/user/modify/userInfo", v1.ModifyUserInfoHandler)
		// 重置词书
		auth.PUT("/user/reset/vocabulary", v1.ResetEngLevelHandler)
		// 设置词书
		auth.PUT("/user/modify/vocabulary", v1.ModifyEngLevelHandler)
		// 设置每日单词要背的量
		auth.PUT("/user/modify/daily/word", v1.ModifyDailyWordHandler)
		// 获取最新的学习记录
		auth.GET("/user/learn/up_to_date", v1.GetUpToDateLearnHandler)
		// 打卡
		auth.POST("/user/learn/finish", v1.FinishLearnHandler)
		// 补打卡
		auth.POST("/user/learn/check", v1.CheckInLearnHandler)
		// 用户获取打卡日期列表
		auth.GET("/user/check/list", v1.CheckLearnListHandler)
		// 用户进入我的页面时获取数据
		auth.GET("/user/mine", v1.MineInfoHandler)
		// 用户点日历看数据
		auth.GET("/user/calendar", v1.CalendarUserInfoHandler)

		// 用户注销账号（关联的一切都要删除，硬删除，并放入注销表（60天冷静期））

		// 获取单词意思（query传单词英文）
		auth.GET("/word/meaning", v1.GetWordMeaningHandler)
		// 记单词时获取单词详情（query传单词id）
		auth.GET("/word/memory/meaning", v1.GetWordDetailHandler)
		// 判断用户是否背完这本书 传 word_type 背完返回1，没有背完返回0    （改service里面的条数判断）
		auth.GET("/word/vocabulary/termination", v1.DetermineVocabularyHandler)
		// 用户获取文章单词列表
		auth.GET("/word/collection/list", v1.GetCollectWordListHandler)
		// 收藏单词
		auth.POST("/word/collection", v1.CollectWordHandler)
		// 取消收藏单词
		auth.DELETE("/word/cancellation/collection", v1.CancelCollectWordHandler)
		// 按词书获取单词列表
		auth.GET("/word/list", v1.GetWordListHandler)

		// 获取文章列表(redis) 还有细分
		auth.GET("/essay/list", v1.GetEssayListHandler)
		// 用户点进去具体文章获取文章信息
		auth.GET("/essay", v1.GetEssayByEssayIdHandler)
		// 用户获取文章收藏列表
		auth.GET("/essay/collection/list", v1.GetCollectEssayListHandler)
		// 收藏文章
		auth.POST("/essay/collection", v1.CollectEssayHandler)
		// 取消收藏文章
		auth.DELETE("/essay/cancellation/collection", v1.CancelCollectEssayHandler)
		// 获取文章列表
		auth.GET("/essays", v1.GetAllEssayHandler)

		// 获取听力列表(redis) 传类型
		auth.GET("/listen/list", v1.GetListenListHandler)
		// 用户点进去具体听力获取听力信息
		auth.GET("/listen", v1.GetListenByListenIdHandler)
		// 用户获取听力收藏列表
		auth.GET("/listen/collection/list", v1.GetCollectListenListHandler)
		// 收藏听力
		auth.POST("/listen/collection", v1.CollectListenHandler)
		// 取消收藏听力
		auth.DELETE("/listen/cancellation/collection", v1.CancelCollectListenHandler)

		// 获取每日一句
		auth.GET("/sentence/daily", v1.DailySentenceHandler)
		// 用户点进去具体句子获取句子信息
		auth.GET("/sentence", v1.GetSentenceBySentenceIdHandler)
		// 用户获取句子收藏列表
		auth.GET("/sentence/collection/list", v1.GetCollectSentenceListHandler)
		// 收藏句子
		auth.POST("/sentence/collection", v1.CollectSentenceHandler)
		// 取消收藏句子
		auth.DELETE("/sentence/cancellation/collection", v1.CancelCollectSentenceHandler)
		auth.POST("/word/post", v1.WordHandler)
		auth.GET("/word/get", v1.WordGetHandler)
	}

	// v1版本后台业务接口
	admin := r.Group("/admin/v1")
	// 管理员登录
	admin.POST("/backend/login", v1.AdminLoginHandler)

	admin.Use(middleware.JWTAuthAdminMiddleware())
	{
		// 数据清库的操作，慎重
		admin.POST("/ljr/delete/empty/db", v1.RefreshDbHandler)
		// 管理员退出登录
		admin.POST("/backend/logout", v1.AdminLogoutHandler)

		// 传入要爬取的文章类型
		admin.POST("/essay/spider/all", v1.EssaySpiderTotalHandler)
		// 手动添加文章
		admin.POST("/essay/insertion", v1.InsertEssayHandler)
		// 删除文章(用户的收藏取消，对应的文章单词也失效)
		admin.DELETE("/essay/delete", v1.DeleteEssayHandler)
		// 根据文章id获取文章具体内容
		admin.GET("/essay", v1.GetEssayHandler)
		// 根据文章id修改文章
		admin.PUT("/essay/modify", v1.UpdateEssayHandler)
		// 获取所有文章(分页)
		admin.GET("/essays", v1.GetEssayAllHandler)

		// 传入要爬取的听力类型
		admin.POST("/listen/spider/all", v1.ListenSpiderTotalHandler)
		// 手动添加听力
		admin.POST("/listen/insertion", v1.InsertListenHandler)
		// 删除听力(用户的收藏取消，对应的听力单词也失效)
		admin.DELETE("/listen/delete", v1.DeleteListenHandler)
		// 根据听力id获取听力具体内容
		admin.GET("/listen", v1.GetListenHandler)
		// 根据听力id修改听力
		admin.PUT("/listen/modify", v1.UpdateListenHandler)
		// 获取所有听力(分页)
		admin.GET("/listens", v1.GetListenAllHandler)

		// 传入要爬取的单词类型
		admin.POST("/word/spider/all", v1.WordSpiderTotalHandler)
		// 手动添加单词
		admin.POST("/word/insertion", v1.InsertWordHandler)
		// 删除单词(用户的收藏取消)
		admin.DELETE("/word/delete", v1.DeleteWordHandler)
		// 根据单词id获取单词具体内容
		admin.GET("/word", v1.GetWordHandler)
		// 输入单词进行查找
		admin.GET("/word/signal", v1.GetWordByWordHandler)
		// 根据单词id修改单词
		admin.PUT("/word/modify", v1.UpdateWordHandler)
		// 获取所有单词(分页)
		admin.GET("/words", v1.GetWordAllHandler)

		// 获取从2019-11-26以来所有的每日一句，写入数据库
		admin.POST("/sentence/spider/all", v1.DailySentenceTotalSpiderHandler)

		// 禁用用户(软删除)
		admin.DELETE("/user/prohibit", v1.DisableUserHandler)
		// 恢复用户权限
		admin.PUT("/user/recovery", v1.RecoverUserHandler)
		// 获取用户列表(分页)
		admin.GET("/users", v1.GetUserAllHandler)
		// 输入用户手机号进行查找
		admin.GET("/user/signal", v1.GetUserByMobileHandler)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "404, page not exists!",
			"data": nil,
		})
	})
	return r
}
