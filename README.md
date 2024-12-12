# Advent of Code 2024 - Go Solutions

This repository contains my solutions to [Advent of Code 2024](https://adventofcode.com) by Eric Wastl, implemented in Golang. 
Each day's solution is organized into its own folder with the necessary files for that day's puzzle.

---

## Project Structure

The repository follows this structure:

```
.
├── 01
│   ├── data.txt        # input data for day 1 (ignored in git)
│   ├── main.go         # solution for day 1
├── 02
│   ├── data.txt        # input data for day 2 (ignored in git)
│   ├── main.go         # solution for day 2
...
├── .gitignore          
├── go.mod              # Go module file for dependency management
```

The `data.txt` files, which contain puzzle inputs, are excluded from version control in `.gitignore` to comply with Advent of Code rules.

---

## Run

Each day's solution is self-contained in its respective folder. You can run the solution for a specific day using the `go run` command:

```bash
# example for day 1
cd 01  
go run main.go
```

Ensure that your own input puzzles as `data.txt` file for the specific day is present in the corresponding folder, as the solutions will read input data from this file.

---

## Prerequisites

- [Go](https://go.dev/) 
