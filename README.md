# Sudoku solver

This is yet another sudoku solver that uses a backtracking algorithm
to get the solution.

## Usage

The sudokus can be passed as *.csv files with the flag -f. Example:

```bash
$ ./sudoku-solver -f sodoku1.csv

Sudoku to be solved:
╔═══════════════════════╗
║ 0 7 5 ║ 0 0 0 ║ 0 2 0 ║
║ 3 6 9 ║ 2 0 0 ║ 7 0 0 ║
║ 0 0 8 ║ 0 3 5 ║ 1 6 9 ║
║═══════O═══════O═══════║
║ 6 0 0 ║ 0 1 3 ║ 0 0 0 ║
║ 0 0 0 ║ 6 0 0 ║ 0 0 5 ║
║ 0 9 4 ║ 0 7 0 ║ 6 0 0 ║
║═══════O═══════O═══════║
║ 4 5 0 ║ 0 0 0 ║ 3 8 7 ║
║ 7 0 6 ║ 0 5 0 ║ 4 0 2 ║
║ 0 0 2 ║ 3 4 7 ║ 0 1 0 ║
╚═══════════════════════╝
Solution:
╔═══════════════════════╗
║ 1 7 5 ║ 4 6 9 ║ 8 2 3 ║
║ 3 6 9 ║ 2 8 1 ║ 7 5 4 ║
║ 2 4 8 ║ 7 3 5 ║ 1 6 9 ║
║═══════O═══════O═══════║
║ 6 2 7 ║ 5 1 3 ║ 9 4 8 ║
║ 8 1 3 ║ 6 9 4 ║ 2 7 5 ║
║ 5 9 4 ║ 8 7 2 ║ 6 3 1 ║
║═══════O═══════O═══════║
║ 4 5 1 ║ 9 2 6 ║ 3 8 7 ║
║ 7 3 6 ║ 1 5 8 ║ 4 9 2 ║
║ 9 8 2 ║ 3 4 7 ║ 5 1 6 ║
╚═══════════════════════╝
It took 65.8µs and 49 iterations to solve the sudoku
```

### Samples

A set of sample sudokus are pre-loaded in the program. If no flag is passed,
the program will run the easiest (lower iterations) of these samples.

Another sample can be chosen with the flag -s.

The hardest sample is a sudoku especifically designed against backtracking
algorithms, so its resolution can take some seconds.

### Verbose mode

A flag -v (verbose) can be used to see how the algorithm is working.

Since the call to print a message on the screen is quite slow, it's not really
useful when the number of iterations is high.

## Limitations

So far the program has been tested exclusively in Linux. Expect some problems
in Windows systems, especially when using the verbose mode.
