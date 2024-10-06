package response

type PageResult struct {
	Lists    interface{} `json:"lists"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
