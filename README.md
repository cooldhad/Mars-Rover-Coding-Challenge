# Mars-Rover-Coding-Challenge
A coding challenge to demonstrate proficiency in solving problems using software engineering tools and processes.

## Prerequisites

- Go 1.23.1 or higher
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd Mars-Rover-Coding-Challenge
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```
   
   Or using Make:
   ```bash
   make deps
   make tidy
   ```

## Running the Project

### Option 1: Using Go directly

Run the program interactively:
```bash
go run cmd/main.go
```

### Option 2: Using Make commands

Run the program:

``` bash
make run
 ```

Build and run the binary:

``` bash
make build
./mars_rover
```

Based on the challenge, The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are assumed to be (0,0).

The rest of the input is information pertaining to the rovers that have been deployed. Each rover has two lines of input. The first line gives the rover's position, and the second line is a series of instructions telling the rover how to explore the plateau.

The position is made up of two integers and a letter separated by spaces, corresponding to the x and y co-ordinates and the rover's orientation.

Below is the example of input that you can use:

```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

And you will get the output like this:
```
1 3 N
5 1 E
```

## Testing
### Run all tests
``` bash
go test -v ./...
```

Or using Make:
``` bash
make test
```

### Run tests with coverage
``` bash
make coverage
```

