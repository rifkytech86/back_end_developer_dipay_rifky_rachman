package handlers

import (
	"github.com/dipay/controller"
)

func (h *MyHandler) DuplicateZeroHandler() controller.IDuplicateZeroController {

	duplicateZeroController := controller.NewDuplicateZeroController(h.Application.ENV.ContextTimeOut, h.Application.Validator)
	return duplicateZeroController
}
