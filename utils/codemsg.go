package utils

import "errors"

const (
	SUCCESS       = 200
	FAIL_BUSINESS = 403
	FAIL_ENTITY   = 405
	FAIL_TOKEN    = 401

	ERROR          = 500
	PAGE_NOT_FOUND = 404
)

// 100xx 基本常用错误码
const (
	REQ_PARAM_LACK_ERR = 10001 + iota
	REQ_PARAM_INVALID_ERR
	RPC_INVOKE_ERR
	INTERNAL_ERR
	TIMEOUT_ERR
	UNKNOW_ERR
	JSON_ERR
	DB_ERR
	DB_SEARCH_ERR //10009
	DB_CREATE_ERR
	DB_UPDATE_ERR
	DB_DELETE_ERR
	REQ_INIT_ERR //10013
	PAGE_ERR     //10014
)

// 200xx 鉴权相关错误
const (
	ERROR_TOKEN_EXIST = 20001 + iota
	ERROR_TOKEN_WRONG
	ERROR_TOKEN_TYPE_WRONG
	ERROR_TOKEN_RUNTIME //4
	ERROR_TOKEN_INVALID
	ERROR_ACTOKEN_CREATE
	ERROR_RETOKEN_CREATE
	ERROR_TOKEN_PARSE
	ERROR_TOKEN_LOGOUT //9
)

// 300xx 用户业务相关错误
const (
	USER_NOTEXIST_ERR                = 30001
	LOGIN_PASSWORD_ERR               = 30002
	EMAIL_ERR                        = 30003
	MOBILE_NULL_ERR                  = 30004
	MOBILE_TYPE_ERR                  = 30005
	SMS_ERR                          = 30006
	MOBILE_ISEXIST_ERR               = 30007
	CODE_SET_ERR                     = 30008
	CODE_CACHE_INVALID_ERR           = 30009
	CODE_INPUT_ERR                   = 30010
	EMAIL_NULL_ERR                   = 30011
	EMAIL_TYPE_ERR                   = 30012
	UPDATE_PWD_ERR                   = 30013
	UPDATE_EMAIL_ERR                 = 30014
	USERINFO_NOTEXIST_ERR            = 30015
	INTEGRAL_UPDATE_ERR              = 30016
	USERID_PARAM_ERR                 = 30017
	USER_STATUS_DISABLED_ERR         = 30018
	USERID_PARAM_DISABLED_ERR        = 30019
	USERINFO_UPDATE_ERR              = 30020
	USER_OLDPASSWORD_ERR             = 30021
	USERINFO_LEXUEAPPID_EXIST_ERR    = 30022
	USERINFO_INVITACODE_EXIST_ERR    = 30023
	USERINFO_INVITACODE_DISABLED_ERR = 30024
	USERID_NULL_ERR                  = 30025
	USER_DISABLE_ERR                 = 30026
	USER_RECOVER_ERR                 = 30027
	ENGLEVEL_UPDATE_ERR              = 30028
	WORDNEEDRECITENUM_UPDATE_ERR     = 30029
	UPDATE_LASTLOGINTIME_ERR         = 30030
	USER_VOCABULARY_ERR              = 30031
	USERDATE_DATE_GET_ERR            = 30032
	USERDATE_UPDATE_ERR              = 30033
	ADMINUSER_PARAM_NULL_ERR         = 30034
	ADMINUSER_NOTEXIST_ERR           = 30035
	ADMINUSER_PASSWORD_ERR           = 30036
	USERLIST_NULL_ERR                = 30037
	USERCHECK_ERR                    = 30038
	USERIN_ERR                       = 30039
	USERCHECKLIST_ERR                = 30040
)

