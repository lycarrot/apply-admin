package email

import (
	"github.com/gin-gonic/gin"
)

type emailPlugin struct {
}

func CreateEmailPlug() *emailPlugin {

	return &emailPlugin{}

}

func (*emailPlugin) Register(group *gin.Engine) {
	//router.RouterGroupApp
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
