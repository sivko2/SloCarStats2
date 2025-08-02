# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

SloCarStats is a simple Go application for car statistics processing. The project appears to be in early development with a basic "Hello, World!" implementation in main.go.

## Development Commands

### Building and Running
- `go build` - Build the application (produces SloCarStats.exe)
- `go run main.go` - Run the application directly
- `./SloCarStats.exe` - Run the compiled executable (Windows)

### Testing and Quality
- `go test ./...` - Run all tests
- `go fmt ./...` - Format all Go files
- `go vet ./...` - Run static analysis
- `go mod tidy` - Clean up module dependencies

## Project Structure

- `main.go` - Main application entry point
- `go.mod` - Go module definition (Go 1.21.1)
- `SloCarStats.exe` - Compiled executable
- `report-2025.xlsx` - Excel report file (likely contains car statistics data)

## Architecture Notes

This is a single-file Go application currently containing only a basic main function. The presence of `report-2025.xlsx` suggests the application will likely process Excel-based car statistics data in the future.

The module is named "SloCarStats" and uses Go 1.21.1.