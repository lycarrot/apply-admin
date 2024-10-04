package utils

var (
	RegisterVerify  = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	AuthorityVerify = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	MenuVerify      = Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify  = Rules{"Title": {NotEmpty()}}
	ApiVerify       = Rules{"Path": {NotEmpty()}, "Group": {NotEmpty()}, "Description": {NotEmpty()}, "Method": {NotEmpty()}}
	CasbinVerify    = Rules{"AuthorityId": {NotEmpty()}}
)
