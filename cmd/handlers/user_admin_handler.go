package handlers

import (
	"github.com/dipay/controller"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
)

func (h *MyHandler) UserAdminHandler() controller.IUserAdminController {
	table := model.NewUserAdmin()
	userAdminRepository := repositories.NewUserAdminRepository(h.Application.MongoDBClient, table)
	userAdminUseCase := usecase.NewUseCaseUserAdmin(userAdminRepository, h.Application.JWT, h.Application.ENV.SecretEndCrypt)
	userAdminController := controller.NewUserAdminController(userAdminUseCase, h.Application.ENV.ContextTimeOut, h.Application.Validator)
	return userAdminController
}
