package system

type ServiceGroup struct {
	JwtService
	MenuService
	InitDBService
	UserService
	AuthorityService
	CasbinService
	ApiService
}
