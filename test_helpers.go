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

// Test helpers and mock data generation for testing purposes

// CreateMockVehicles generates a slice of mock vehicles for testing
func CreateMockVehicles() []Vehicle {
	return []Vehicle{
		{
			Name:              "Toyota Corolla",
			Values:            [12]int{45, 52, 68, 71, 85, 92, 88, 95, 103, 89, 75, 67},
			CurrentYearCount:  930,
			PreviousYearCount: 850,
		},
		{
			Name:              "BMW X5",
			Values:            [12]int{30, 35, 42, 48, 55, 60, 58, 62, 65, 58, 50, 45},
			CurrentYearCount:  608,
			PreviousYearCount: 580,
		},
		{
			Name:              "Volkswagen Golf",
			Values:            [12]int{55, 60, 72, 78, 85, 90, 88, 92, 95, 85, 78, 70},
			CurrentYearCount:  948,
			PreviousYearCount: 920,
		},
		{
			Name:              "Audi A4",
			Values:            [12]int{25, 28, 32, 35, 40, 45, 42, 48, 50, 45, 38, 32},
			CurrentYearCount:  460,
			PreviousYearCount: 440,
		},
		{
			Name:              "Mercedes C-Class",
			Values:            [12]int{35, 40, 48, 52, 58, 62, 60, 65, 68, 62, 55, 48},
			CurrentYearCount:  653,
			PreviousYearCount: 620,
		},
		{
			Name:              "Škoda Octavia",
			Values:            [12]int{65, 70, 82, 88, 95, 100, 98, 105, 108, 98, 88, 80},
			CurrentYearCount:  1077,
			PreviousYearCount: 1050,
		},
		{
			Name:              "Renault Clio",
			Values:            [12]int{40, 45, 52, 58, 62, 68, 65, 70, 72, 65, 58, 52},
			CurrentYearCount:  707,
			PreviousYearCount: 680,
		},
		{
			Name:              "Ford Focus",
			Values:            [12]int{30, 32, 38, 42, 45, 50, 48, 52, 55, 50, 42, 38},
			CurrentYearCount:  522,
			PreviousYearCount: 500,
		},
	}
}

// CreateEmptyVehicle returns a vehicle with zero values for testing edge cases
func CreateEmptyVehicle() Vehicle {
	return Vehicle{
		Name:              "Test Empty",
		Values:            [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		CurrentYearCount:  0,
		PreviousYearCount: 0,
	}
}

// CreateVehicleWithName creates a custom vehicle with specified name for testing
func CreateVehicleWithName(name string, currentYearCount int) Vehicle {
	var values [12]int
	for i := 0; i < 12; i++ {
		values[i] = currentYearCount / 12
	}

	return Vehicle{
		Name:              name,
		Values:            values,
		CurrentYearCount:  currentYearCount,
		PreviousYearCount: currentYearCount - 50,
	}
}

// GetTotalForVehicle calculates the sum of all monthly values
func GetTotalForVehicle(v Vehicle) int {
	total := 0
	for _, val := range v.Values {
		total += val
	}
	return total
}
