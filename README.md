# identifier
Very fast generator of unique IDs.

- 8 bytes decoded integer
  - int64 (only 63 bits used)
  - contains a unix timestamp with millisecond precision
  - sequential
  - sortable by time
  - used in databases
- 13 bytes encoded string
  - appears random
  - used in JSON

## Install
```sh
go get github.com/webmafia/identifier
```

## Usage
```
id := identifier.Generate()
```