package repository

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	BaseRepository[model.Vehicle]
}

type vehicleRepository struct {
	db *gorm.DB
}

func (v *vehicleRepository) Search(by map[string]interface{}) ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := v.db.Where(by).Find(&vehicles)
	if err := result.Error; err != nil {
		return vehicles, err
	}
	return vehicles, nil
}

func (v *vehicleRepository) List() ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := v.db.Find(&vehicles)
	if err := result.Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (v *vehicleRepository) Get(id string) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	result := v.db.First(&vehicle, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (v *vehicleRepository) Save(payload *model.Vehicle) error {
	result := v.db.Save(payload)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (v *vehicleRepository) Delete(id string) error {
	return v.db.Delete(&model.Vehicle{}, "id=?", id).Error
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}
