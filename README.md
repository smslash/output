# output

Second optional task of ASCII-ART. File handling

## Author

* [@smslash](https://github.com/smslash)

## Objectives

- You must follow the same [instructions](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/README.md) as in the first subject while writing the result into a file.
The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name which will contain the output.

- The flag must have exactly the same format as above, any other formats must return the following usage message:

```bash
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.

Additionally, the program must still be able to run with a single `[STRING]` argument.

## Instructions

- Your project must be written in **Go**.

- The code must respect the [good practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).


## Usage

```
$ go run . --output=banner.txt "hello" standard
$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run . --output=banner.txt "Hello There!" shadow
$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$
```

## Allowed packages

- Only the [standard Go](https://pkg.go.dev/std) packages are allowed
This project will help you learn about :

- The Go file system(**fs**) API

- Data manipulation