# SloCarStats Test Suite

This document describes the test files and how to use them.

## Test Files

### 1. `main_test.go`
Unit tests for core functionality:
- **TestVehicleStruct**: Tests the Vehicle struct creation and field assignment
- **TestVehicleSorting**: Tests sorting vehicles by current year count (descending)
- **TestNameFiltering**: Tests case-insensitive name filtering logic
- **TestMonthValueSelection**: Tests monthly vs yearly value selection logic
- **TestVehicleArrayIndexing**: Tests array indexing for monthly values
- **BenchmarkVehicleSorting**: Performance benchmark for sorting 1000 vehicles
- **BenchmarkNameFiltering**: Performance benchmark for filtering 1000 vehicles

### 2. `integration_test.go`
Integration tests using mock data:
- **TestIntegrationVehicleProcessing**: Tests full vehicle processing pipeline
- **TestIntegrationFilteringAndSorting**: Tests combined filtering and sorting
- **TestIntegrationMultipleBrands**: Tests handling multiple vehicle brands
- **TestIntegrationMonthlyDataAccess**: Tests accessing monthly data for all vehicles
- **TestIntegrationYearComparison**: Tests year-over-year comparison logic
- **TestIntegrationEmptyVehicleHandling**: Tests edge case with zero values
- **TestIntegrationCustomVehicleCreation**: Tests custom vehicle creation helper
- **TestIntegrationTotalCalculation**: Tests total calculation accuracy
- **TestIntegrationTopPerformers**: Tests top N vehicles selection
- **TestIntegrationCaseInsensitiveFiltering**: Tests various case combinations in filtering

### 3. `test_helpers.go`
Helper functions for testing:
- **CreateMockVehicles()**: Generates realistic mock vehicle data for 8 vehicles
- **CreateEmptyVehicle()**: Creates a vehicle with all zero values for edge case testing
- **CreateVehicleWithName(name, count)**: Creates a custom vehicle with specified parameters
- **GetTotalForVehicle(v)**: Calculates sum of all monthly values

## Running Tests

### Run All Tests
```bash
go test ./...
```

### Run Tests with Verbose Output
```bash
go test -v ./...
```

### Run Specific Test
```bash
go test -run TestVehicleSorting
```

### Run Integration Tests Only
```bash
go test -run Integration
```

### Run Benchmarks
```bash
go test -bench=.
```

### Run Benchmarks with Memory Stats
```bash
go test -bench=. -benchmem
```

### Check Test Coverage
```bash
go test -cover
```

### Generate Coverage Report
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Coverage

The test suite covers:
- Vehicle struct operations
- Sorting algorithms
- Name filtering (case-insensitive, partial matching)
- Monthly and yearly value selection
- Array indexing and bounds checking
- Edge cases (empty vehicles, invalid months)
- Performance benchmarks

## Mock Data

The test suite includes realistic mock data for 8 vehicle models:
1. Škoda Octavia (1077 units) - Top performer
2. Volkswagen Golf (948 units)
3. Toyota Corolla (930 units)
4. Renault Clio (707 units)
5. Mercedes C-Class (653 units)
6. BMW X5 (608 units)
7. Ford Focus (522 units)
8. Audi A4 (460 units)

Each vehicle includes:
- 12 monthly values
- Current year total count
- Previous year total count

## Continuous Integration

To integrate with CI/CD pipelines:
```bash
# Run tests with race detection
go test -race ./...

# Run tests with coverage threshold
go test -cover -coverprofile=coverage.out && \
  go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//'
```

## Best Practices

1. Run tests before committing: `go test ./...`
2. Check formatting: `go fmt ./...`
3. Run static analysis: `go vet ./...`
4. Clean dependencies: `go mod tidy`
5. Run benchmarks periodically to catch performance regressions
