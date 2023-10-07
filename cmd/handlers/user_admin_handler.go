package handlers

import (
	"github.com/dipay/controller"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
)

func (h *MyHandler) UserAdminHandler() controller.IUserAdminController {
	modelUserAdmin := model.NewUserAdmin()
	userAdminRepository := repositories.NewUserAdminRepository(h.Application.MongoDBClient, modelUserAdmin)
	userAdminUseCase := usecase.NewUseCaseUserAdmin(userAdminRepository, modelUserAdmin, h.Application.JWT)
	userAdminController := controller.NewUserAdminController(userAdminUseCase, h.Application.ENV.ContextTimeOut, h.Application.Validator)
	return userAdminController
}
