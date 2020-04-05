Usage examples

For make reverse index:
```bash
go run main.go make --dir='PATH_TO_FILES_DIR' --index='OUTPUT.json'

Example: go run main.go make --dir books/backup --index idx.json
```

For search:
```bash
go run main.go -search --str='STR_TOKENS_TO_FIND' --index='PATH_TO_REVERSE_IDX.json' 

Example: go run main.go search --str "Author help" --index idx.json
``` 