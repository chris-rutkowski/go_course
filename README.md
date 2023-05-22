https://www.udemy.com/course/go-the-complete-developers-guide

playground https://go.dev/play/

projects:

- cards: some data structures with methods and unit tests
- channels: parallel code
- evenodd: trivial app
- fileprinter: accepts command line argument and prints the file with error handling
- helloworld: trivial app
- http: prints the body of the URL
- interfaces: demonstrates interfaces with various examples
- maps: equivalent of dictionary/hashmap
- playground: ignore
- structs demonstrates data structures

structs (pointers) slices are data structure that holds pointer to underlying array, that's why we can pass slices without a pointer (it creates a copy of slice, but that copy has the same array pointer)
similar other reference types: maps, channels, pointers, functions

`go run main.go`

```
go build main.go 
./main.go
```
