package usecase

import (
	"context"
	"github.com/dipay/repositories"
)

type countriesUseCase struct {
	CountriesRepository repositories.ICountriesRepository
}

type ICountriesUseCase interface {
	GetCountries(context.Context) ([]repositories.CountriesPayload, error)
}

func NewCountriesUseCase(countriesRepository repositories.ICountriesRepository) ICountriesUseCase {
	return &countriesUseCase{
		CountriesRepository: countriesRepository,
	}
}

func (co *countriesUseCase) GetCountries(ctx context.Context) ([]repositories.CountriesPayload, error) {
	listCountries, err := co.CountriesRepository.GetCountries(ctx)
	if err != nil {
		return nil, err
	}
	return listCountries, nil
}
