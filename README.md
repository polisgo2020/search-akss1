Usage examples

For run http search server:
```bash
go run main.go server --host 127.0.0.1 --port 8080 --index idx.json
```

After the index is readed, and the server starts, we need to send the http GET request with "search" method and "str" query param":
```bash
    curl 127.0.0.1:8080/search?str=Author+help
or
    http://127.0.0.1:8080/search?str=Author help
```
<br>
<br>

For make reverse index:
```bash
go run main.go make --dir 'PATH_TO_FILES_DIR' --index 'OUTPUT.json'

Example: go run main.go make --dir books/backup --index idx.json
```

For search:
```bash
go run main.go -search --str 'STR_TOKENS_TO_FIND' --index 'PATH_TO_REVERSE_IDX.json' 

Example: go run main.go search --str "Author help" --index idx.json
``` 