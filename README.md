# tour
My solutions to A Tour of Go exercises

## Compile and Run

### Exercise: Stringers

```
go run stringer.go
```

### Exercise: Readers

```sh
go run reader.go
```

Had to first run:

* `go mod init github.com/dskecse/tour` - created `go.mod`
* `go get golang.org/x/tour/reader` - created `go.sum` and appended to `go.mod`

https://www.digitalocean.com/community/tutorials/importing-packages-in-go

### Exercise: rot13Reader

```sh
go run rot13.go
```

### Exercise: Images

```sh
go run images.go
```

NOTE: The `pic` package is already available as part of the

```
require golang.org/x/tour v0.1.0 // indirect
```

in `go.mod`

### Exercise: Generic Types

```sh
go run generics.go
```

https://gobyexample.com/generics
