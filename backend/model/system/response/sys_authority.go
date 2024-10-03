package response

import "gin-pro/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}
