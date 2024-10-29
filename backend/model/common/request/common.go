package request

type PageQuery struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

type IdQuery struct {
	Id int `json:"id" form:"id"`
}
