# go-excel-picker

Pick cell values from Excel file.

## Usage

`-file` and `-cell` arguments are required.

Print picked values as JSON format.

```
$ excel-picker -file=test.xlsx -cell=Sheet1/B2
{"Sheet1/B2":"a","_filename":"test.xlsx"}
```

Use active sheet when sheet name is not set.

```
$ ./excel-picker -file=test.xlsx -cell=B2
{"B2":"a","_filename":"test.xlsx"}
```

Pick multiple values by multiple `-cell` arguments.

```
$ ./excel-picker -file=test.xlsx -cell=B2 -cell=D3 -cell=B4 -cell=D6
{"B2":"a","B4":"1","D3":"あ","D6":"aa\tbb\tcc\n1\n22","_filename":"test.xlsx"}
```

Pick values from multiple file by multiple `-file` arguments.

```
$ ./excel-picker -file=test.xlsx -file=test2.xlsx -cell=B2 -cell=D3 -cell=B4 -cell=D6
{"B2":"a","B4":"1","D3":"あ","D6":"aa\tbb\tcc\n1\n22","_filename":"test.xlsx"}
{"B2":"a","B4":"1","D3":"あ","D6":"aa\tbb\tcc\n1\n22","_filename":"test2.xlsx"}
```

Print picked values as CSV format by `-type=csv` argument.

```
$ ./excel-picker -file=test.xlsx -cell=B2  -type=csv
_filename,B2
test.xlsx,a
```

## License

MIT License