// 400xx 单词业务相关错误
const (
	DELETE_ESSAYWORD_ERR      = 40001
	ESSAYWORD_GETLIST_ERR     = 40002
	DELETE_LISTENWORD_ERR     = 40003
	WORD_TRANSLATION_ERR      = 40004
	WORD_NULL_ERR             = 40005
	LISTENWORD_GETLIST_ERR    = 40006
	WORD_TYPE_ERR             = 40007
	WORDID_NULL_ERR           = 40008
	WORD_NOTEXIST_ERR         = 40009
	WORD_NOTEXIST_COLLECT_ERR = 40010
	WORD_COLLECT_ERR          = 40011
	WORD_COLLECT_ISEXIST_ERR  = 40012
	WORD_CANCEL_COLLECT_ERR   = 40013
	WORD_COLLECT_LIST_ERR     = 40014
	WORD_EXIST_ERR            = 40015
	WORD_UPDATE_ERR           = 40016
	DELETE_WORD_ERR           = 40017
	WORDTYPE_NULL_ERR         = 40018
	WORDLEARN_NULL_ERR        = 40019
	WORDLEARN_RESET_ERR       = 40020
	WORDLIST_NULL_ERR         = 40021
	VOCABULARY_TYPE_ERR       = 40022
)

// 500xx 阅读业务相关错误
const (
	DELETE_ESSAY_ID_ERR               = 50001
	DELETE_ESSAY_ERR                  = 50002
	ESSAY_NOTEXIST_ERR                = 50003
	ESSAY_UPDATE_ERR                  = 50004
	ESSAY_EXIST_ERR                   = 50005
	ESSAY_COLLECT_ISEXIST_ERR         = 50006
	ESSAY_COLLECT_ERR                 = 50007
	ESSAY_CANCEL_COLLECT_ERR          = 50008
	ESSAY_NOTEXIST_COLLECT_ERR        = 50009
	ESSAYLIST_GET_ERR                 = 50010
	ESSAY_NOTEXIST_CANCEL_COLLECT_ERR = 50011
	ESSAY_TYPE_ERR                    = 50012
	ESSAY_CACHE_ERR                   = 50013
	ESSAY_CACHE_GET_ERR               = 50014
	ESSAY_COLLECT_LIST_ERR            = 50015
	ESSAYLIST_NULL_ERR                = 50016
)

// 600xx 听力业务相关错误
const (
	LISTEN_EXIST_ERR                   = 60001
	DELETE_LISTEN_ID_ERR               = 60002
	DELETE_LISTEN_ERR                  = 60003
	LISTEN_CANCEL_COLLECT_ERR          = 60004
	LISTEN_NOTEXIST_ERR                = 60005
	LISTEN_UPDATE_ERR                  = 60006
	LISTEN_TYPE_ERR                    = 60007
	LISTENLIST_GET_ERR                 = 60008
	LISTEN_CACHE_ERR                   = 60009
	LISTEN_CACHE_GET_ERR               = 60010
	LISTEN_NOTEXIST_COLLECT_ERR        = 60011
	LISTEN_COLLECT_ERR                 = 60012
	LISTEN_COLLECT_ISEXIST_ERR         = 60013
	LISTEN_NOTEXIST_CANCEL_COLLECT_ERR = 60014
	LISTEN_COLLECT_LIST_ERR            = 60015
	LISTENLIST_NULL_ERR                = 60016
)

// 700xx 句子业务相关错误
const (
	SENTENCE_EXIST_ERR            = 70001
	SENTENCE_SPIDER_ERR           = 70002
	PARAM_DISABLED_ERR            = 70003
	SENTENCE_NOTEXIST_ERR         = 70004
	SENTENCE_NOTEXIST_COLLECT_ERR = 70005
	SENTENCE_UPDATE_ERR           = 70006
	SENTENCE_COLLECT_ERR          = 70007
	SENTENCE_COLLECT_ISEXIST_ERR  = 70008
	SENTENCE_CANCEL_COLLECT_ERR   = 70009
	SENTENCEID_NULL_ERR           = 70010
	SENTEN_COLLECT_LIST_ERR       = 70011
)

