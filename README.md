# goCleaner

### SimpleUse

```bash
go run .\main.go -root ".." -target "node_modules"
```

-help -> shows you the flags

```bash
go run .\main.go -help
```

-root -> use a string to set the initial path for search.

-target -> a unique string with one or multiple values to search, example

```bash
go run .\main.go -root ".." -target "node_modules"
```

this delete all node_modules it find,
actually you can use multiple values like this

```bash
go run .\main.go -root ".." -target "node_modules,index.js"
```
