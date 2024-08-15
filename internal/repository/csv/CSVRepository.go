package csvRepository

import (
	"encoding/csv"
	"fmt"
	"os"

	"Github.com/julianob-dev/ejercicio-2/internal/domain/interfaces"
	"Github.com/julianob-dev/ejercicio-2/pkg/model"
)

type CSVRepository struct {
	data map[string]*model.Sector
}

func New() (interfaces.Repository, error) {
	lines, err := readCSVFile("sectores.csv")

	if err != nil {
		return nil, err
	}

	var sectorMap = make(map[string]*model.Sector)
	sectorSlice := []model.Sector{}
	var sectorId int = 1
	var especialidadId int = 1
	for _, row := range lines[1:] {
		esp := model.Especialidad{
			Id:     especialidadId,
			Nombre: row[1],
		}
		especialidadId++

		// Añado la especialidad al sector si lo encuentro y continuo con la siguiente linea
		if v, ok := sectorMap[row[0]]; ok {
			v.Especialidades = append(v.Especialidades, esp)
			continue
		}

		//si no existe creo el sector y lo agrego al mapa
		sector := model.Sector{
			Id:             sectorId,
			Nombre:         row[0],
			Especialidades: append([]model.Especialidad{}, esp),
		}

		sectorMap[row[0]] = &sector
		sectorSlice = append(sectorSlice, sector)
		sectorId++
	}

	return CSVRepository{
		data: sectorMap,
	}, nil
}


func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)

	if err != nil {
		fmt.Println("no se encuentra la ruta ", filepath)
		return nil, err
	}

	f, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()

	
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (r CSVRepository) GetAll() map[string]*model.Sector {
	return r.data
}