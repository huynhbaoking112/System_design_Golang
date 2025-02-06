package routers

import (
	"github.com/huynhbaoking112/System_design_Golang/internal/routers/manager"
	"github.com/huynhbaoking112/System_design_Golang/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
