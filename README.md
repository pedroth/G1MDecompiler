# G1M Decompiler
[Visit project page on humaidq.ae](https://humaidq.ae/projects/g1mdecompiler/)
## 1. Purpose
The goal of G1M Decompiler is to allow programmers to decode Casio Basic's "`.g1m`"
file into a readable text similarly what is shown on the calculator.

This is a one-way converter, and not all symbols are supported by this program (most
common symbols are).

## 2. Requirements

The following packages must be installed on your system

- Go
- Git

You also need a Casio Basic file to decode.

## 3. Copying and contributing

This program is written by Humaid AlQassimi, and is distributed under
the [MIT](https://humaidq.ae/license/mit) license.  

## 4. Download and install

```sh
$ go get -u git.sr.ht/~humaid/g1mdecompiler
$ go install git.sr.ht/~humaid/g1mdecompiler
```

## 5. Usage
The program outputs the decoded program to standard output.
```sh
$ g1mdecompiler [file] > [output]
```
