package models

type Room struct {
	Name string
	X    string
	Y    string
}

type Path struct {
	Rooms []Room
	Ants  []Ant
}
