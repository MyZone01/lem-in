package model

type Room struct {
	Name string
	X    string
	Y    string
	ant Ant
}

type Path struct {
	Rooms []Room
	Ants  []Ant
}
