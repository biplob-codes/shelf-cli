# shelf

A CLI tool for saving, organizing, and retrieving links — without ever leaving your terminal.

`shelf` lets you group links into collections, tag them for quick filtering, and manage everything through a fast, local SQLite-backed command line interface. Built as a hands-on project for learning idiomatic Go, Cobra, and clean CLI architecture.

```
$ shelf link add https://go.dev/doc/effective_go -t golang -c reading-list
✓ Saved https://go.dev/doc/effective_go

$ shelf link list -c reading-list
ID  URL                                    TAG      CREATED
1   https://go.dev/doc/effective_go        golang   2026-07-03
```

## Features

- **Collections** — group related links under named collections (`work`, `reading-list`, etc.)
- **Tags** — attach a single tag to any link for lightweight filtering
- **Local & fast** — everything is stored in a local SQLite database, no network calls, no accounts
- **Clean CLI output** — styled success/error/info messages instead of raw stack traces
- **Safe by default** — operations validate that the collection or link you're referencing actually exists before touching the database

## Tech Stack

| Layer         | Choice                                                                                   |
| ------------- | ---------------------------------------------------------------------------------------- |
| Language      | Go                                                                                       |
| CLI framework | [Cobra](https://github.com/spf13/cobra)                                                  |
| Database      | SQLite via [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) (pure Go, no CGO) |
| Architecture  | Repository pattern, one repository per domain entity                                     |

## Installation

**Requires Go 1.21+**

```bash
git clone https://github.com/biplob-codes/shelf-cli.git
cd shelf-cli
go build -o shelf .
```

Move the binary onto your `$PATH` to use it from anywhere:

```bash
mv shelf /usr/local/bin/
```

On first run, `shelf` creates and migrates a local `shelf.db` SQLite file in the current directory.

## Usage

### Collections

```bash
shelf collection create work            # create a new collection
shelf collection list                   # list all collections
shelf collection update work archive    # rename "work" to "archive"
shelf collection delete archive         # delete a collection
```

Alias: `shelf c ...`

### Links

```bash
shelf link add https://example.com -t news -c reading-list   # save a link
shelf link list -c reading-list                               # list links in a collection
shelf link list                                                # list uncategorized links
shelf link update 3 tech-articles                              # retag a link by id
shelf link delete 3                                             # delete a link by id
```

Aliases: `shelf lnk ...`, `shelf l ...`

Run `shelf --help` or `shelf <command> --help` for full flag details on any command.

## Project Structure

```
shelf-cli/
├── main.go                  # wiring: db connection, migrations, repo + command setup
├── cmd/                     # Cobra command definitions
│   ├── root.go
│   ├── collection.go
│   └── link.go
├── internal/
│   ├── db/                  # connection + migrations
│   ├── store/                # data access layer
│   │   ├── errors.go          # shared sentinel errors (ErrNotFound)
│   │   ├── collection.go      # CollectionRepository
│   │   └── link.go            # LinkRepository
│   └── ui/                   # terminal output styling & rendering
└── go.mod
```

## Architecture Notes

- **Repository pattern, split by domain.** Rather than one monolithic `Repository` struct handling every table, `CollectionRepository` and `LinkRepository` each own their table's persistence logic. `LinkRepository` holds a reference to `CollectionRepository` and reuses its lookup logic when it needs to validate a collection exists — the dependency between the two entities is explicit in the constructor signature, not implicit in duplicated SQL.
- **Errors propagate, they don't `Fatal` from inside commands.** Every Cobra command uses `RunE` instead of `Run`, returning wrapped errors (`fmt.Errorf("...: %w", err)`) up through `cmd.RootCMD.Execute()` and back to `main`, which is the single place that prints the final error and exits. `SilenceErrors` and `SilenceUsage` are set on the root command so Cobra doesn't double-print.
- **No silent failures.** Every write operation checks and validates its result — a rename or delete that affects zero rows returns a wrapped `store.ErrNotFound` instead of exiting quietly as if it succeeded, and referencing a nonexistent collection when adding a link fails loudly rather than inserting a `NULL` foreign key.

## Roadmap

- [ ] Full-text search across saved links
- [ ] Support multiple tags per link
- [ ] Open a link directly from the terminal (`shelf link open <id>`)
- [ ] Export/import collections as JSON

## License

MIT

## Author

Built by [Biplob](https://github.com/biplob-codes) as a learning project for backend/systems-focused Go development — CLI design, package architecture, and idiomatic error handling.
