### [WIP] this is my blog's api version2 (distribution version)

### Response type
```go
type Responee struct {
    Code    int         `json:"code"`
    Msg     string      `json:"msg"`
    Data    interface{} `json:"data"` 
}
```

the code `666` to represent success 
and `777` to represent failed