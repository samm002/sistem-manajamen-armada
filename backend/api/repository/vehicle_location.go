package repository

import (
	"errors"
	"fmt"
	"sistem-manajemen-armada/database/model"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Create(vehicleLocation *model.VehicleLocation) error
	FindAll() ([]*model.VehicleLocation, error)
	FindHistory(vehicleId string, start *int, end *int) ([]*model.VehicleLocation, error)
	FindLatestLocationById(vehicleId string) (*model.VehicleLocation, error)

	// Tidak dipakai
	Update(vehicleId string, vehicleLocation map[string]interface{}) error

	// Tidak dipakai
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

	if err := r.db.Order("vehicle_id ASC, timestamp DESC").Find(&vehicleLocations).Error; err != nil {
		return nil, err
	}

	return vehicleLocations, nil
}

func (r *repository) FindHistory(vehicleId string, start *int, end *int) ([]*model.VehicleLocation, error) {
	var vehicleLocations []*model.VehicleLocation

	query := r.db.Where("vehicle_id = ?", vehicleId)

	if start != nil {
		query = query.Where("timestamp >= ?", *start)
	}

	if end != nil {
		query = query.Where("timestamp <= ?", *end)
	}

	if err := query.Find(&vehicleLocations).Error; err != nil {
		return nil, err
	}

	if len(vehicleLocations) == 0 {
		return nil, fmt.Errorf("vehicle with id %s not found", vehicleId)
	}

	return vehicleLocations, nil
}

func (r *repository) FindLatestLocationById(vehicleId string) (*model.VehicleLocation, error) {
	var vehicleLocation model.VehicleLocation

	if err := r.db.Where("vehicle_id = ?", vehicleId).Order("timestamp DESC").First(&vehicleLocation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("vehicle with id %s not found", vehicleId)
		}

		return nil, err
	}

	return &vehicleLocation, nil
}

// Tidak dipakai
func (r *repository) Update(vehicleId string, vehicleLocation map[string]interface{}) error {
	return r.db.Model(&model.VehicleLocation{VehicleId: vehicleId}).Updates(vehicleLocation).Error
}

// Tidak dipakai
func (r *repository) Delete(vehicleId string) error {
	return r.db.Delete(&model.VehicleLocation{}, vehicleId).Error
}
