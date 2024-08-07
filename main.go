package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Sector struct {
	Id     int
	Nombre string
	Especialidades []Especialidad
}

type Especialidad struct {
	Id     int
	Nombre string
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



func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No hay suficientes parametros")
		os.Exit(1)
	}
	// 0, 1 .. n
	input := args[1]

	lines, err := readCSVFile(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sectorMap = make(map[string]*Sector)
	var sectorId int = 1
	var especialidadId int = 1
	for _, row := range lines[1:] {
		esp := Especialidad{
			Id: especialidadId,
			Nombre: row[1],
		}
		especialidadId++
		
		// Añado la especialidad al sector si lo encuentro y continuo con la siguiente linea
		if v, ok := sectorMap[row[0]]; ok {
			v.Especialidades = append(v.Especialidades, esp)
			continue
		}
		
		//si no existe creo el sector y lo agrego al mapa
		sector := Sector{
			Id: sectorId,
			Nombre: row[0],
			Especialidades: append([]Especialidad{}, esp),
		}

		sectorMap[row[0]] = &sector
		sectorId++
	}

	for _, v := range sectorMap {
		log.Printf("%v\n", v)
	}
}