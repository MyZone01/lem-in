package main

import (
	"fmt"
	"lemin/lib"
	"lemin/models"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		wantAnts    int
		wantErr     bool
		wantAntFarm models.AntFarm
	}{
		{
			name:     "Test file with valid input",
			fileName: "./samples/testfile1",
			wantAnts: 10,
			wantErr:  false,
			wantAntFarm: models.AntFarm{
				Start: models.Room{Name: "A", X: "1", Y: "0"},
				End:   models.Room{Name: "E", X: "0", Y: "3"},
				Rooms: map[string]models.Room{
					"A": {Name: "A", X: "1", Y: "0"},
					"B": {Name: "B", X: "0", Y: "1"},
					"C": {Name: "C", X: "2", Y: "1"},
					"D": {Name: "D", X: "1", Y: "2"},
					"E": {Name: "E", X: "0", Y: "3"},
				},
				Links: []models.Link{
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
			fileName: "./samples/testfile3",
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
			wantAntFarm: models.AntFarm{
				Start: models.Room{Name: "A", X: "1", Y: "0"},
				End:   models.Room{Name: "D", X: "1", Y: "2"},
				Rooms: map[string]models.Room{
					"A": {Name: "A", X: "1", Y: "0"},
					"B": {Name: "B", X: "0", Y: "1"},
					"C": {Name: "C", X: "2", Y: "1"},
					"D": {Name: "D", X: "1", Y: "2"},
				},
				Links: []models.Link{
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
			wantAnts: 10,
			wantErr:  false,
			wantAntFarm: models.AntFarm{
				Start: models.Room{Name: "A", X: "1", Y: "0"},
				End:   models.Room{Name: "G", X: "6", Y: "1"},
				Rooms: map[string]models.Room{
					"A": {Name: "A", X: "1", Y: "0"},
					"B": {Name: "B", X: "0", Y: "1"},
					"C": {Name: "C", X: "2", Y: "1"},
					"D": {Name: "D", X: "3", Y: "2"},
					"E": {Name: "E", X: "4", Y: "1"},
					"F": {Name: "F", X: "5", Y: "0"},
					"G": {Name: "G", X: "6", Y: "1"},
				},
				Links: []models.Link{
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.fileName)
			t.Log(test.name)
			gotAnts, gotAntFarm, _, gotErr := lib.ParseFile(test.fileName)
			_gotAntFarm := fmt.Sprintf("%v", gotAntFarm)
			_wantAntFarm := fmt.Sprintf("%v", test.wantAntFarm)

			if gotErr != test.wantErr {
				t.Errorf("❌ ParseFile() error = %v, wantErr %v", gotErr, test.wantErr)
				return
			}

			if gotAnts != test.wantAnts {
				t.Errorf("❌ ParseFile() gotAnts = %v, wantAnts %v", gotAnts, test.wantAnts)
				return
			}

			if _gotAntFarm != _wantAntFarm {
				t.Errorf("❌ ParseFile() gotAntFarm = %v, wantAntFarm %v", gotAntFarm, test.wantAntFarm)
			} else {
				t.Log("✅ ", test.name, " Succeeded")
			}
		})
	}
}

func TestFindPaths(t *testing.T) {
	tests := []struct {
		name      string
		antFarm   models.AntFarm
		wantPaths []models.Path
	}{
		{
			name: "Single path",
			antFarm: models.AntFarm{
				Start: models.Room{Name: "A"},
				End:   models.Room{Name: "C"},
				Rooms: map[string]models.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
				},
				Links: []models.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
				},
			},
			wantPaths: []models.Path{
				{
					Rooms: []models.Room{
						{Name: "B"},
						{Name: "C"},
					},
				},
			},
		},
		{
			name: "Multiple paths",
			antFarm: models.AntFarm{
				Start: models.Room{Name: "A"},
				End:   models.Room{Name: "C"},
				Rooms: map[string]models.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
					"D": {Name: "D"},
				},
				Links: []models.Link{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "A", To: "D"},
					{From: "D", To: "C"},
				},
			},
			wantPaths: []models.Path{
				{
					Rooms: []models.Room{
						{Name: "B"},
						{Name: "C"},
					},
				},
				{
					Rooms: []models.Room{
						{Name: "D"},
						{Name: "C"},
					},
				},
			},
		},
		{
			name: "No path",
			antFarm: models.AntFarm{
				Start: models.Room{Name: "A"},
				End:   models.Room{Name: "C"},
				Rooms: map[string]models.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
				},
				Links: []models.Link{
					{From: "A", To: "B"},
				},
			},
			wantPaths: []models.Path{},
		},
		{
			name: "Multiple rooms but no path",
			antFarm: models.AntFarm{
				Start: models.Room{Name: "A"},
				End:   models.Room{Name: "C"},
				Rooms: map[string]models.Room{
					"A": {Name: "A"},
					"B": {Name: "B"},
					"C": {Name: "C"},
				},
				Links: []models.Link{
					{From: "A", To: "B"},
					{From: "B", To: "B"},
				},
			},
			wantPaths: []models.Path{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPaths := lib.FindPaths(tt.antFarm); !reflect.DeepEqual(gotPaths, tt.wantPaths) && !(len(gotPaths) == 0 && len(tt.wantPaths) == 0) {
				t.Errorf("FindPaths() = %v, want %v", gotPaths, tt.wantPaths)
			}
		})
	}
}
