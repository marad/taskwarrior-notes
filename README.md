# TaskWarrior Notes CLI

TaskWarrior Notes CLI (`twn`) is a command-line tool for managing notes associated with TaskWarrior tasks. It allows you to synchronize task metadata with notes, retrieve note paths, and perform other note-related operations for tasks.

## Installation

Requirements:

- Go 1.24+
- Installed [frontmatter](https://github.com/marad/frontmatter)
- Installed [`yq`](https://github.com/mikefarah/yq)

To build and install the tool:

```sh
make install
```

Or manually:

```sh
go build -o twn .
cp twn $(go env GOPATH)/bin/
```

### Installing the TaskWarrior Hook

To automatically sync notes when tasks are modified, install the provided TaskWarrior hook:

```sh
make install_hook
```

This copies `hooks/on-modify-sync-note.sh` to your `~/.task/hooks/` directory. The hook will call `twn` to sync notes whenever a task is modified.

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
