package repository

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"gorm.io/gorm"
)

type BrandRepository interface {
	BaseRepository[model.Brand]
	BaseRepositoryCount[model.Brand]
	ListBrandWithVehicle() ([]model.Brand, error)
	GetBrandWithVehicle(brandId string) (*model.Brand, error)
}

type brandRepository struct {
	db *gorm.DB
}

func (b *brandRepository) Delete(id string) error {
	return b.db.Delete(&model.Brand{}, "id=?", id).Error
}

func (b *brandRepository) Get(id string) (*model.Brand, error) {
	var brand model.Brand
	result := b.db.First(&brand, "id=?", id).Error
	if result != nil {
		return nil, result
	}
	return &brand, nil
}

func (b *brandRepository) List() ([]model.Brand, error) {
	var brands []model.Brand
	result := b.db.Find(&brands).Error
	if result != nil {
		return nil, result
	}
	return brands, nil
}

func (b *brandRepository) Save(payload *model.Brand) error {
	return b.db.Save(payload).Error
}

func (b *brandRepository) Search(by map[string]interface{}) ([]model.Brand, error) {
	var brands []model.Brand
	result := b.db.Where(by).Find(&brands).Error
	if result != nil {
		return nil, result
	}
	return brands, nil
}

func (b *brandRepository) GetBrandWithVehicle(brandId string) (*model.Brand, error) {
	var brand model.Brand
	result := b.db.Preload("Brand").First(&brand, "id=?", brandId).Error
	if result != nil {
		return nil, result
	}
	return &brand, nil
}

func (b *brandRepository) ListBrandWithVehicle() ([]model.Brand, error) {
	var brands []model.Brand
	result := b.db.Preload("Brand").Find(&brands).Error
	if result != nil {
		return nil, result
	}
	return brands, nil
}

func (b *brandRepository) CountData(fieldName string, id string) error {
	var count int64
	var result *gorm.DB
	if id != "" {
		result = b.db.Unscoped().Model(&model.Brand{}).Where("name ILIKE ? AND id <> ?", "%"+fieldName+"%", id).Count(&count)
	} else {
		result = b.db.Unscoped().Model(&model.Brand{}).Where("name ILIKE ?", "%"+fieldName+"%").Count(&count)
	}

	if result.Error != nil {
		return result.Error
	}

	if count > 0 {
		return fmt.Errorf("field with name %s already exists", fieldName)
	}
	return nil
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}
