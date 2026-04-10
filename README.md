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

## FAQ 

- Commands are executed in order.
- Exists if command errors.
