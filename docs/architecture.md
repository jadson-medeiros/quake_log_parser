### docs/architecture.md

The `architecture.md` provides a detailed overview of your project's architecture. This helps new developers understand how the system is structured and how different components interact.

# Quake Log Parser Architecture

## Overview

The Quake Log Parser consists of several layers and components that collaborate to read game logs, analyze data, and generate JSON format reports. Below is a detailed description of each component and its responsibilities.

## Project Structure

```
quake_log_parser/
├── assets/
│   └── logs/
│       └── qgames.log
├── cmd/
│   └── quake_log_parser/
│       └── main.go
├── internal/
│   ├── input/
│   │   └── input.go
│   │   └── input_test.go
│   ├── output/
│   │   └── json.go
│   │   └── output_test.go
│   │   └── report.go
│   ├── parser/
│   │   └── parser.go
│   │   └── parser_test.go
│   │   └── register.go
│   └── pkg/
│       └── data/
│           └── match.go
├── Makefile
├── .gitignore
├── dockerfile
├── .dockerignore
└── go.mod
```

## Main Components

### Input

- Responsible for reading input logs.
- Files: `internal/input/input.go`, `internal/input/input_test.go`.

### Parser

- Processes logs and extracts relevant information.
- Files: `internal/parser/parser.go`, `internal/parser/parser_test.go`.

### Output

- Generates formatted output (JSON, reports).
- Files: `internal/output/json.go`, `internal/output/report.go`, `internal/output/output_test.go`.

## Data Flow

1. **Log Reading:**: Logs are read from the `qgames.log` file by the `input` component.
2. **Parsing**: The `parser` component analyzes the logs and extracts relevant data.
3. **Output Generation**: The `output` component formats the extracted data and generates the desired output (JSON, report).

## Technical Details

- Language: Go
- Data Structures: Structs to represent matches, players, etc.
- Algorithms: Log parsing, report generation.

## Component Descriptions

### assets/

Contains static files, such as game logs.

### logs/:

Directory where log files are stored. Example: `qgames.log`.

### cmd/

Contains the application's entry point.

### quake_log_parser/

Directory containing `main.go`, which is the main entry point of the application.

### internal/

Contains the core components of the application logic, divided into subdirectories based on their responsibilities.

### input/

Contains logic for reading and processing the log file input.

- **input.go**: Code for reading logs.
- **input_test.go**: Tests for log reading logic.

### output/

Contains logic for formatting and saving results.

- **json.go**: Code for formatting data into JSON.
- **output_test.go**: Tests for JSON formatting logic.
- **report.go**: Code for generating reports from analyzed data.

### parser/

Contains logic for parsing game logs.

- **parser.go**: Main code for parsing logs.
- **parser_test.go**: Tests for log parsing logic.
- **register.go**: Code for registering game events (such as kills and players).

### pkg/data/

Contains data structures used in the application..

- **match.go**: Data structures for storing match information.

### Makefile

Contains commands to automate common tasks such as building, testing, and formatting code.

### .gitignore

List of files and directories to be ignored by Git.

### dockerfile

Configuration file to build the Docker image of the application.

### .dockerignore

List of files and directories to be ignored by Docker.

### go.mod

Dependency management file for Go.

## Final Considerations

This architecture documentation provides a detailed view of the components and data flow within the Quake Log Parser. It should serve as a guide for understanding the project structure and facilitate maintenance and contribution.

## Future Plans

- Add support for different log formats.
- Improve parsing performance with more efficient algorithms.
- Implement a graphical interface for report visualization.
- Implement more test scenarios for critical functions.
