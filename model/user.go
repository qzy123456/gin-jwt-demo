package model


const (
	dbName     = "myBlog.db"
	userBucket = "user"
)

// User 用户类
type User struct {
	Id         string `json:"userId"`
	Name       string `json:"userName"`
	Gender     string `json:"gender"`
	Phone      string `json:"userMobile"`
	Pwd        string `json:"pwd"`
	Permission string `json:"permission"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}



// 反序列化


// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(phone string, pwd string) error {
	return nil
}
// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, User, error) {

	resultUser := User{}
	resultBool := true

	loginReq.Phone = resultUser.Phone
	loginReq.Pwd = resultUser.Pwd

	return resultBool, resultUser, nil
}

// EditUserReq 更新用户信息数据类
type EditUserReq struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	UserGender string `json:"gender"`
}


