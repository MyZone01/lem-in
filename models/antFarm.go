package models

type AntFarm struct {
	Start Room
	End   Room
	Rooms map[string]Room
	Links []Link
}
