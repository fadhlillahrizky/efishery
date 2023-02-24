package resources

import (
	"efishery-golang/repository"
)

type ServiceInterface interface {
	List() (data []ListResponse, err error)
}

type Service struct {
	repo repository.Resources
}

func NewService(repo repository.Resources) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) List() (data []ListResponse, err error) {
	getData, err := s.repo.GetResourceList()

	for _, item := range getData {
		data = append(data, ListResponse{
			UUID:         item.UUID,
			Komoditas:    item.Komoditas,
			AreaProvinsi: item.AreaProvinsi,
			AreaKota:     item.AreaKota,
			Size:         item.Size,
			Price:        item.Price,
			TglParsed:    item.TglParsed,
			PriceUSD:     item.Price,
		})
	}
	return data, err
}
