# Fullname parser

## Description

ParseFullname() is designed to parse names with a single first name, presupposing no middle / suffix / title

parseFullName():

1. accepts a string containing a person's full name
2. analyzes and attempts to detect the format of that name,
3. (if possible) parses the name into its component parts, and
4. returns an object containing first  & last name:
    - first (string): first name or initial
    - last (string): last name or initial

## Use

### Instalation
```
	go get github.com/lucasgabrielson/fullname_parser
```

### Basic Use

```go
package main
import fp "github.com/amonsat/fullname_parser"

func main() {
    parsedFullname := fp.ParseFullname("Mr. David Davis")
    println(parsedFullname.First)
    println(parsedFullname.Last)
}
```

### Parsedname struct

```go
type ParsedName struct {
	First  string
	Last   string
}
```

## Credits and precursors

My thanks to David Schnell-Davis & amonsat for sharing their work.

David Schnell-Davis's parse-full-name
https://github.com/dschnelldavis/parse-full-name

amonsat's fullname_parser
https://github.com/amonsat/fullname_parser
