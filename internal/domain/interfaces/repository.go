package interfaces

import "Github.com/julianob-dev/ejercicio-2/pkg/model"

type Repository interface {
	GetAll() map[string]*model.Sector
}
