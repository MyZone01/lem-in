package main

import (
	"lemin/model"
	"lemin/utils"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		wantAnts    int
		wantErr     bool
		wantAntFarm model.AntFarm
	}{
		{
			name:     "Test file with valid input",
			fileName: "./samples/testfile1",
			wantAnts: 10,
			wantErr:  false,
			wantAntFarm: model.AntFarm{
				Start: model.Room{Name: "A", X: "1", Y: "0"},
				End:   model.Room{Name: "E", X: "0", Y: "3"},
				Rooms: map[string]model.Room{
					"A": {Name: "A", X: "1", Y: "0"},
					"B": {Name: "B", X: "0", Y: "1"},
					"C": {Name: "C", X: "2", Y: "1"},
					"D": {Name: "D", X: "1", Y: "2"},
					"E": {Name: "E", X: "0", Y: "3"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "A", To: "C"},
					{From: "B", To: "D"},
					{From: "C", To: "D"},
					{From: "D", To: "E"},
				},
			},
		},
		{
			name:     "Test file with no ants",
			fileName: "./samples/testfile2",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "Test file with invalid data",
			fileName: "testfile3",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "Test file with missing start or end room",
			fileName: "./samples/testfile4",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "Test file with circular paths",
			fileName: "./samples/testfile5",
			wantAnts: 0,
			wantErr:  true,
			wantAntFarm: model.AntFarm{
				Start: model.Room{Name: "A", X: "1", Y: "0"},
				End:   model.Room{Name: "D", X: "1", Y: "2"},
				Rooms: map[string]model.Room{
					"A": {Name: "A", X: "1", Y: "0"},
					"B": {Name: "B", X: "0", Y: "1"},
					"C": {Name: "C", X: "2", Y: "1"},
					"D": {Name: "D", X: "1", Y: "2"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "C", To: "A"},
					{From: "A", To: "D"},
				},
			},
		},
		{
			name:     "Test file with disconnected rooms",
			fileName: "./samples/testfile6",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "Test file with disconnected rooms",
			fileName: "./samples/testfile7",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "Test file with complex layout with multiple valid paths",
			fileName: "./samples/testfile8",
			wantAnts: 0,
			wantErr:  true,
			wantAntFarm: model.AntFarm{
				Start: model.Room{Name: "A", X: "1", Y: "0"},
				End:   model.Room{Name: "G", X: "6", Y: "1"},
				Rooms: map[string]model.Room{
					"A": {Name: "A", X: "1", Y: "0"},
					"B": {Name: "B", X: "0", Y: "1"},
					"C": {Name: "C", X: "2", Y: "1"},
					"D": {Name: "D", X: "3", Y: "2"},
					"E": {Name: "E", X: "4", Y: "1"},
					"F": {Name: "F", X: "5", Y: "0"},
					"G": {Name: "G", X: "6", Y: "1"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "C", To: "D"},
					{From: "D", To: "E"},
					{From: "E", To: "F"},
					{From: "F", To: "G"},
					{From: "B", To: "D"},
					{From: "D", To: "G"},
				},
			},
		},
		{
			name:     "Test file with link pointing to non-existent room",
			fileName: "./samples/testfile9",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "Test file with ant count as string instead of number",
			fileName: "./samples/testfile10",
			wantAnts: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnts, gotAntFarm, gotErr := utils.ParseFile(tt.fileName)

			if gotErr != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			
			if gotAnts != tt.wantAnts {
				t.Errorf("ParseFile() gotAnts = %v, wantAnts %v", gotAnts, tt.wantAnts)
				return
			}

			if reflect.DeepEqual(gotAntFarm, tt.wantAntFarm) {
				t.Errorf("ParseFile() gotAntFarm = %v, wantAntFarm %v", gotAntFarm, tt.wantAntFarm)
			}
		})
	}
}

func TestFindPaths(t *testing.T) {
	tests := []struct {
		name      string
		antFarm   model.AntFarm
		wantPaths []model.Path
	}{
		{
			name: "Single path",
			antFarm: model.AntFarm{
				Start: model.Room{Name: "A"},
				End:   model.Room{Name: "C"},
				Rooms: map[string]model.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
				},
			},
			wantPaths: []model.Path{
				{
					Rooms: []model.Room{
						{Name: "A"},
						{Name: "B"},
						{Name: "C"},
					},
				},
			},
		},
		{
			name: "Multiple paths",
			antFarm: model.AntFarm{
				Start: model.Room{Name: "A"},
				End:   model.Room{Name: "C"},
				Rooms: map[string]model.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
					"D": {Name: "D"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "A", To: "D"},
					{From: "D", To: "C"},
				},
			},
			wantPaths: []model.Path{
				{
					Rooms: []model.Room{
						{Name: "A"},
						{Name: "B"},
						{Name: "C"},
					},
				},
				{
					Rooms: []model.Room{
						{Name: "A"},
						{Name: "D"},
						{Name: "C"},
					},
				},
			},
		},
		{
			name: "No path",
			antFarm: model.AntFarm{
				Start: model.Room{Name: "A"},
				End:   model.Room{Name: "C"},
				Rooms: map[string]model.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
				},
			},
			wantPaths: []model.Path{},
		},
		{
			name: "Multiple paths but one already visited",
			antFarm: model.AntFarm{
				Start: model.Room{Name: "A"},
				End:   model.Room{Name: "C"},
				Rooms: map[string]model.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
					"D": {Name: "D"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "A", To: "D"},
					{From: "D", To: "C"},
				},
			},
			wantPaths: []model.Path{
				{
					Rooms: []model.Room{
						{Name: "A"},
						{Name: "B"},
						{Name: "C"},
					},
				},
			},
		},
		{
			name: "Multiple rooms but no path",
			antFarm: model.AntFarm{
				Start: model.Room{Name: "A"},
				End:   model.Room{Name: "C"},
				Rooms: map[string]model.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
				},
				Links: []model.Link{
					{From: "A", To: "B"},
					{From: "B", To: "B"},
				},
			},
			wantPaths: []model.Path{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPaths := utils.FindPaths(tt.antFarm); !reflect.DeepEqual(gotPaths, tt.wantPaths) {
				t.Errorf("FindPaths() = %v, want %v", gotPaths, tt.wantPaths)
			}
		})
	}
}
