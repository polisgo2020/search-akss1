Usage examples

Before the start you need to set environment variables: <b>LISTEN</b> (default: "localhost:8080") and <b>LOG_LEVEL</b> (default: "info")


For run http search server:
```bash
go run main.go server --index idx.json
```

After the index is readed, and the server starts, we need to send the http GET request with "search" method and "str" query param":
```bash
    curl 127.0.0.1:8080/search?q=Author+help
or
    http://127.0.0.1:8080/search?q=Author help
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