package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type Vehicle struct {
	Name              string
	Values            [12]int
	CurrentYearCount  int
	PreviousYearCount int
}

func main() {
	nameFilter := flag.String("n", "", "Filter vehicles by name (case-insensitive partial match)")
	yearFlag := flag.Int("y", time.Now().Year(), "Year for report file (default: current year)")
	monthFlag := flag.Int("m", 0, "Show specific month data (1-12, default: 0 for current year total)")

	flag.Parse()
	
	currentYear := *yearFlag
	
	// Convert nameFilter to uppercase for consistent comparison
	filterUpper := strings.ToUpper(*nameFilter)

	// Create dynamic report filename with current year
	reportFilename := "report-" + strconv.Itoa(currentYear) + ".xlsx"

	f, err := excelize.OpenFile(reportFilename)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all worksheet names
	sheets := f.GetSheetList()

	// Read only tab 8 (index 7)
	if len(sheets) < 8 {
		fmt.Printf("Error: Only %d worksheets found, but requested tab 8\n", len(sheets))
		return
	}

	sheetName := sheets[7] // 8th tab (0-indexed)
	
	rows, err := f.GetRows(sheetName)

	if err != nil {
		fmt.Printf("Error reading sheet %s: %v\n", sheetName, err)
		return
	}

	// Display only rows where first column is an integer and second column is not empty
	var vehicles []Vehicle
	
	for _, row := range rows {
		if len(row) < 3 {
			continue
		}
		
		// Check if first column is an integer
		_, err := strconv.Atoi(row[0])

		if err != nil {
			continue // Skip rows where first column is not an integer
		}
		
		// Check if second column is not empty
		if row[1] == "" {
			continue // Skip rows where second column is empty
		}
		
		// Read second column as vehicle brand, third column as vehicle model
		vehicleBrand := row[1]
		vehicleModel := ""

		if len(row) > 2 {
			vehicleModel = row[2]
		}
		
		// Combine brand and model into vehicle name
		vehicleName := vehicleBrand + " " + vehicleModel
		
		// Extract array of 12 integers from columns 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26
		var values [12]int

		for i := 0; i < 12; i++ {
			columnIndex := 3 + (i * 2) // Column 4 is index 3, then 5 (index 5), 7 (index 6), etc.
			if columnIndex < len(row) {
				if val, err := strconv.Atoi(row[columnIndex]); err == nil {
					values[i] = val
				}
			}
		}
		
		// Extract current year count from column 28 (index 27)
		var currentYearCount int

		if len(row) > 27 {
			if val, err := strconv.Atoi(row[27]); err == nil {
				currentYearCount = val
			}
		}
		
		// Extract previous year count from column 32 (index 31)
		var previousYearCount int

		if len(row) > 31 {
			if val, err := strconv.Atoi(row[31]); err == nil {
				previousYearCount = val
			}
		}
		
		vehicle := Vehicle{
			Name:              vehicleName,
			Values:            values,
			CurrentYearCount:  currentYearCount,
			PreviousYearCount: previousYearCount,
		}
		
		vehicles = append(vehicles, vehicle)
	}
	
	// Sort vehicles by current year count descending
	sort.Slice(vehicles, func(i, j int) bool {
		return vehicles[i].CurrentYearCount > vehicles[j].CurrentYearCount
	})
	
	// Display vehicles filtered by name (if provided)
	count := 0
	totalCurrentYear := 0

	for i, vehicle := range vehicles {
		vehicleNameUpper := strings.ToUpper(vehicle.Name)
		
		if filterUpper != "" && !strings.Contains(vehicleNameUpper, filterUpper) {
			continue
		}

		var displayValue int
		
		if *monthFlag > 0 && *monthFlag <= 12 {
			displayValue = vehicle.Values[*monthFlag - 1] // Convert 1-based month to 0-based index
		} else {
			displayValue = vehicle.CurrentYearCount
		}
		
		fmt.Printf("%3d %-35s %5d\n", i, vehicle.Name, displayValue)
		count++
		totalCurrentYear += displayValue
	}
	
	fmt.Printf("%-39s %5d\n", "TOTAL", totalCurrentYear)
}

