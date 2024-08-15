package memory

import (
	"Github.com/julianob-dev/ejercicio-2/internal/domain/interfaces"
	"Github.com/julianob-dev/ejercicio-2/pkg/model"
)

type MemoryRepository struct {
	data map[string]*model.Sector
}

func New() (interfaces.Repository, error) {
	var sectoresMap = make(map[string]*model.Sector)
	sector := model.Sector{
		Id:     2,
		Nombre: "COMUNICACIONES",
	}

	sectoresMap["COMUNICACIONES"] = &sector
	return MemoryRepository{
		data: sectoresMap,
	}, nil
}



func (r MemoryRepository) GetAll() map[string]*model.Sector {
	return r.data
}