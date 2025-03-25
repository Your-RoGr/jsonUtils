![Version](https://img.shields.io/badge/Version-1.0.0-blue)
![Code Coverage](https://img.shields.io/codecov/c/github/Your-RoGr/jsonUtils)
![License](https://img.shields.io/github/license/Your-RoGr/jsonUtils)
![Downloads](https://img.shields.io/github/downloads/Your-RoGr/jsonUtils/total)
![Go Report Card](https://goreportcard.com/badge/Your-RoGr/jsonUtils)
![Latest Release](https://img.shields.io/github/v/release/Your-RoGr/jsonUtils)
![GitHub Stars](https://img.shields.io/github/stars/Your-RoGr/jsonUtils?style=social)

# jsonUtils

[English](README.md) | [Русский](README.ru.md)

This package provides utility functions for working with JSON in Go. It includes two main components: `FindAndParseJSON` and `numberedJSON`.

- [Features](#Features)
- [Usage](#Usage)
- [Dependencies](#Dependencies)
- [Installation](#Installation)
- [License](#License)

## Features

### FindAndParseJSON

The `FindAndParseJSON` function extracts and parses JSON objects from a given text string.

```go
func FindAndParseJSON(text string) (*[]map[string]interface{}, error)
```

- Uses regular expressions to find JSON-like structures in the text
- Attempts to parse each found structure as JSON
- Returns a slice of successfully parsed JSON objects as maps
- Skips invalid JSON structures
- Returns an error if no valid JSON is found

### numberedJSON

The `numberedJSON` struct and its associated methods provide a way to create numbered JSON objects from a slice of strings.

```go
type numberedJSON struct {
    items      *[]string
    key        *string
    generalKey *string
}
```

#### Methods

1. `GetString() (string, error)`
   - Converts the numbered JSON to a JSON string
   - Returns an error if items are nil

2. `GetMap() *map[string]map[string]string`
   - Converts the numbered JSON to a nested map structure
   - Returns an empty map if items are nil

3. `CreateNumberedJSON(items *[]string, key string, generalKey string) *numberedJSON`
   - Creates a new `numberedJSON` instance

## Usage

### FindAndParseJSON

```go
text := `Some text {{"key": "value"}} more text {"another": "object"}`
result, err := jsonUtils.FindAndParseJSON(text)
if err != nil {
    // Handle error
}
// Use result
```

### numberedJSON

```go
items := []string{"item1", "item2", "item3"}
nj := jsonUtils.CreateNumberedJSON(&items, "Item", "General")

// Get as string
jsonString, err := nj.GetString()
if err != nil {
    // Handle error
}

// Get as map
jsonMap := nj.GetMap()
```

### Error Handling

Both components include error handling for various scenarios:

- `FindAndParseJSON` returns errors for no JSON found or no valid JSON parsed
- `numberedJSON.GetString()` returns an error if items are nil

## Dependencies

This package uses only standard library packages:

- `encoding/json`
- `errors`
- `regexp`
- `fmt`

## Installation

To install the package, use the command:

```bash
go get github.com/Your-RoGr/jsonUtils
```

## License

jsonUtils is MIT-Licensed
