# Golang Advent of Code 2023

This Go program is a REST API that provides solutions for Advent of Code challenges. The solutions are organized by day in folders named "first", "second", and so on. Currently, solutions are available for the first three days.

Each day's folder contains the following subfolders:

-   entities: Contains data structures used in the solutions.
-   handlers: Includes HTTP handlers to process API requests.
-   helpers: Provides auxiliary functions and utilities used in the solutions.
-   mapper: Contains functions for mapping and converting data.

## Prerequisites

-   Go installed on the system. You can download it from [Oficial Download Page](https://golang.org/dl/)

# Running Dev environment

1. Clone this repository:

```bash
git clone https://github.com/Rdani2005/GolangAdventOfCode.git
```

2. Navigate to project directory:

```bash
cd GolangAdventOfCode
```

3. Install project dependencies:

```bash
go mod tidy
```

4. Run the application on development mode:

```bash
# Using Go command
go run ./main.go

# Using Compile Daemon
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -command=".GolangAdventageCode.exe" # Windows base system
CompileDaemon -command="./GolangAdventageCode" # Unix Base system
```

The server will be running at [localhost:8080](http://localhost:8080).

## Compilation and execution in the production environment

1. Navigate to project directory:

```bash
cd GolangAdventOfCode
```

2. Build app Executable:

```bash
go build
```

3. Run Executable:

```bash
GolangAdventageCode.exe # Windows base system
./GolangAdventageCode # Unix Base system
```

The server will be running at [localhost:8080](http://localhost:8080).

## API Structure

The API is structured using the [gorilla/mux](https://github.com/gorilla/mux) package for routing. Below is the breakdown of the API structure based on the provided code:

-   **Day 1**:

    -   Endpoint for reading results by numbers only: [/api/v1/dayone/readResultsByNumbersOnly](http://localhost:8080/api/v1/dayone/readResultsByNumbersOnly)
    -   Endpoint for reading results by words and numbers: [/api/v1/dayone/readResultsByWordsAndNumbers](http://localhost:8080/api/v1/dayone/readResultsByWordsAndNumbers)

-   **Day 2**:

    -   Endpoint for reading valid games: [/api/v1/daytwo/readValidGames](http://localhost:8080/api/v1/daytwo/readValidGames)
    -   Endpoint for reading total power: [/api/v1/daytwo/readTotalPower](http://localhost:8080/api/v1/daytwo/readTotalPower)

-   **Day 3**:
    -   Endpoint for reading by number: [/api/v1/daythree/readByNumber](http://localhost:8080/api/v1/daythree/readByNumber)
    -   Endpoint for reading all ratios' sum by gear: [/api/v1/daythree/readByGear](http://localhost:8080/api/v1/daythree/readByGear)

Each endpoint corresponds to a specific function (Handle\*) that handles the corresponding Advent of Code challenge for that day. The methods specified for each endpoint are HTTP GET methods. You can customize these endpoints and methods based on the specific challenges and requirements for each day.
