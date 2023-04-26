package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
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

type BrandRepoTestSuite struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	mockDb *gorm.DB
}

func (suite *BrandRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New() // sqlmock.New akan membuat sebuah mock dari *sql.DB
	assert.NoError(suite.T(), err)
	suite.mock = mock
	// karane menggunakan gorm.Db kita membuat simulasi open koneksi ke gorm
	dialect := postgres.New(postgres.Config{
		Conn: db,
	})
	// kemudian kita assign hasil dari open koneksi ke dalam suite.MockDb (*gorm.DB)
	suite.mockDb, _ = gorm.Open(dialect)
}

func TestBrandRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BrandRepoTestSuite))
}

func (suite *BrandRepoTestSuite) TestListBrand_Success() {
	brandRowDummies := make([]model.Brand, len(brandDummies))
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	for i, brand := range brandDummies {
		brandRowDummies[i] = brand
		rows.AddRow(brand.ID, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "mst_brand"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	expectedVehicleQuery := `SELECT * FROM "mst_vehicle" WHERE "mst_vehicle"."brand_id" IN ($1,$2,$3)`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedVehicleQuery)).WillReturnRows(rows)
	repo := NewBrandRepository(suite.mockDb)
	brands, err := repo.List()
	assert.Equal(suite.T(), brandDummies, brands)
	assert.NoError(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestListBrand_Fail() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	for _, brand := range brandDummies {
		rows.AddRow(brand.ID, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "mst_brand"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("db error"))
	repo := NewBrandRepository(suite.mockDb)
	brands, err := repo.List()
	assert.Nil(suite.T(), brands)
	assert.Error(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestGetBrand_Success() {
	brandRowDummies := brandDummies[0]
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	rows.AddRow(brandRowDummies.ID, brandRowDummies.Name, brandRowDummies.CreatedAt, brandRowDummies.UpdatedAt)
	// id='1'
	expectedQuery := `SELECT * FROM "mst_brand" WHERE id=$1 AND "mst_brand"."deleted_at" IS NULL`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(brandRowDummies.ID).WillReturnRows(rows)
	//expectedQuery := `SELECT (.+) FROM "mst_brand" WHERE id(\s*)=(\s*)\$1 AND "mst_brand"."deleted_at" IS NULL ORDER BY "mst_brand"."id" LIMIT 1`
	//suite.mock.ExpectQuery(expectedQuery).WithArgs(brandRowDummies.ID).WillReturnRows(rows)
	expectedVehicleQuery := `SELECT * FROM "mst_vehicle" WHERE "mst_vehicle"."brand_id" = $1`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedVehicleQuery)).WillReturnRows(rows)
	repo := NewBrandRepository(suite.mockDb)
	brand, err := repo.Get(brandRowDummies.ID)
	assert.Equal(suite.T(), brandRowDummies, *brand)
	assert.NoError(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestGetBrand_Fail() {
	brandRowDummies := &brandDummies[0]
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	rows.AddRow(brandRowDummies.ID, brandRowDummies.Name, brandRowDummies.CreatedAt, brandRowDummies.UpdatedAt)
	// id='1'
	expectedQuery := `SELECT * FROM "mst_brand" WHERE id=$1 AND "mst_brand"."deleted_at" IS NULL`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(brandRowDummies.ID).WillReturnError(errors.New("db error"))
	//expectedQuery := `SELECT (.+) FROM "mst_brand" WHERE id(\s*)=(\s*)\$1 AND "mst_brand"."deleted_at" IS NULL ORDER BY "mst_brand"."id" LIMIT 1`
	//suite.mock.ExpectQuery(expectedQuery).WithArgs(brandRowDummies.ID).WillReturnRows(rows)
	repo := NewBrandRepository(suite.mockDb)
	brand, err := repo.Get(brandRowDummies.ID)
	assert.Nil(suite.T(), brand)
	assert.Error(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestSearchBrand_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	for _, brand := range brandDummies {
		rows.AddRow(brand.ID, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	}
	//expectedQuery := `SELECT \* FROM "mst_brand" WHERE \"name\"(\s*)=(\s*)\$1 AND "mst_brand"."deleted_at" IS NULL`
	expectedQuery := `SELECT * FROM "mst_brand" WHERE "name" = $1 AND "mst_brand"."deleted_at" IS NULL`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("Honda").WillReturnRows(rows)
	expectedVehicleQuery := `SELECT * FROM "mst_vehicle" WHERE "mst_vehicle"."brand_id" IN ($1,$2,$3)`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedVehicleQuery)).WillReturnRows(rows)
	repo := NewBrandRepository(suite.mockDb)
	filter := map[string]interface{}{"name": "Honda"}
	brands, err := repo.Search(filter)
	assert.Equal(suite.T(), brandDummies, brands)
	assert.NoError(suite.T(), err)
}

func (suite *BrandRepoTestSuite) TestSearchBrand_Fail() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	for _, brand := range brandDummies {
		rows.AddRow(brand.ID, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	}
	//expectedQuery := `SELECT \* FROM "mst_brand" WHERE \"name\"(\s*)=(\s*)\$1 AND "mst_brand"."deleted_at" IS NULL`
	expectedQuery := `SELECT * FROM "mst_brand" WHERE "name" = $1 AND "mst_brand"."deleted_at" IS NULL`
	suite.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("Honda").WillReturnError(errors.New("db error"))
	repo := NewBrandRepository(suite.mockDb)
	filter := map[string]interface{}{"name": "Honda"}
	brands, err := repo.Search(filter)
	assert.Nil(suite.T(), brands)
	assert.Error(suite.T(), err)
}
