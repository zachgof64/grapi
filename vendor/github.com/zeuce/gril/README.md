# Gril

A GoLang library full of some random functions

## Installation

RUN:

```
go get github.com/zeuce/gril
```

## Usage

- [Check Function](#check-function)
- [GetByte Function](#getbyte-function)

### Check Function

```
package main


import (
    "github.com/zeuce/gril"
)

func main() {
    err := someFunctionThatReturnsError()
    gril.Check(err)
}
```

### GetByte Function

```
package main


import (
    "github.com/zeuce/gril"
)

func main() {
    gril.getByte("path/to/file")
}
```
