package service

import (
	"sistem-manajemen-armada/api/dto"
	"sistem-manajemen-armada/api/repository"
)

type Service interface {
	Create(payload *dto.CreateVehicleLocationDto) (*dto.VehicleLocationDto, error)
	FindAll() ([]*dto.VehicleLocationDto, error)
	FindById(vehicleId string) (*dto.VehicleLocationDto, error)
	Update(vehicleId string, payload *dto.UpdateVehicleLocationDto) (*dto.VehicleLocationDto, error)
	Delete(vehicleId string) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{repository}
}

func (s *service) Create(payload *dto.CreateVehicleLocationDto) (*dto.VehicleLocationDto, error) {
	vehicleLocation := payload.ToModel()

	if err := s.repository.Create(&vehicleLocation); err != nil {
		return nil, err
	}

	response := dto.ToResponse(&vehicleLocation)

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

func (s *service) FindById(vehicleId string) (*dto.VehicleLocationDto, error) {
	vehicleLocation, err := s.repository.FindById(vehicleId)

	if err != nil {
		return nil, err
	}

	response := dto.ToResponse(vehicleLocation)

	return &response, nil
}

func (s *service) Update(vehicleId string, payload *dto.UpdateVehicleLocationDto) (*dto.VehicleLocationDto, error) {
	if _, err := s.repository.FindById(vehicleId); err != nil {
		return nil, err
	}

	sanitizedPayload := payload.ConstructUpdatePayload()

	if err := s.repository.Update(vehicleId, sanitizedPayload); err != nil {
		return nil, err
	}

	response, err := s.FindById(vehicleId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *service) Delete(vehicleId string) error {
	return s.repository.Delete(vehicleId)
}
