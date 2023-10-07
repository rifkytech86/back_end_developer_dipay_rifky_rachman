package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type employeeUserAdmin struct {
	EmployeeRepository repositories.IEmployeeRepository
	CompanyRepository  repositories.ICompanyRepository
}

//go:generate mockery --name IEmployeeUseCase
type IEmployeeUseCase interface {
	AddEmployee(context.Context, string, *api.AddEmployeeJSONBody) (string, string, error)
	GetEmployeeByID(context.Context, string) (*model.Employees, error)
	GetEmployeeByCompanyID(context.Context, string) ([]*model.Employees, error)
	UpdateEmployeeData(ctx context.Context, companyID string, employeeID string, req *api.UpdateEmployeeDataJSONRequestBody) (string, string, error)
	DeleteEmployeeByID(context.Context, string) error
}

func NewEmployeeUseCase(employeeRepository repositories.IEmployeeRepository, companyRepository repositories.ICompanyRepository) IEmployeeUseCase {
	return &employeeUserAdmin{
		EmployeeRepository: employeeRepository,
		CompanyRepository:  companyRepository,
	}
}

func (e *employeeUserAdmin) AddEmployee(ctx context.Context, companyId string, req *api.AddEmployeeJSONBody) (string, string, error) {
	// check if company exist
	var company model.Companies
	companyIDHex, err := primitive.ObjectIDFromHex(companyId)
	if err != nil {
		return "", "", errors.New(internal.ErrorInvalidRequest.String())
	}

	err = e.CompanyRepository.FetchOne(ctx, bson.M{"_id": companyIDHex}, &company)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", "", errors.New(internal.ErrDataCompanyNotFound.String())
		}
		return "", "", errors.New(internal.ErrorInternalServer.String())
	}

	dataEmployeeInsert := model.Employees{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		JobTitle:    strings.ReplaceAll(strings.ToLower(req.Jobtitle), " ", ""),
		CompanyID:   companyId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	lastInsertedID, err := e.EmployeeRepository.Create(ctx, dataEmployeeInsert)
	if err != nil {
		var writeException mongo.WriteException
		if errors.As(err, &writeException) {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 {
					return "", "", errors.New(internal.ErrorInvalidDuplicateCompanyAndEmail.String())
				}
			}
		}
		return "", "", errors.New(internal.ErrorInternalServer.String())
	}

	// NOte this is only for make sure data correctly, for Assignment test only
	var employee model.Employees
	err = e.EmployeeRepository.FetchOne(ctx, bson.M{"_id": lastInsertedID}, &employee)
	if err != nil {
		return "", "", errors.New(internal.ErrorInternalServer.String())
	}

	// send email
	err = e.EmployeeRepository.SendEmail(ctx, employee.Email)
	if err != nil {
		fmt.Println("warning send email failed")
	}

	return employee.ID.Hex(), employee.CompanyID, nil
}

func (e *employeeUserAdmin) GetEmployeeByID(ctx context.Context, employeeID string) (*model.Employees, error) {
	employeeIDHex, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return nil, errors.New(internal.ErrorInvalidParameterID.String())
	}

	var employee model.Employees
	err = e.EmployeeRepository.FetchOne(ctx, bson.M{"_id": employeeIDHex}, &employee)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New(internal.ErrorInvalidDataNotFound.String())
		}
		return nil, errors.New(internal.ErrorInternalServer.String())
	}

	return &employee, nil

}

func (e *employeeUserAdmin) GetEmployeeByCompanyID(ctx context.Context, companyID string) ([]*model.Employees, error) {
	listEmployee, err := e.EmployeeRepository.Fetch(ctx, bson.M{"company_id": companyID})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New(internal.ErrorInvalidDataNotFound.String())
		}
		return nil, errors.New(internal.ErrorInternalServer.String())
	}

	return listEmployee, nil

}

func (e *employeeUserAdmin) UpdateEmployeeData(ctx context.Context, companyID string, employeeID string, req *api.UpdateEmployeeDataJSONRequestBody) (string, string, error) {
	jobTittle := strings.ReplaceAll(strings.ToLower(req.Jobtitle), " ", "")

	update := bson.M{"$set": bson.M{"name": req.Name, "phone_number": req.PhoneNumber, "jobtitle": jobTittle, "email": req.Email, "updated_at": time.Now()}}
	employeeIDHex, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return "", "", errors.New(internal.ErrorInvalidParameterID.String())
	}
	err = e.EmployeeRepository.Update(ctx, bson.M{"company_id": companyID, "_id": employeeIDHex}, update)
	if err != nil {
		if err.Error() == internal.ErrNoModifyUpdate.String() {
			return "", "", errors.New(internal.ErrNoModifyUpdate.String())
		}
		var writeException mongo.WriteException
		if errors.As(err, &writeException) {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 {
					return "", "", errors.New(internal.ErrorInvalidDuplicateCompanyAndEmail.String())
				}
			}
		}

		return "", "", errors.New(internal.ErrorInternalServer.String())
	}

	var employee model.Employees
	err = e.EmployeeRepository.FetchOne(ctx, bson.M{"company_id": companyID, "_id": employeeIDHex}, &employee)
	if err != nil {
		return "", "", errors.New(internal.ErrorInternalServer.String())
	}
	return employee.ID.Hex(), employee.CompanyID, nil

}

func (e *employeeUserAdmin) DeleteEmployeeByID(ctx context.Context, employeeID string) error {
	employeeIDHex, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return errors.New(internal.ErrorInvalidParameterID.String())
	}

	err = e.EmployeeRepository.Delete(ctx, bson.M{"_id": employeeIDHex})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New(internal.ErrorInvalidDataNotFound.String())
		}
		return errors.New(internal.ErrorInternalServer.String())
	}
	return nil

}
