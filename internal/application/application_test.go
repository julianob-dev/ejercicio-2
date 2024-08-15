package application

import (
	"testing"

	"Github.com/julianob-dev/ejercicio-2/internal/repository/memory"
)

func TestNewStruct(t *testing.T) {
	//Arrange
	repo, err := memory.New()
	
	if err != nil {
		t.Errorf("Error creating repository")
	}

	app := New(repo)
	
	if app == nil {
		t.Errorf("application fail to be created")
	}

	sector, err := app.GetById("COMUNICACION")

	if err != nil {
		t.Errorf("error al obtener el sector")
	}

	if sector == nil {
		t.Errorf("sector vacio")
	}
}