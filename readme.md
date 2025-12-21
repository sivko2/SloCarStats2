# SloCarStats - Slovenian Car Statistics Analyzer

A Go-based command-line application for analyzing Slovenian vehicle registration statistics. This tool processes Excel files from the Slovenian Automobile Association (ADS) and provides filtered, sortable views of vehicle registration data by brand, model, year, and month.

## Description

SloCarStats reads Excel spreadsheets containing vehicle registration data and provides:
- Vehicle statistics sorted by registration count (descending)
- Filtering by vehicle brand/model name (case-insensitive partial matching)
- Year-specific data analysis
- Month-specific registration counts
- Total registration summaries

The application processes data from the 8th worksheet tab of the Excel file, extracting vehicle brand, model, monthly registration counts, and yearly totals.

## Requirements

- Go 1.23.0 or later
- Excel file with vehicle statistics (`.xlsx` format)

## Dependencies

This project uses the following external libraries:

- **github.com/xuri/excelize/v2** (v2.9.1) - Excel file reading and processing
  - **github.com/xuri/efp** (v0.0.1) - Excel formula parser
  - **github.com/xuri/nfp** (v0.0.1) - Number format parser
- **github.com/richardlehane/mscfb** (v1.0.4) - Microsoft Compound Document File Binary format
- **github.com/richardlehane/msoleps** (v1.0.4) - Microsoft OLE Property Set format
- **github.com/tiendc/go-deepcopy** (v1.6.0) - Deep copy utilities
- **golang.org/x/crypto** (v0.38.0) - Cryptographic functions
- **golang.org/x/net** (v0.40.0) - Network utilities
- **golang.org/x/text** (v0.25.0) - Text processing

## Installation

### Prerequisites
1. Install Go 1.23.0 or later from [golang.org](https://golang.org/downloads/)
2. Clone or download this repository

### Build from Source
```shell
# Clone the repository (if not already done)
git clone https://github.com/sivko2/SloCarStats2.git
cd SloCarStats2

# Download dependencies
go mod tidy

# Build the executable
go build

# This creates carstats.exe on Windows or carstats on Unix systems
```

### Alternative Build Commands
```shell
# Build with specific output name
go build -o carstats

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o carstats-linux
GOOS=windows GOARCH=amd64 go build -o carstats.exe
```

## Data Source

Statistics data is available from the Slovenian Automobile Association (ADS):
- Website: https://www.ads-slo.org/statistika
- Download the Excel file for the desired year
- Save as `report-<YEAR>.xlsx` (e.g., `report-2025.xlsx`) in the application directory

## Usage

### Basic Execution

**Windows:**
```shell
carstats.exe
```

**Linux/macOS:**
```shell
./carstats
```

### Command Line Options

| Option | Description | Example |
|--------|-------------|---------|
| `-n` | Filter by vehicle name (case-insensitive partial match) | `-n volkswagen` |
| `-y` | Specify year for data (default: current year) | `-y 2025` |
| `-m` | Show specific month data (1-12, default: yearly total) | `-m 3` |

### Examples

```shell
# Show all vehicles for current year
carstats.exe

# Filter by brand name
carstats.exe -n toyota

# Show 2024 statistics
carstats.exe -y 2024

# Show March statistics
carstats.exe -m 3

# Show April 2024 Toyota statistics
carstats.exe -n toyota -y 2024 -m 4

# Show Volkswagen statistics for 2024
carstats.exe -n volkswagen -y 2024
```

### Output Format
```
#001   1 VOLKSWAGEN GOLF                   1250
#002   6 TOYOTA RAV4                       1180
#003  13 RENAULT CLIO                      1050
...
TOTAL                                       15430
```

## Development

### Testing

The project includes a comprehensive test suite covering unit tests, integration tests, and performance benchmarks.

#### Test Files
- **main_test.go** - Unit tests for core functionality
- **integration_test.go** - Integration tests with mock data
- **test_helpers.go** - Test utilities and mock data generators
- **TESTING.md** - Detailed test documentation

#### Running Tests

```shell
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestVehicleSorting

# Run integration tests only
go test -run Integration

# Run benchmarks
go test -bench=.

# Run benchmarks with memory stats
go test -bench=. -benchmem

# Check test coverage
go test -cover

# Generate HTML coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run tests with race detection
go test -race ./...
```

#### Test Coverage

The test suite includes:
- **Vehicle struct operations** - Creation, field assignment, and data integrity
- **Sorting algorithms** - Descending sort by registration count
- **Name filtering** - Case-insensitive partial matching
- **Monthly/yearly data** - Value selection and indexing
- **Edge cases** - Empty vehicles, invalid months, boundary conditions
- **Performance benchmarks** - Sorting and filtering performance metrics
- **Integration scenarios** - Full processing pipeline, multi-brand handling, year-over-year comparisons

#### Mock Test Data

The test suite includes realistic mock data for 8 vehicle models with monthly registration counts:
1. Škoda Octavia (1077 units)
2. Volkswagen Golf (948 units)
3. Toyota Corolla (930 units)
4. Renault Clio (707 units)
5. Mercedes C-Class (653 units)
6. BMW X5 (608 units)
7. Ford Focus (522 units)
8. Audi A4 (460 units)

### Static Analysis
```shell
go vet ./...
```

### Code Formatting
```shell
go fmt ./...
```

### Clean Dependencies
```shell
go mod tidy
```

## Project Structure

```
SloCarStats2/
├── main.go              # Main application source code
├── main_test.go         # Unit tests
├── integration_test.go  # Integration tests
├── test_helpers.go      # Test utilities and mock data
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── readme.md            # This documentation
├── TESTING.md           # Test documentation
├── CLAUDE.md            # Development guidelines
├── LICENSE              # GNU GPL v3 license text
├── carstats.exe         # Compiled Windows executable
└── report-<YEAR>.xlsx   # Excel data files
```

## Contributing

This application was developed with assistance from Claude Code. When contributing:

1. Follow Go best practices and formatting (`go fmt`)
2. Run tests and static analysis before submitting changes
3. Update documentation for new features or changes
4. Ensure compatibility with the expected Excel file format

## License

This program is free software: you can redistribute it and/or modify it under the terms of the **GNU General Public License** as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

### GNU GPL v3

SloCarStats is licensed under the GNU General Public License v3.0 (GPL-3.0). This means:

**Permissions:**
- Commercial use
- Modification
- Distribution
- Patent use
- Private use

**Conditions:**
- Disclose source code
- Include original license and copyright
- State changes made to the code
- Use same license for derivatives (copyleft)

**Limitations:**
- No liability
- No warranty

For the full license text, visit: https://www.gnu.org/licenses/gpl-3.0.en.html


