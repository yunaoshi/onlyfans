# onlyfans

Simple command-line tool for downloading user lists or media from OnlyFans.

## Features
- Download user lists and media from a logged-in OnlyFans account
- Structured logging using go.uber.org/zap (see `internal/logger`)
- Configurable log level with `--verbose` flag

## Usage

### Build the binary (optional)
```sh
go build -o bin/onlyfans
```

### Run the tool
You can run the compiled binary or use `go run` directly:

**Using the binary:**
```sh
bin/onlyfans --help
bin/onlyfans download --user <username>
```

**Using go run:**
```sh
go run . --help
go run . download --user <username>
```

### Common options
- `--verbose` : Enable verbose logging
- `--output <path>` : Specify output directory

See `--help` for all available commands and options.