package handlers

import (
	"github.com/dipay/bootstrap"
	"github.com/dipay/internal/db/mocks"
	"github.com/dipay/internal/env"
	"testing"
)

func TestMyHandler_CompanyHandler(t *testing.T) {
	mockMongo := new(mocks.Database)
	mockApp := bootstrap.Application{
		MongoDBClient: mockMongo,
		ENV:           &env.ENV{},
	}
	h := &MyHandler{
		Application: mockApp,
	}
	companyHandler := h.CompanyHandler()
	if companyHandler == nil {
		t.Errorf("Expected non-nil value, but got nil")
	}
}
