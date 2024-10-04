package request

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type CasbinInReceive struct {
	AuthorityId uint         `json:"authorityId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/auth/login", Method: "POST"},
		{Path: "/menu/getMenu", Method: "POST"},
		{Path: "/api/create", Method: "POST"},
		//{Path: "/jwt/jsonInBlacklist", Method: "POST"},
		//{Path: "/user/changePassword", Method: "POST"},
		//{Path: "/user/setUserAuthority", Method: "POST"},
		{Path: "/user/getUserInfo", Method: "GET"},
		//{Path: "/user/setSelfInfo", Method: "PUT"},
		//{Path: "/fileUploadAndDownload/upload", Method: "POST"},
		//{Path: "/sysDictionary/findSysDictionary", Method: "GET"},
	}
}
