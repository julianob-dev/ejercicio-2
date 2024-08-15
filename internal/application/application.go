package application

import (
	"fmt"

	"Github.com/julianob-dev/ejercicio-2/internal/domain/interfaces"
	"Github.com/julianob-dev/ejercicio-2/pkg/model"
)

type Application struct {
	repo interfaces.Repository
}

func New(repo interfaces.Repository) *Application {
	return &Application{repo}
}

func (a *Application)GetById(key string) (*model.Sector, error)  {
	sectores := a.repo.GetAll()

	if s, ok :=sectores[key]; !ok {
		return nil, fmt.Errorf("no se encuentra el sector especificado")
	} else {
		return s, nil
	}
}