package response

const (
	OK                  = "ok"
	UserExit            = "用户已存在"
	UserNoExit          = "用户不存在"
	PasswordError       = "密码错误"
	SystemError         = "服务器出错了"
	PhoneRegistered     = "该手机号已被注册"
	OldPasswordError    = "旧密码错误"
	PhoneNumberError    = "手机号必须为11位"
	MailTypeError       = "邮箱格式错误"
	PasswordNumberError = "密码至少6位"

	TokenExpired   = "token过期"
	Unauthorized   = "权限不足"
	AuthorityError = "权限设置错误"

	NameError  = "昵称不能为空"
	TitleError = "标题不能为空"
	CoverError = "请上传封面图片"
	DateError  = "日期错误"

	ReviewScoreExit = "您已点评过此电影"
	MovieNotExit    = "视频不存在"

	FailUploadFile = "文件上传失败"

	RequestError = "请求错误"

	CommentOrReplyError = "评论或回复内容不能为空"
	PageError           = "页码或请求数量错误"

	ParameterError = "参数错误"

	FailUploadImage = "图片上传失败"
	ImageTypeError  = "图片不符合要求"
	SaveImageError  = "图片保存失败"

	RequestTooMany = "请求数量过多"

	SearchNotEmpty = "搜索内容不能为空"

	ReviewScoreError = "打分范围为1-10分"

	MovieCoverError = "电影封面上传失败"

	FileTypeError = "文件格式错误"

	FileSaveError = "文件保存失败"

	FileSizeError = "文件大小不符合要求"

	CollectError = "禁止重复收藏"

	NoExitAdmin = "无该管理员或者权限不足"
)
