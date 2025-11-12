# AGENTS.md – Guidelines for taskwarrior-notes
Build: `make build` (binary `twn`), run: `make run`, install: `make install`.
Hook install: `make install_hook` (copies shell hook to `~/.task/hooks/`).
Tests: none yet; add `_test.go` then run `go test ./...`.
Single test: `go test -run 'TestName$' ./path/to/pkg`.
Benchmark: `go test -bench . ./path/to/pkg`.
Lint/format: `go fmt ./...`; vet: `go vet ./...`; optional: `staticcheck ./...`.
Imports: group stdlib, blank line, third-party, blank line, internal `taskwarrior-notes/...`.
Avoid unused imports; run `go mod tidy` to clean modules.
Types: exported structs CamelCase; JSON tags lowercase; use meaningful field names.
Constants: use CamelCase; only all-caps for stable external formats (e.g. `TW_TIME_FORMAT`).
Errors: return `error` not panic; wrap with `fmt.Errorf("context: %w", err)`; CLI commands print user-facing errors then `return`.
Fatal config issues may write to `os.Stderr` and `os.Exit(1)` (see `config.ReadNotesRoot`).
Do not log in `tw` or `util` packages beyond returning errors; keep them pure.
Naming: short but clear; avoid abbreviations except common (`cmd`, `tw`); file names lowercase.
Functions with side effects start with verb (`GetTasks`, `ReadNotesRoot`).
Avoid global state; use function parameters; only `rootCmd` as Cobra pattern.
External commands (`task`, `frontmatter`, `grep`) use `exec.Command`; capture and include output in error.
Sanitize user input (filenames) by replacing forbidden runes (see `TaskNotePath`).
Before adding new tools or tests, update this file; keep ~20 lines total.
