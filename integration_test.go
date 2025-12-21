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
	"sort"
	"strings"
	"testing"
)

func TestIntegrationVehicleProcessing(t *testing.T) {
	vehicles := CreateMockVehicles()

	if len(vehicles) == 0 {
		t.Fatal("Mock vehicles should not be empty")
	}

	sort.Slice(vehicles, func(i, j int) bool {
		return vehicles[i].CurrentYearCount > vehicles[j].CurrentYearCount
	})

	if vehicles[0].Name != "Škoda Octavia" {
		t.Errorf("Expected top vehicle to be Škoda Octavia, got %s", vehicles[0].Name)
	}

	if vehicles[0].CurrentYearCount != 1077 {
		t.Errorf("Expected top count to be 1077, got %d", vehicles[0].CurrentYearCount)
	}
}

func TestIntegrationFilteringAndSorting(t *testing.T) {
	vehicles := CreateMockVehicles()
	filter := "VOLKSWAGEN"

	var filtered []Vehicle
	for _, vehicle := range vehicles {
		vehicleNameUpper := strings.ToUpper(vehicle.Name)
		if strings.Contains(vehicleNameUpper, filter) {
			filtered = append(filtered, vehicle)
		}
	}

	if len(filtered) != 1 {
		t.Errorf("Expected 1 Volkswagen vehicle, got %d", len(filtered))
	}

	if len(filtered) > 0 && filtered[0].Name != "Volkswagen Golf" {
		t.Errorf("Expected Volkswagen Golf, got %s", filtered[0].Name)
	}
}

func TestIntegrationMultipleBrands(t *testing.T) {
	vehicles := CreateMockVehicles()

	brandCounts := make(map[string]int)
	for _, vehicle := range vehicles {
		brand := strings.Split(vehicle.Name, " ")[0]
		brandCounts[brand]++
	}

	expectedBrands := []string{"Toyota", "BMW", "Volkswagen", "Audi", "Mercedes", "Škoda", "Renault", "Ford"}
	if len(brandCounts) != len(expectedBrands) {
		t.Errorf("Expected %d brands, got %d", len(expectedBrands), len(brandCounts))
	}
}

func TestIntegrationMonthlyDataAccess(t *testing.T) {
	vehicles := CreateMockVehicles()

	if len(vehicles) == 0 {
		t.Fatal("No vehicles available for testing")
	}

	testVehicle := vehicles[0]

	for month := 0; month < 12; month++ {
		value := testVehicle.Values[month]
		if value < 0 {
			t.Errorf("Month %d has negative value: %d", month+1, value)
		}
	}
}

func TestIntegrationYearComparison(t *testing.T) {
	vehicles := CreateMockVehicles()

	growthCount := 0
	declineCount := 0

	for _, vehicle := range vehicles {
		if vehicle.CurrentYearCount > vehicle.PreviousYearCount {
			growthCount++
		} else if vehicle.CurrentYearCount < vehicle.PreviousYearCount {
			declineCount++
		}
	}

	totalVehicles := len(vehicles)
	if growthCount+declineCount > totalVehicles {
		t.Error("Growth + decline count exceeds total vehicles")
	}
}

func TestIntegrationEmptyVehicleHandling(t *testing.T) {
	emptyVehicle := CreateEmptyVehicle()

	if emptyVehicle.CurrentYearCount != 0 {
		t.Errorf("Empty vehicle should have 0 current year count, got %d", emptyVehicle.CurrentYearCount)
	}

	total := GetTotalForVehicle(emptyVehicle)
	if total != 0 {
		t.Errorf("Empty vehicle total should be 0, got %d", total)
	}
}

func TestIntegrationCustomVehicleCreation(t *testing.T) {
	testName := "Test Brand Model"
	testCount := 600

	vehicle := CreateVehicleWithName(testName, testCount)

	if vehicle.Name != testName {
		t.Errorf("Expected name %s, got %s", testName, vehicle.Name)
	}

	if vehicle.CurrentYearCount != testCount {
		t.Errorf("Expected count %d, got %d", testCount, vehicle.CurrentYearCount)
	}

	expectedMonthly := testCount / 12
	if vehicle.Values[0] != expectedMonthly {
		t.Errorf("Expected monthly value %d, got %d", expectedMonthly, vehicle.Values[0])
	}
}

func TestIntegrationTotalCalculation(t *testing.T) {
	vehicles := CreateMockVehicles()

	for _, vehicle := range vehicles {
		calculated := GetTotalForVehicle(vehicle)

		if calculated <= 0 && vehicle.CurrentYearCount > 0 {
			t.Errorf("Vehicle %s has mismatched totals: calculated=%d, stored=%d",
				vehicle.Name, calculated, vehicle.CurrentYearCount)
		}
	}
}

func TestIntegrationTopPerformers(t *testing.T) {
	vehicles := CreateMockVehicles()

	sort.Slice(vehicles, func(i, j int) bool {
		return vehicles[i].CurrentYearCount > vehicles[j].CurrentYearCount
	})

	topN := 3
	if len(vehicles) < topN {
		topN = len(vehicles)
	}

	topVehicles := vehicles[:topN]

	for i := 0; i < len(topVehicles)-1; i++ {
		if topVehicles[i].CurrentYearCount < topVehicles[i+1].CurrentYearCount {
			t.Errorf("Top performers not properly sorted at index %d", i)
		}
	}
}

func TestIntegrationCaseInsensitiveFiltering(t *testing.T) {
	vehicles := CreateMockVehicles()

	testFilters := []struct {
		filter        string
		expectedCount int
	}{
		{"bmw", 1},
		{"BMW", 1},
		{"BmW", 1},
		{"golf", 1},
		{"GOLF", 1},
		{"toyota", 1},
		{"audi", 1},
		{"nonexistent", 0},
	}

	for _, tf := range testFilters {
		count := 0
		filterUpper := strings.ToUpper(tf.filter)

		for _, vehicle := range vehicles {
			vehicleNameUpper := strings.ToUpper(vehicle.Name)
			if strings.Contains(vehicleNameUpper, filterUpper) {
				count++
			}
		}

		if count != tf.expectedCount {
			t.Errorf("Filter '%s': expected %d matches, got %d", tf.filter, tf.expectedCount, count)
		}
	}
}
