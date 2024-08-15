package model

type Sector struct {
	Id             int
	Nombre         string
	Especialidades []Especialidad
}

type Especialidad struct {
	Id     int
	Nombre string
}