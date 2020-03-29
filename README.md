Usage examples

For make reverse index:
```bash
go run main.go -m=make -in='PATH_TO_FILES_DIR' -out='OUTPUT.json'

Example: go run main.go -m=make -in=/techno/search-akss1/books/ -out=idx.json
```

For search:
```bash
go run main.go -m=search -in='PATH_TO_REVERSE_IDX.json' -str='STR_TOKENS_TO_FIND'

Example: go run main.go -m=search -in=/techno/search-akss1/idx.json -str="Hello friend"
```