// 900xx 配置初始化错误或其他配置错误
const (
	VIPER_READ_ERR           = 90001
	VIPER_UNMARSHAL_ERR      = 90002
	INIT_SETTING_ERR         = 90003
	INIT_LOGGER_ERR          = 90004
	MYSQL_INIT_ERR           = 90005
	MYSQL_GETINSTANCE_ERR    = 90006
	MYSQL_CLOSE_ERR          = 90007
	SONWFLAKE_NOT_INIT       = 90008
	SONWFLAKE_INIT_ERR       = 90009
	REDIS_INIT_ERR           = 90010
	REDIS_CLOSE_ERR          = 90011
	TRANSLATE_INIT_ERR       = 90012
	SELF_VERIFY_REGISTER_ERR = 90013
	SELF_TRANSLATE_ERR       = 90014
	APISMS_SEND_ERR          = 90015
)

var codeMsg = map[int]string{
	SUCCESS:               "操作成功",
	ERROR:                 "服务器错误，无法完成请求",
	REQ_PARAM_LACK_ERR:    "请求缺少参数",
	REQ_PARAM_INVALID_ERR: "请求参数无效",
	RPC_INVOKE_ERR:        "远程调用出错",
	INTERNAL_ERR:          "内部错误",
	TIMEOUT_ERR:           "请求超时",
	PAGE_NOT_FOUND:        "页面找不到",
	UNKNOW_ERR:            "未知错误",
	JSON_ERR:              "JSON序列化错误",
	DB_ERR:                "数据库错误",
	DB_SEARCH_ERR:         "查询失败",
	DB_CREATE_ERR:         "创建失败",
	DB_UPDATE_ERR:         "更新失败",
	DB_DELETE_ERR:         "删除失败",
	REQ_INIT_ERR:          "请求发起错误",
	PAGE_ERR:              "页数和偏移量不能为空",

	ERROR_TOKEN_EXIST:      "TOKEN不存在，无权限访问，请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确，请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误，请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期，请重新登陆",
	ERROR_TOKEN_INVALID:    "TOKEN无效，请重新登陆",
	ERROR_ACTOKEN_CREATE:   "ACCESS TOKEN颁发失败，请重新生成",
	ERROR_RETOKEN_CREATE:   "REFRESH TOKEN颁发失败，请重新生成",
	ERROR_TOKEN_PARSE:      "TOKEN解析错误",
	ERROR_TOKEN_LOGOUT:     "TOKEN已注销，请重新登陆",

	VIPER_READ_ERR:           "配置文件读取错误",
	VIPER_UNMARSHAL_ERR:      "配置文件反序列化至结构体失败",
	INIT_SETTING_ERR:         "初始化配置信息失败",
	INIT_LOGGER_ERR:          "初始化日志失败",
	MYSQL_INIT_ERR:           "创建mysql客户端失败",
	MYSQL_GETINSTANCE_ERR:    "获取mysql客户端实例失败",
	MYSQL_CLOSE_ERR:          "mysql关闭失败",
	SONWFLAKE_NOT_INIT:       "雪花算法没有初始化",
	SONWFLAKE_INIT_ERR:       "雪花算法初始化错误",
	REDIS_INIT_ERR:           "redis初始化失败",
	REDIS_CLOSE_ERR:          "redis关闭失败",
	TRANSLATE_INIT_ERR:       "初始化翻译器失败",
	SELF_VERIFY_REGISTER_ERR: "自定义校验方法出错",
	SELF_TRANSLATE_ERR:       "自定义翻译方法出错",
	APISMS_SEND_ERR:          "短信发送失败，请重新获取",

	USER_NOTEXIST_ERR:                "该用户不存在",
	LOGIN_PASSWORD_ERR:               "密码输入错误",
	EMAIL_ERR:                        ">_<|||请大佬重新获取一下邮箱验证码",
	MOBILE_NULL_ERR:                  "手机号码不能为空",
	MOBILE_TYPE_ERR:                  "手机号码格式错误",
	SMS_ERR:                          ">_<|||请大佬重新获取一下短信验证码",
	MOBILE_ISEXIST_ERR:               "该手机号已注册",
	CODE_SET_ERR:                     "验证码获取失败",
	CODE_CACHE_INVALID_ERR:           "验证码失效，请重新获取",
	CODE_INPUT_ERR:                   "验证码输入错误",
	EMAIL_NULL_ERR:                   "邮箱号不能为空",
	EMAIL_TYPE_ERR:                   "邮箱号格式错误",
	UPDATE_PWD_ERR:                   "更新密码失败，请稍后重试",
	UPDATE_EMAIL_ERR:                 "绑定邮箱失败",
	USERINFO_NOTEXIST_ERR:            "该用户的用户信息不存在",
	INTEGRAL_UPDATE_ERR:              "用户积分更新失败",
	USERID_PARAM_ERR:                 "用户id为空，获取用户信息失败",
	USER_STATUS_DISABLED_ERR:         "你的账号已被禁用，请联系管理员解禁",
	USERID_PARAM_DISABLED_ERR:        "请求信息有误",
	USERINFO_UPDATE_ERR:              "用户信息更新失败",
	USER_OLDPASSWORD_ERR:             "旧密码输入不正确",
	USERINFO_LEXUEAPPID_EXIST_ERR:    "该ID已被使用，修改失败",
	USERINFO_INVITACODE_EXIST_ERR:    "该邀请码已存在",
	USERINFO_INVITACODE_DISABLED_ERR: "该邀请码无效，请重新输入",
	USERID_NULL_ERR:                  "用户id为空",
	USER_DISABLE_ERR:                 "用户禁用失败",
	USER_RECOVER_ERR:                 "用户权限恢复失败",
	ENGLEVEL_UPDATE_ERR:              "用户词书更新失败",
	WORDNEEDRECITENUM_UPDATE_ERR:     "用户每日单词量更新失败",
	UPDATE_LASTLOGINTIME_ERR:         "更新登录时间失败，请稍后重试",
	USER_VOCABULARY_ERR:              "用户词书等级获取失败",
	USERDATE_DATE_GET_ERR:            "没有学习记录",
	USERDATE_UPDATE_ERR:              "用户更新打卡记录失败",
	ADMINUSER_PARAM_NULL_ERR:         "管理员用户名或密码不能为空",
	ADMINUSER_NOTEXIST_ERR:           "管理员不存在",
	ADMINUSER_PASSWORD_ERR:           "密码不正确",
	USERLIST_NULL_ERR:                "用户资源列表加载失败",
	USERCHECK_ERR:                    "今天你已经打卡啦",
	USERIN_ERR:                       "积分不够打卡哦",
	USERCHECKLIST_ERR:                "该日无打卡记录",

	DELETE_ESSAYWORD_ERR:      "删除文章中的单词失败",
	ESSAYWORD_GETLIST_ERR:     "文章单词获取失败",
	DELETE_LISTENWORD_ERR:     "删除听力中的单词失败",
	WORD_TRANSLATION_ERR:      "字典中不存在该单词",
	WORD_NULL_ERR:             "单词不能为空哦",
	LISTENWORD_GETLIST_ERR:    "听力单词获取失败",
	WORD_TYPE_ERR:             "单词类型错误",
	WORDID_NULL_ERR:           "单词id不能为空哦",
	WORD_NOTEXIST_ERR:         "单词不存在",
	WORD_NOTEXIST_COLLECT_ERR: "单词不存在，收藏失败，请联系管理员处理",
	WORD_COLLECT_ERR:          "单词收藏失败",
	WORD_COLLECT_ISEXIST_ERR:  "您已收藏该单词，请勿重复操作",
	WORD_CANCEL_COLLECT_ERR:   "单词取消收藏失败",
	WORD_COLLECT_LIST_ERR:     "单词收藏列表获取失败",
	WORD_EXIST_ERR:            "单词已存在",
	WORD_UPDATE_ERR:           "单词信息更新失败",
	DELETE_WORD_ERR:           "单词删除失败",
	WORDTYPE_NULL_ERR:         "词书类型不能为空哦",
	WORDLEARN_NULL_ERR:        "用户还未学习",
	WORDLEARN_RESET_ERR:       "词书重置失败",
	WORDLIST_NULL_ERR:         "单词资源列表加载失败",
	VOCABULARY_TYPE_ERR:       "词书类型错误",

	DELETE_ESSAY_ID_ERR:               "文章id无效，请重新获取文章",
	DELETE_ESSAY_ERR:                  "文章删除失败",
	ESSAY_NOTEXIST_ERR:                "文章不存在",
	ESSAY_UPDATE_ERR:                  "文章信息更新失败",
	ESSAY_EXIST_ERR:                   "文章已存在",
	ESSAY_COLLECT_ISEXIST_ERR:         "您已收藏该文章，请勿重复操作",
	ESSAY_COLLECT_ERR:                 "文章收藏失败",
	ESSAY_CANCEL_COLLECT_ERR:          "文章取消收藏失败",
	ESSAY_NOTEXIST_COLLECT_ERR:        "文章不存在，收藏失败，请联系管理员处理",
	ESSAYLIST_GET_ERR:                 "文章列表获取失败，请重新刷新",
	ESSAY_NOTEXIST_CANCEL_COLLECT_ERR: "文章不存在，取消收藏失败，请联系管理员处理",
	ESSAY_TYPE_ERR:                    "文章类型错误",
	ESSAY_CACHE_ERR:                   "文章缓存设置失败",
	ESSAY_CACHE_GET_ERR:               "文章缓存获取失败",
	ESSAY_COLLECT_LIST_ERR:            "文章收藏列表获取失败",
	ESSAYLIST_NULL_ERR:                "文章资源列表加载失败",

	SENTENCE_EXIST_ERR:            "句子已存在",
	SENTENCE_SPIDER_ERR:           "句子爬取失败，请稍后重试",
	PARAM_DISABLED_ERR:            "请求信息有误",
	SENTENCE_NOTEXIST_ERR:         "句子不存在",
	SENTENCE_NOTEXIST_COLLECT_ERR: "句子不存在，收藏失败，请联系管理员处理",
	SENTENCE_UPDATE_ERR:           "句子信息更新失败",
	SENTENCE_COLLECT_ERR:          "句子收藏失败",
	SENTENCE_COLLECT_ISEXIST_ERR:  "您已收藏该文章，请勿重复操作",
	SENTENCE_CANCEL_COLLECT_ERR:   "句子取消收藏失败",
	SENTENCEID_NULL_ERR:           "句子id为空，请重新获取信息",
	SENTEN_COLLECT_LIST_ERR:       "句子收藏列表获取失败",

	LISTEN_EXIST_ERR:                   "听力已存在",
	DELETE_LISTEN_ID_ERR:               "听力id无效，请重新获取听力",
	DELETE_LISTEN_ERR:                  "听力删除失败",
	LISTEN_CANCEL_COLLECT_ERR:          "听力取消收藏失败",
	LISTEN_NOTEXIST_ERR:                "听力不存在",
	LISTEN_UPDATE_ERR:                  "听力信息更新失败",
	LISTEN_TYPE_ERR:                    "听力类型错误",
	LISTENLIST_GET_ERR:                 "听力列表获取失败，请重新刷新",
	LISTEN_CACHE_ERR:                   "听力缓存设置失败",
	LISTEN_CACHE_GET_ERR:               "听力缓存获取失败",
	LISTEN_NOTEXIST_COLLECT_ERR:        "听力不存在，收藏失败，请联系管理员处理",
	LISTEN_COLLECT_ERR:                 "听力收藏失败",
	LISTEN_COLLECT_ISEXIST_ERR:         "您已收藏该听力，请勿重复操作",
	LISTEN_NOTEXIST_CANCEL_COLLECT_ERR: "听力不存在，取消收藏失败，请联系管理员处理",
	LISTEN_COLLECT_LIST_ERR:            "听力收藏列表获取失败",
	LISTENLIST_NULL_ERR:                "听力资源列表加载失败",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}

func GetError(code int) error {
	return errors.New(codeMsg[code])
}

func GetStrAndError(str string, code int) error {
	return errors.New(str + codeMsg[code])
}

func GetErrorAndError(str string, code int) error {
	return errors.New(codeMsg[code] + str)
}
