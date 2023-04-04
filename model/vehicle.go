package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	BrandID        string
	Model          string `gorm:"varchar;size:30"`
	ProductionYear int    `gorm:"size:4"`
	Color          string `gorm:"varchar;size:30"`
	IsAutomatic    bool
	Stock          int        `gorm:"check:stock >= 0"`
	SalePrice      int        `gorm:"check:sale_price > 0"`
	Status         string     `gorm:"check:status IN ('baru', 'bekas')"`
	Customers      []Customer `gorm:"many2many:customer_vehicles;"`
	BaseModel
}

func (Vehicle) TableName() string {
	return "mst_vehicle"
}

func (v *Vehicle) IsValidStatus() bool {
	return v.Status == "baru" || v.Status == "bekas"
}

func (v *Vehicle) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New().String()
	return nil
}
