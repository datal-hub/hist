# HIST

This program allows to count histogram of ascii symbols for files from the directory.

### Init test files

If you want to create files to test the program, you can run it with the following arguments:
```bash
./hist init --fc 10000 --sc 1000 --path ./examples
```

In this case 10000 files will be created in directory ./examples. 
Each file will contain 1000 random ascii symbols.

### Count histogram ascii symbols from directory

If you want to count histogram, you should run the program with the following arguments:
```bash
./hist hist --p 8 --a 100 --path ./examples
```

In this case the program will be launched with variable GOMAXPROCS equal 8.
Parameter _a_ is the number of aggregators. Aggregator is routine that aggregates results
of counting histogram for each file. It is recommended to set it less than the number of files.
_Path_ is path to directory with text files.

As a result, you will get array with 128 elements where each element corresponds to the number 
of repetitions of the corresponding ascii symbol.
