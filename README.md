# Sudoku solver

This is yet another sudoku solver that uses a backtracking algorithm
to get the solution.

## Usage

### Files

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

If the delimiter character is other than a comma, just use the flag -d. Example:

```bash
./sudoku-solver -f example2.csv -d " "
```

### Standard input

If no flags are passed, the program will read the standard input expecting to find 9
lines with 9 characters each.

By default, the characters should be separated with a space but this can be changed with
the flag -d.

Example:

```bash
$ ./sudoku-solver
Please introduce the sudoku to be solved using " " as separator:

6 . . . . . 1 9 3
. . . . 3 . . . .
. . . . . . . . 6
. . 7 3 . . . . .
. . 6 2 . 1 7 8 .
. 1 . . 6 . . 4 .
2 7 . . . 4 . . .
. . . 9 . . . 2 .
3 . . . 8 . . . 9

Sudoku to be solved:
╔═══════════════════════╗
║ 6 0 0 ║ 0 0 0 ║ 1 9 3 ║
║ 0 0 0 ║ 0 3 0 ║ 0 0 0 ║
║ 0 0 0 ║ 0 0 0 ║ 0 0 6 ║
║═══════O═══════O═══════║
║ 0 0 7 ║ 3 0 0 ║ 0 0 0 ║
║ 0 0 6 ║ 2 0 1 ║ 7 8 0 ║
║ 0 1 0 ║ 0 6 0 ║ 0 4 0 ║
║═══════O═══════O═══════║
║ 2 7 0 ║ 0 0 4 ║ 0 0 0 ║
║ 0 0 0 ║ 9 0 0 ║ 0 2 0 ║
║ 3 0 0 ║ 0 8 0 ║ 0 0 9 ║
╚═══════════════════════╝

Solution:
╔═══════════════════════╗
║ 6 8 2 ║ 4 7 5 ║ 1 9 3 ║
║ 1 9 5 ║ 8 3 6 ║ 2 7 4 ║
║ 7 4 3 ║ 1 2 9 ║ 8 5 6 ║
║═══════O═══════O═══════║
║ 5 2 7 ║ 3 4 8 ║ 9 6 1 ║
║ 4 3 6 ║ 2 9 1 ║ 7 8 5 ║
║ 9 1 8 ║ 5 6 7 ║ 3 4 2 ║
║═══════O═══════O═══════║
║ 2 7 9 ║ 6 1 4 ║ 5 3 8 ║
║ 8 6 1 ║ 9 5 3 ║ 4 2 7 ║
║ 3 5 4 ║ 7 8 2 ║ 6 1 9 ║
╚═══════════════════════╝
It took 75.9684ms and 482463 iterations to solve the sudoku
```

### Samples

A set of 3 sample sudokus are pre-loaded in the program and can be selected
with the flag -s.

The sample "hardest" is interesting because it's a sudoku especifically designed against backtracking
algorithms. While a regular sudoku takes between micro and milliseconds to solve, this one takes some
seonds.

### Verbose mode

A flag -v (verbose) can be used to see how the algorithm is working.

Since the call to print a message on the screen is quite slow, it's not really
useful when the number of iterations is high.

## Limitations

So far the program has been tested only in Linux. It's possible that the
output is not shown correctly in Windows systems when verbose mode is enabled.
