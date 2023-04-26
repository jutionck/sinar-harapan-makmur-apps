package mock_usecase

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/stretchr/testify/mock"
)

type BrandUseCaseMock struct {
	mock.Mock
}

func (b *BrandUseCaseMock) FindById(id string) (*model.Brand, error) {
	args := b.Called(id)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Brand), nil
}

func (b *BrandUseCaseMock) FindAll() ([]model.Brand, error) {
	args := b.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Brand), nil
}

func (b *BrandUseCaseMock) SaveData(payload *model.Brand) error {
	args := b.Called(payload)
	return args.Error(0)
}

func (b *BrandUseCaseMock) DeleteData(id string) error {
	args := b.Called(id)
	return args.Error(0)
}
func (b *BrandUseCaseMock) SearchBy(by map[string]interface{}) ([]model.Brand, error) {
	args := b.Called(by)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Brand), nil
}

func (b *BrandUseCaseMock) Pagination(requestQueryParams dto.RequestQueryParams) ([]model.Brand, dto.Paging, error) {
	args := b.Called(requestQueryParams)
	return args.Get(0).([]model.Brand), args.Get(1).(dto.Paging), args.Error(2)
}
