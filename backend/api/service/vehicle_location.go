package service

import (
	"log"
	"sistem-manajemen-armada/api/common/util"
	"sistem-manajemen-armada/api/dto"
	"sistem-manajemen-armada/api/repository"
	"sistem-manajemen-armada/common/constant"
	"sistem-manajemen-armada/pkg/rabbitmq"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(payload *dto.CreateVehicleLocationDto) (*dto.VehicleLocationDto, error)
	FindAll() ([]*dto.VehicleLocationDto, error)
	FindHistory(vehicleId string, start *int, end *int) ([]*dto.VehicleLocationDto, error)
	FindLatestLocationById(vehicleId string) (*dto.VehicleLocationDto, error)

	// Tidak dipakai
	Update(vehicleId string, payload *dto.UpdateVehicleLocationDto) (*map[string]interface{}, error)

	// Tidak dipakai
	Delete(vehicleId string) error
}

type service struct {
	repository repository.Repository
	validator  *validator.Validate
}

func NewService(repository repository.Repository, validator *validator.Validate) Service {
	return &service{repository, validator}
}

func (s *service) Create(payload *dto.CreateVehicleLocationDto) (*dto.VehicleLocationDto, error) {
	vehicleLocation := payload.ToModel()

	if err := s.repository.Create(&vehicleLocation); err != nil {
		return nil, err
	}

	response := dto.ToResponse(&vehicleLocation)

	distance := util.CalculateCoordinateDistance(
		payload.Latitude, payload.Longitude,
		constant.GeofenceLatitude, constant.GeofenceLongitude,
	)

	if distance <= 50 {
		encodedPayload, err := util.GenerateGeofenceEventMessage(payload, s.validator)

		if err != nil {
			log.Fatal(err)
		}

		rabbitmq.Publish(*encodedPayload)

	}

	return &response, nil
}

func (s *service) FindAll() ([]*dto.VehicleLocationDto, error) {
	vehicleLocations, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}

	var response []*dto.VehicleLocationDto

	for _, v := range vehicleLocations {
		vehicleLocation := dto.ToResponse(v)
		response = append(response, &vehicleLocation)
	}

	return response, nil
}

func (s *service) FindHistory(vehicleId string, start *int, end *int) ([]*dto.VehicleLocationDto, error) {
	vehicleLocations, err := s.repository.FindHistory(vehicleId, start, end)

	if err != nil {
		return nil, err
	}

	var response []*dto.VehicleLocationDto

	for _, v := range vehicleLocations {
		vehicleLocation := dto.ToResponse(v)
		response = append(response, &vehicleLocation)
	}

	return response, nil
}

func (s *service) FindLatestLocationById(vehicleId string) (*dto.VehicleLocationDto, error) {
	vehicleLocation, err := s.repository.FindLatestLocationById(vehicleId)

	if err != nil {
		return nil, err
	}

	response := dto.ToResponse(vehicleLocation)

	return &response, nil
}

// Tidak dipakai
func (s *service) Update(vehicleId string, payload *dto.UpdateVehicleLocationDto) (*map[string]interface{}, error) {
	sanitizedPayload := payload.ConstructUpdatePayload()

	if err := s.repository.Update(vehicleId, sanitizedPayload); err != nil {
		return nil, err
	}

	return &sanitizedPayload, nil
}

// Tidak dipakai
func (s *service) Delete(vehicleId string) error {
	return s.repository.Delete(vehicleId)
}
