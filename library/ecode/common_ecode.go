package ecode

// All common ecode
var (
	OK = add(0) // 正确

	SystemErr          = add(-1) // 内部错误
	NotMatchErr        = add(-2) // 参数校验失败（参数为空、失效等）
	SignCheckErr       = add(-3) // API校验密匙错误
	MethodNoPermission = add(-4) // 调用方对该Method没有权限
	ServiceUpdate      = add(-5) // 系统升级中

	NoLogin      = add(-101) // 账号未登录
	UserDisabled = add(-102) // 账号被封停
	UserNotExist = add(-103) // 用户不存在
	UserExpired  = add(-104) // 用户登陆身份已过期
	//LoginConflict            = add(-105) // 用户登录冲突
	LoginConflict                = add(10001) // 用户登录冲突
	LackOfItems                  = add(-106)  // 道具不足
	TaskLimitExceed              = add(-107)  // 已达到任务完成次数上限
	GoldLimitDelivered           = add(-108)  // 已达到金币奖励总数上限
	GoldLimitSingleDelivered     = add(-109)  // 超过单次金币奖励上限
	GoldAbnormalAccountDelivered = add(-110)  // 账号行为异常
	RPCError                     = add(-111)  // RPC 错误


	NotModified       = add(-304) // 木有改动
	TemporaryRedirect = add(-307) // 撞车跳转
	RequestErr        = add(-400) // 请求错误
	Unauthorized      = add(-401) // 未认证
	AccessDenied      = add(-403) // 访问权限不足
	NothingFound      = add(-404) // 啥都木有
	MethodNotAllowed  = add(-405) // 不支持该方法

	ServerErr          = add(-500) // 服务器错误
	ServiceUnavailable = add(-503) // 过载保护,服务暂不可用
	Deadline           = add(-504) // 服务调用超时
	LimitExceed        = add(-509) // 超出限制
	FileNotExists      = add(-616) // 上传文件不存在
	FileTooLarge       = add(-617) // 上传文件太大
	FailedTooManyTimes = add(-625) // 登录失败次数太多

	PasswordTooLeak       = add(-628) // 密码太弱
	UsernameOrPasswordErr = add(-629) // 用户名或密码错误
	TargetBlocked         = add(-643) // 被锁定
	UserLevelLow          = add(-650) // 用户等级太低
	UserDuplicate         = add(-652) // 重复的用户
	AccessTokenExpires    = add(-658) // Token 过期
	PasswordHashExpires   = add(-662) // 密码时间戳过期
	AreaLimit             = add(-688) // 地理区域限制

	LeagueTimeErr       = add(-701)   //联赛奖励时间错误
	LeagueUserErr       = add(-702)   //联赛用户信息错误

	
)
