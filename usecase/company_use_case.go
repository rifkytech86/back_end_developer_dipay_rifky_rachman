package usecase

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type companyUseCase struct {
	CompanyRepository repositories.ICompanyRepository
	JWT               jwt.IJWTRSAToken
}

type ICompanyUseCase interface {
	AddCompany(context.Context, *api.AddCompanyJSONBody) (string, error)
	GetCompany(context.Context) ([]*model.Companies, error)
	UpdateCompanyStatusActive(ctx context.Context, id string) (idAffected string, isActive bool, err error)
}

func NewCompanyUseCase(companyRepository repositories.ICompanyRepository, jwt jwt.IJWTRSAToken) ICompanyUseCase {
	return &companyUseCase{
		CompanyRepository: companyRepository,
		JWT:               jwt,
	}
}

func (u *companyUseCase) AddCompany(ctx context.Context, req *api.AddCompanyJSONBody) (string, error) {
	companies := model.Companies{
		CompanyName:     req.CompanyName,
		TelephoneNumber: req.TelephoneNumber,
		IsActive:        false,
		Address:         req.Address,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	lastInsertId, err := u.CompanyRepository.Create(ctx, companies)
	if err != nil {
		return "", errors.New(internal.ErrorInternalServer.String())
	}

	return lastInsertId.Hex(), nil
}

func (u *companyUseCase) GetCompany(ctx context.Context) ([]*model.Companies, error) {
	listCompanies, err := u.CompanyRepository.Fetch(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []*model.Companies{}, nil
		}
		return nil, errors.New(internal.ErrorInternalServer.String())
	}

	return listCompanies, nil
}

func (u *companyUseCase) UpdateCompanyStatusActive(ctx context.Context, id string) (idAffected string, isActive bool, err error) {
	update := bson.M{"$set": bson.M{"is_active": true, "updated_at": time.Now()}}

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", false, errors.New(internal.ErrorInvalidParameterID.String())
	}

	err = u.CompanyRepository.Update(ctx, bson.M{"_id": idHex}, update)
	if err != nil {
		if err.Error() == internal.ErrNoModifyUpdate.String() {
			return "", false, errors.New(internal.ErrNoModifyUpdate.String())
		}
		return "", false, errors.New(internal.ErrorInternalServer.String())
	}

	var company model.Companies
	err = u.CompanyRepository.FetchOne(ctx, bson.M{"_id": idHex}, &company)
	return company.ID.Hex(), company.IsActive, nil
}
