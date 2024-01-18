# celi

celi is a simple CEL interpreter for CLI.

## Usage

celi reads a CEL query from stdin, compiles it, optionally checks it, executes
it, then prints the resulting type to stdout. Run `celi -help` for available
flags.

## Installation

Currently no binaries are provided and the only option is to build from source.
```
go install github.com/fore5fire/celi
```

If anyone ever cares to actually use this, open an issue and I'll do a real
release with binaries for popular platforms.
