# TaskWarrior Notes CLI

TaskWarrior Notes CLI (`twn`) is a command-line tool for managing notes associated with TaskWarrior tasks. It allows you to synchronize task metadata with notes, retrieve note paths, and perform other note-related operations for tasks.

## Installation

Requirements:

- Go 1.24+

To build and install the tool:

```sh
make install
```

Or manually:

```sh
go build -o twn .
cp twn $(go env GOPATH)/bin/
```

## Building

To build the binary:

```sh
make build
```

To run directly:

```sh
make run
```

## License

MIT
