package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/mock/mock_usecase"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
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
		Name:     "Honda",
		Vehicles: []model.Vehicle{},
	},
	{
		BaseModel: model.BaseModel{
			ID:        "2",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:     "Toyota",
		Vehicles: []model.Vehicle{},
	},
	{
		BaseModel: model.BaseModel{
			ID:        "3",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:     "BMW",
		Vehicles: []model.Vehicle{},
	},
}

type BrandControllerTestSuite struct {
	suite.Suite
	useCaseMock *mock_usecase.BrandUseCaseMock
	routerMock  *gin.Engine
}

func (suite *BrandControllerTestSuite) SetupTest() {
	suite.useCaseMock = new(mock_usecase.BrandUseCaseMock)
	suite.routerMock = gin.Default()
}

func TestBrandControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BrandControllerTestSuite))
}

func (suite *BrandControllerTestSuite) TestGetHandler_Success() {
	brandDm := &brandDummies[0]
	// disimulasikan sebagai HTTP Response (status code, body dll)
	r := httptest.NewRecorder()
	// disimulasikan sebagai HTTP Requst (method, url, body payload)
	request, err := http.NewRequest(http.MethodGet, "/brands/1", nil)
	suite.useCaseMock.On("FindById", brandDm.ID).Return(brandDm, nil)
	NewBrandController(suite.routerMock, suite.useCaseMock)
	suite.routerMock.ServeHTTP(r, request)
	var brandActual struct {
		Status map[string]interface{}
		Data   model.Brand
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &brandActual)
	// expected 200, actual 200
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, 1)
	assert.Equal(suite.T(), brandDm.Name, brandActual.Data.Name)
}

func (suite *BrandControllerTestSuite) TestGetHandler_Fail() {
	brandDm := &brandDummies[0]
	// disimulasikan sebagai HTTP Response (status code, body dll)
	r := httptest.NewRecorder()
	// disimulasikan sebagai HTTP Requst (method, url, body payload)
	request, err := http.NewRequest(http.MethodGet, "/brands/1", nil)
	suite.useCaseMock.On("FindById", brandDm.ID).Return(nil, errors.New("failed"))
	NewBrandController(suite.routerMock, suite.useCaseMock)
	suite.routerMock.ServeHTTP(r, request)
	var errActual struct {
		Code        int
		Description string
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &errActual)
	assert.Equal(suite.T(), http.StatusNotFound, r.Code)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "failed", errActual.Description)
}

// CreateHandler -> Ketika sukses
// CreateHandler -> Ketika Payload yang dikirimkan kosong
// CreateHandler -> Ketika tetapi saat masuk ke usecase gagal

func (suite *BrandControllerTestSuite) TestCreateHandler_Success() {
	var payload = model.Brand{Name: "Honda"}
	suite.useCaseMock.On("SaveData", &payload).Return(nil)
	NewBrandController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "/brands", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	var brandActual struct {
		Status map[string]interface{}
		Data   model.Brand
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &brandActual)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), payload.Name, brandActual.Data.Name)
}

func (suite *BrandControllerTestSuite) TestCreateHandler_FailedBinding() {
	r := httptest.NewRecorder()
	NewBrandController(suite.routerMock, suite.useCaseMock)
	request, _ := http.NewRequest(http.MethodPost, "/brands", nil)
	suite.routerMock.ServeHTTP(r, request)
	var errActual struct {
		Code        int
		Description string
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &errActual)
	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.NotEmpty(suite.T(), errActual)
}

func (suite *BrandControllerTestSuite) TestCreateHandler_FailedUseCase() {
	var payload = model.Brand{Name: "Honda"}
	suite.useCaseMock.On("SaveData", &payload).Return(errors.New("failed"))
	NewBrandController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, "/brands", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	var errActual struct {
		Code        int
		Description string
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &errActual)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	assert.Equal(suite.T(), "failed", errActual.Description)
}

func (suite *BrandControllerTestSuite) TestUpdateHandler_Success() {
	var payload = model.Brand{BaseModel: model.BaseModel{ID: "1"}, Name: "Honda"}
	suite.useCaseMock.On("SaveData", &payload).Return(nil)
	NewBrandController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPut, "/brands", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	var brandActual struct {
		Status map[string]interface{}
		Data   model.Brand
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &brandActual)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), payload.Name, brandActual.Data.Name)
}

func (suite *BrandControllerTestSuite) TestUpdateHandler_FailedBinding() {
	r := httptest.NewRecorder()
	NewBrandController(suite.routerMock, suite.useCaseMock)
	request, _ := http.NewRequest(http.MethodPut, "/brands", nil)
	suite.routerMock.ServeHTTP(r, request)
	var errActual struct {
		Code        int
		Description string
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &errActual)
	assert.Equal(suite.T(), http.StatusBadRequest, r.Code)
	assert.NotEmpty(suite.T(), errActual)
}

func (suite *BrandControllerTestSuite) TestUpdateHandler_FailedUseCase() {
	var payload = model.Brand{BaseModel: model.BaseModel{ID: "1"}, Name: "Honda"}
	suite.useCaseMock.On("SaveData", &payload).Return(errors.New("failed"))
	NewBrandController(suite.routerMock, suite.useCaseMock)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPut, "/brands", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	var errActual struct {
		Code        int
		Description string
	}
	resp := r.Body.String()
	json.Unmarshal([]byte(resp), &errActual)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	assert.Equal(suite.T(), "failed", errActual.Description)
}
