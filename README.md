# Sudoku solver

This is yet another sudoku solver that uses a backtracking algorithm
to get the solution.

## Usage

### Files

The sudokus can be passed as files with the flag -f. Example:

```
$ ./sudoku-solver -f examples/example3.txt
Parsing file using " " as delimiter...

Sudoku to be solved:
╔═══════════════════════╗
║ 0 1 0 | 2 0 4 | 0 8 0 ║
║ 0 0 4 | 0 0 0 | 0 0 0 ║
║ 9 0 2 | 0 0 8 | 0 0 0 ║
║-------+-------+-------║
║ 0 3 0 | 7 0 0 | 0 4 0 ║
║ 0 0 0 | 0 6 0 | 0 0 0 ║
║ 0 9 0 | 0 2 0 | 6 0 7 ║
║-------+-------+-------║
║ 0 7 0 | 9 0 0 | 2 6 0 ║
║ 3 4 0 | 0 0 0 | 7 0 1 ║
║ 0 0 0 | 0 0 0 | 0 0 9 ║
╚═══════════════════════╝

Solution:
╔═══════════════════════╗
║ 7 1 3 | 2 9 4 | 5 8 6 ║
║ 8 5 4 | 3 7 6 | 1 9 2 ║
║ 9 6 2 | 1 5 8 | 4 7 3 ║
║-------+-------+-------║
║ 2 3 6 | 7 1 5 | 9 4 8 ║
║ 1 8 7 | 4 6 9 | 3 2 5 ║
║ 4 9 5 | 8 2 3 | 6 1 7 ║
║-------+-------+-------║
║ 5 7 8 | 9 3 1 | 2 6 4 ║
║ 3 4 9 | 6 8 2 | 7 5 1 ║
║ 6 2 1 | 5 4 7 | 8 3 9 ║
╚═══════════════════════╝
It took 18.0414ms and 80055 iterations to solve the sudoku
```

The default delimiter is a space, but this can be easily changed using the flag -d. Example:

```bash
./sudoku-solver -f examples/example1.txt -d ","
```

### Standard input

If no flags are passed, the program will read the standard input expecting to find 9
lines with 9 characters each.

```
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
║ 6 0 0 | 0 0 0 | 1 9 3 ║
║ 0 0 0 | 0 3 0 | 0 0 0 ║
║ 0 0 0 | 0 0 0 | 0 0 6 ║
║-------+-------+-------║
║ 0 0 7 | 3 0 0 | 0 0 0 ║
║ 0 0 6 | 2 0 1 | 7 8 0 ║
║ 0 1 0 | 0 6 0 | 0 4 0 ║
║-------+-------+-------║
║ 2 7 0 | 0 0 4 | 0 0 0 ║
║ 0 0 0 | 9 0 0 | 0 2 0 ║
║ 3 0 0 | 0 8 0 | 0 0 9 ║
╚═══════════════════════╝

Solution:
╔═══════════════════════╗
║ 6 8 2 | 4 7 5 | 1 9 3 ║
║ 1 9 5 | 8 3 6 | 2 7 4 ║
║ 7 4 3 | 1 2 9 | 8 5 6 ║
║-------+-------+-------║
║ 5 2 7 | 3 4 8 | 9 6 1 ║
║ 4 3 6 | 2 9 1 | 7 8 5 ║
║ 9 1 8 | 5 6 7 | 3 4 2 ║
║-------+-------+-------║
║ 2 7 9 | 6 1 4 | 5 3 8 ║
║ 8 6 1 | 9 5 3 | 4 2 7 ║
║ 3 5 4 | 7 8 2 | 6 1 9 ║
╚═══════════════════════╝
It took 78.3289ms and 482463 iterations to solve the sudoku
```

Note that the blank cells can be represented by any character other than the separator itself.

### Samples

A set of 3 sample sudokus are pre-loaded in the program and can be selected
with the flag -s.

The sample "hardest" is interesting because it's a sudoku especifically designed against backtracking
algorithms. While a regular sudoku takes between micro and milliseconds to solve, this one takes some
seonds.

### Verbose mode

A flag -v (verbose) can be used to see how the algorithm is working. By default, the output will be updated
120 times per second, although this can be changed using the flag -r.

Note that in any case, the verbose mode will slow down the solution significantly.

## Limitations

So far the program has been tested only in Linux. It's possible that the
output is not shown correctly in Windows systems when verbose mode is enabled.
