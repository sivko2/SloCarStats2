# Slovenian Car Statistics for year 2025 and 2024

This application was written by Claude Code and dictated/supervised by me :)

## Build

Execute:
```shell
go build
```

## Statistics Database

The statistics is found at https://www.ads-slo.org/statistika.

Download the latest XSLX file for specifi year and save it as *report-<YEAR>.xslx* (e.g., *report-2025.xslx*).

## Run

On Windows 11, inside the command prompt or powershell execute:
```shell
carstats.exe
```

On Linux, inside the shell execute:
```shell
./carstats
```

## Options

To filter by name, use argument *-n* (e.g., *carstats.exe -n renault*).

To show statistics from a specific year, use argument *-y* (e.g., *carstats.exe -y 2024*).

To show statistics from a specific month, use argument *-m* where 1 is January and 12 is December (e.g., *carstats.exe -m 3*).

An example to show statistics for April 2024, only for Toyota:
```shell
carstats.exe -n toyota -y 2024 -m 4
```

