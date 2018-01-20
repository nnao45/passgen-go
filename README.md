# passgen
Password generator written in Go.
## Usage
```bash
$ go run passgen.go -h
usage: passgen [<flags>]

A password generate in Go application.

Flags:
  -h, --help        Show context-sensitive help (also try --help-long and --help-man).
  -N, --noVaildNum  No valid Number in a generating password.
  -S, --noVaildSym  No valid Symbol in a generating password.
  -R, --noVaildUnr  No valid Unreadable character in a generating password.
  -n, --number=1    Number of generating password.
  -l, --length=31   Length of generating password.
```
