//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/huynhbaoking112/System_design_Golang/internal/controller"
	"github.com/huynhbaoking112/System_design_Golang/internal/repo"
	"github.com/huynhbaoking112/System_design_Golang/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
