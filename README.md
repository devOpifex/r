# r

Create simple shortcuts for R commands

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
