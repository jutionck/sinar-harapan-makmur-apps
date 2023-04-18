package mock_repo

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/stretchr/testify/mock"
)

type BrandRepoMock struct {
	mock.Mock
}

func (b *BrandRepoMock) Get(id string) (*model.Brand, error) {
	// membuat seolah memanggil method asli (Get) pada brandRepository
	args := b.Called(id)
	// Cek error dengan memasukkan index 1 (error)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	// Jika sukses balikan index ke 0 (*model.Brand) (casting dulu ya karena interface{})
	return args.Get(0).(*model.Brand), nil
}

func (b *BrandRepoMock) List() ([]model.Brand, error) {
	args := b.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Brand), nil
}

func (b *BrandRepoMock) Save(payload *model.Brand) error {
	args := b.Called(payload)
	return args.Error(0)
}

func (b *BrandRepoMock) Delete(id string) error {
	args := b.Called(id)
	return args.Error(0)
}
func (b *BrandRepoMock) Search(by map[string]interface{}) ([]model.Brand, error) {
	args := b.Called(by)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Brand), nil
}

func (b *BrandRepoMock) CountData(fieldName string, id string) error {
	args := b.Called(fieldName, id)
	return args.Error(0)
}

func (b *BrandRepoMock) Paging(requestQueryParams dto.RequestQueryParams) ([]model.Brand, dto.Paging, error) {
	args := b.Called(requestQueryParams)
	return args.Get(0).([]model.Brand), args.Get(1).(dto.Paging), args.Error(2)
}
