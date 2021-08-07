package consts

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_USER:                 "已存在该用户",
	ERROR_DELETE_ERROR:              "删除失败",
	ERROR_UPDATE_ERROR:              "修改失败",
	ERROR_ADD_ROLE_ERROR:             "添加角色失败",
	ERROR_EXIST_MENU_URL:            "url地址存在",
}

/**
根据code返回相应的信息
@param	code	key
@return	msg		value
*/
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
