package main

import (
	"encoding/csv"
	"fmt"
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

	for _, row := range lines[1:] {
		if v, ok := sectorMap[row[0]]; ok {
			esp := Especialidad{
				Id: 1,
				Nombre: row[1],
			}
			v.Especialidades = append(v.Especialidades, esp)
			continue
		}
		sector := Sector{
			Id: 1,
			Nombre: row[0],
			Especialidades: []Especialidad{},
		}

		sectorMap[row[0]] = &sector
	}

	fmt.Println(sectorMap)
	fmt.Println(sectorMap["ALMACENES"])
}