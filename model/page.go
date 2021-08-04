package model

// page 分页
type Page struct {
	PageNum     int `json:"page_num"`//开始的页码
	PageSize    int  `json:"page_size"`//每页显示多少
	Query   	string  `json:"query_info"`//查询的数据
}


