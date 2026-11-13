# r

Create simple shortcuts for R commands

## Install

Requires [Go](https://go.dev/dl/) 1.26 or later, or build the [POSIX-only](https://github.com/devOpifex/rc)

```bash
go install github.com/devOpifex/r@latest
```

Make sure `$(go env GOPATH)/bin` is on your `PATH`.

## Example

Config file `~/.rconfig`:

```json
[
  {
    "name": "d",
    "command": "devtools::document()"
  },
  {
    "name": "i",
    "command": "devtools::install()"
  },
  {
    "name": "c",
    "command": "devtools::check()"
  }
]
```

Usage: `r d c i`

Place any number of commands in the config file and run `r <command>`.

## FAQ 

- Commands are executed in order.
- Exits if command errors.
