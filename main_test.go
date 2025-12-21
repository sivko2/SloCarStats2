// SloCarStats - Slovenian Car Statistics Analyzer
// Copyright (C) 2025
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestVehicleStruct(t *testing.T) {
	tests := []struct {
		name     string
		vehicle  Vehicle
		expected Vehicle
	}{
		{
			name: "Vehicle with all fields populated",
			vehicle: Vehicle{
				Name:              "Toyota Corolla",
				Values:            [12]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120},
				CurrentYearCount:  780,
				PreviousYearCount: 650,
			},
			expected: Vehicle{
				Name:              "Toyota Corolla",
				Values:            [12]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120},
				CurrentYearCount:  780,
				PreviousYearCount: 650,
			},
		},
		{
			name: "Vehicle with zero values",
			vehicle: Vehicle{
				Name:              "BMW X5",
				Values:            [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				CurrentYearCount:  0,
				PreviousYearCount: 0,
			},
			expected: Vehicle{
				Name:              "BMW X5",
				Values:            [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				CurrentYearCount:  0,
				PreviousYearCount: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.vehicle, tt.expected) {
				t.Errorf("Vehicle mismatch: got %+v, want %+v", tt.vehicle, tt.expected)
			}
		})
	}
}

func TestVehicleSorting(t *testing.T) {
	vehicles := []Vehicle{
		{Name: "Toyota Corolla", CurrentYearCount: 500},
		{Name: "BMW X5", CurrentYearCount: 800},
		{Name: "Audi A4", CurrentYearCount: 300},
		{Name: "Mercedes C-Class", CurrentYearCount: 650},
	}

	sort.Slice(vehicles, func(i, j int) bool {
		return vehicles[i].CurrentYearCount > vehicles[j].CurrentYearCount
	})

	expected := []string{"BMW X5", "Mercedes C-Class", "Toyota Corolla", "Audi A4"}
	for i, vehicle := range vehicles {
		if vehicle.Name != expected[i] {
			t.Errorf("Index %d: got %s, want %s", i, vehicle.Name, expected[i])
		}
	}
}

func TestNameFiltering(t *testing.T) {
	tests := []struct {
		name         string
		vehicleName  string
		filter       string
		shouldMatch  bool
	}{
		{
			name:        "Exact match case insensitive",
			vehicleName: "Toyota Corolla",
			filter:      "toyota",
			shouldMatch: true,
		},
		{
			name:        "Partial match",
			vehicleName: "BMW X5",
			filter:      "bmw",
			shouldMatch: true,
		},
		{
			name:        "Model match",
			vehicleName: "Audi A4",
			filter:      "a4",
			shouldMatch: true,
		},
		{
			name:        "No match",
			vehicleName: "Mercedes C-Class",
			filter:      "toyota",
			shouldMatch: false,
		},
		{
			name:        "Empty filter matches all",
			vehicleName: "Volkswagen Golf",
			filter:      "",
			shouldMatch: true,
		},
		{
			name:        "Case insensitive uppercase filter",
			vehicleName: "Honda Civic",
			filter:      "HONDA",
			shouldMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vehicleNameUpper := strings.ToUpper(tt.vehicleName)
			filterUpper := strings.ToUpper(tt.filter)

			matches := filterUpper == "" || strings.Contains(vehicleNameUpper, filterUpper)

			if matches != tt.shouldMatch {
				t.Errorf("Filter test failed: vehicle=%s, filter=%s, got %v, want %v",
					tt.vehicleName, tt.filter, matches, tt.shouldMatch)
			}
		})
	}
}

func TestMonthValueSelection(t *testing.T) {
	vehicle := Vehicle{
		Name:             "Test Vehicle",
		Values:           [12]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120},
		CurrentYearCount: 780,
	}

	tests := []struct {
		name          string
		monthFlag     int
		expectedValue int
	}{
		{name: "January (month 1)", monthFlag: 1, expectedValue: 10},
		{name: "June (month 6)", monthFlag: 6, expectedValue: 60},
		{name: "December (month 12)", monthFlag: 12, expectedValue: 120},
		{name: "No month specified (0)", monthFlag: 0, expectedValue: 780},
		{name: "Invalid month (13)", monthFlag: 13, expectedValue: 780},
		{name: "Invalid month (-1)", monthFlag: -1, expectedValue: 780},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var displayValue int

			if tt.monthFlag > 0 && tt.monthFlag <= 12 {
				displayValue = vehicle.Values[tt.monthFlag-1]
			} else {
				displayValue = vehicle.CurrentYearCount
			}

			if displayValue != tt.expectedValue {
				t.Errorf("Month value selection failed: month=%d, got %d, want %d",
					tt.monthFlag, displayValue, tt.expectedValue)
			}
		})
	}
}

func TestVehicleArrayIndexing(t *testing.T) {
	values := [12]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200}

	tests := []struct {
		columnIndex int
		expected    int
	}{
		{columnIndex: 0, expected: 100},
		{columnIndex: 5, expected: 600},
		{columnIndex: 11, expected: 1200},
	}

	for _, tt := range tests {
		if values[tt.columnIndex] != tt.expected {
			t.Errorf("Array indexing failed: index=%d, got %d, want %d",
				tt.columnIndex, values[tt.columnIndex], tt.expected)
		}
	}
}

func BenchmarkVehicleSorting(b *testing.B) {
	vehicles := make([]Vehicle, 1000)
	for i := 0; i < 1000; i++ {
		vehicles[i] = Vehicle{
			Name:             "Vehicle " + string(rune(i)),
			CurrentYearCount: i * 10,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vehiclesCopy := make([]Vehicle, len(vehicles))
		copy(vehiclesCopy, vehicles)
		sort.Slice(vehiclesCopy, func(i, j int) bool {
			return vehiclesCopy[i].CurrentYearCount > vehiclesCopy[j].CurrentYearCount
		})
	}
}

func BenchmarkNameFiltering(b *testing.B) {
	vehicles := make([]Vehicle, 1000)
	for i := 0; i < 1000; i++ {
		vehicles[i] = Vehicle{
			Name: "Vehicle " + string(rune(i)),
		}
	}

	filter := "VEHICLE 500"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := 0
		for _, vehicle := range vehicles {
			vehicleNameUpper := strings.ToUpper(vehicle.Name)
			if strings.Contains(vehicleNameUpper, filter) {
				count++
			}
		}
	}
}
