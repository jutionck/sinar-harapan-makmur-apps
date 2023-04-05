package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type VehicleUseCase interface {
	BaseUseCase[model.Vehicle]
	Paging(requestQueryParams dto.RequestQueryParams) ([]model.Vehicle, dto.Paging, error)
	UpdateVehicleStock(count int, id string) error
}

type vehicleUseCase struct {
	repo         repository.VehicleRepository
	brandUseCase BrandUseCase
}

func (v *vehicleUseCase) SearchBy(by map[string]interface{}) ([]model.Vehicle, error) {
	return v.repo.Search(by)
}

func (v *vehicleUseCase) FindAll() ([]model.Vehicle, error) {
	return v.repo.List()
}

func (v *vehicleUseCase) FindById(id string) (*model.Vehicle, error) {
	vehicle, err := v.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("vehicle with ID %s not found", id)
	}
	return vehicle, nil
}

func (v *vehicleUseCase) SaveData(payload *model.Vehicle) error {
	err := payload.Validate()
	if err != nil {
		return err
	}
	brand, err := v.brandUseCase.FindById(payload.BrandID)
	if err != nil {
		return fmt.Errorf("brand with ID %s not found", payload.ID)
	}
	payload.BrandID = brand.ID
	if payload.ID != "" {
		_, err := v.FindById(payload.ID)
		if err != nil {
			return fmt.Errorf("vehicle with ID %s not found", payload.ID)
		}
	}
	return v.repo.Save(payload)
}

func (v *vehicleUseCase) DeleteData(id string) error {
	vehicle, err := v.FindById(id)
	if err != nil {
		return fmt.Errorf("vehicle with ID %s not found", id)
	}
	return v.repo.Delete(vehicle.ID)
}

func (v *vehicleUseCase) UpdateVehicleStock(count int, id string) error {
	return v.repo.UpdateStock(count, id)
}

func (v *vehicleUseCase) Paging(requestQueryParams dto.RequestQueryParams) ([]model.Vehicle, dto.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, dto.Paging{}, fmt.Errorf("invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return v.repo.Paging(requestQueryParams)
}

func NewVehicleUseCase(repo repository.VehicleRepository, brandUseCase BrandUseCase) VehicleUseCase {
	return &vehicleUseCase{repo: repo, brandUseCase: brandUseCase}
}
