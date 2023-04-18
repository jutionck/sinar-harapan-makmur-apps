package usecase

import (
	"errors"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/mock/mock_repo"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"time"
)

var brandDummies = []model.Brand{
	{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "Honda",
	},
	{
		BaseModel: model.BaseModel{
			ID:        "2",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "Toyota",
	},
	{
		BaseModel: model.BaseModel{
			ID:        "3",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "BMW",
	},
}

type BrandUseCaseTestSuite struct {
	suite.Suite
	repo *mock_repo.BrandRepoMock
}

func (suite *BrandUseCaseTestSuite) TestFindById_Success() {
	brandDm := brandDummies[0]
	suite.repo.On("Get", "1").Return(&brandDm, nil)
	useCase := NewBrandUseCase(suite.repo)
	brand, err := useCase.FindById("1")
	assert.Equal(suite.T(), brandDm, *brand)
	assert.Nil(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestFindById_Fail() {
	suite.repo.On("Get", "1").Return(nil, errors.New("repo error"))
	useCase := NewBrandUseCase(suite.repo)
	brand, err := useCase.FindById("1")
	assert.Nil(suite.T(), brand)
	assert.Error(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestFindAll_Success() {
	suite.repo.On("List").Return(brandDummies, nil)
	useCase := NewBrandUseCase(suite.repo)
	brands, err := useCase.FindAll()
	assert.Equal(suite.T(), brandDummies, brands)
	assert.Nil(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestFindAll_Fail() {
	suite.repo.On("List").Return(nil, errors.New("repo error"))
	useCase := NewBrandUseCase(suite.repo)
	brands, err := useCase.FindAll()
	assert.Nil(suite.T(), brands)
	assert.Error(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestSearchBy_Success() {
	filter := map[string]interface{}{"brand": "Honda"}
	suite.repo.On("Search", filter).Return(brandDummies, nil)
	useCase := NewBrandUseCase(suite.repo)
	brands, err := useCase.SearchBy(filter)
	assert.Equal(suite.T(), brandDummies, brands)
	assert.Nil(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestSearchBy_Fail() {
	filter := map[string]interface{}{"brand": "Honda"}
	suite.repo.On("Search", filter).Return(nil, errors.New("repo error"))
	useCase := NewBrandUseCase(suite.repo)
	brands, err := useCase.SearchBy(filter)
	assert.Nil(suite.T(), brands)
	assert.Error(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestDeleteData_Success() {
	brandDm := brandDummies[0]
	suite.repo.On("Get", "1").Return(&brandDm, nil)
	suite.repo.On("Delete", "1").Return(nil)
	useCase := NewBrandUseCase(suite.repo)
	err := useCase.DeleteData("1")
	assert.Nil(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestDeleteData_Fail() {
	suite.repo.On("Get", "1").Return(nil, errors.New("repo error"))
	suite.repo.On("Delete", "1").Return(errors.New("repo error"))
	useCase := NewBrandUseCase(suite.repo)
	err := useCase.DeleteData("1")
	assert.Error(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestSaveData_Success() {
	brandDm := brandDummies[0]
	suite.repo.On("CountData", "Honda", "1").Return(nil)
	suite.repo.On("Get", "1").Return(&brandDm, nil)
	suite.repo.On("Save", &brandDm).Return(nil)
	useCase := NewBrandUseCase(suite.repo)
	err := useCase.SaveData(&brandDm)
	assert.Nil(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestSaveData_Fail() {
	brandDm := brandDummies[0]
	suite.repo.On("CountData", "Honda", "1").Return(errors.New("count data repo error"))
	suite.repo.On("Save", &brandDm).Return(errors.New("repo error"))
	// error check validate before save
	brandDm.Name = ""
	useCase := NewBrandUseCase(suite.repo)
	err := useCase.SaveData(&brandDm)
	assert.Error(suite.T(), err)
	// error check data exists
	brandDm = brandDummies[0]
	err = useCase.SaveData(&brandDm)
	assert.Error(suite.T(), err)
}

func (suite *BrandUseCaseTestSuite) TestSaveDataFindById_Fail() {
	brandDm := brandDummies[0]
	suite.repo.On("CountData", "Honda", "1").Return(nil)
	suite.repo.On("Get", "1").Return(nil, errors.New("not found"))
	useCase := NewBrandUseCase(suite.repo)
	err := useCase.SaveData(&brandDm)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "brand with ID 1 not found", err.Error())
}

func (suite *BrandUseCaseTestSuite) SetupTest() {
	suite.repo = new(mock_repo.BrandRepoMock)
}

func TestBrandUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BrandUseCaseTestSuite))
}
