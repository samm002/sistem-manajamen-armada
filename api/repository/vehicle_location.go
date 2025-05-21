package repository

import (
	"errors"
	"sistem-manajemen-armada/database/model"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Create(vehicleLocation *model.VehicleLocation) error
	FindAll() ([]*model.VehicleLocation, error)
	FindById(vehicleId string) (*model.VehicleLocation, error)
	Update(vehicleId string, vehicleLocation map[string]interface{}) error
	Delete(vehicleId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(vehicleLocation *model.VehicleLocation) error {
	err := r.db.Create(vehicleLocation).Error

	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return errors.New("vehicle location already exist with inputted timestamp")
		}
	}

	return err
}

func (r *repository) FindAll() ([]*model.VehicleLocation, error) {
	var vehicleLocations []*model.VehicleLocation

	if err := r.db.Find(&vehicleLocations).Error; err != nil {
		return nil, err
	}

	return vehicleLocations, nil
}

func (r *repository) FindById(vehicleId string) (*model.VehicleLocation, error) {
	var vehicleLocation model.VehicleLocation

	if err := r.db.Where("vehicle_id = ?", vehicleId).First(&vehicleLocation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("vehicle not found")
		}

		return nil, err
	}

	return &vehicleLocation, nil
}

func (r *repository) Update(vehicleId string, vehicleLocation map[string]interface{}) error {
	_, err := r.FindById(vehicleId)

	if err != nil {
		return err
	}

	return r.db.Model(&model.VehicleLocation{VehicleId: vehicleId}).Updates(vehicleLocation).Error
}

func (r *repository) Delete(vehicleId string) error {
	_, err := r.FindById(vehicleId)

	if err != nil {
		return err
	}

	return r.db.Delete(&model.VehicleLocation{}, vehicleId).Error
}
