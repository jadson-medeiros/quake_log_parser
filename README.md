# Quake Log Parser

Quake Log Parser is a tool for analyzing Quake game logs and generating reports in JSON format. It was developed to fulfill [this assignment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1).

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Architecture Overview](docs/architecture.md)

## Introduction

This project aims to provide a simple and efficient tool to analyze Quake game logs and generate detailed reports in JSON format.

## Requirements

To use this project, you need the following:

- Go programming language
- Git
- Docker (optional)

## Installation

To compile and run the project locally, follow the steps below:

1. Clone the repository:

   ```bash
   git clone https://github.com/jadson-medeiros/quake_log_parser.git
   cd quake_log_parser
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the project:

   ```bash
   make build
   ```

4. Run tests:
   ```bash
   make test
   ```

Alternatively, you can use Docker:

5. Build the Docker image:

   ```bash
   make docker-build
   ```

6. Run the Docker container:
   ```bash
   make docker-run
   ```
