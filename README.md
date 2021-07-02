# Humantouch

[![GoDoc](https://pkg.go.dev/github.com/masv3971/humantouch?status.svg)](https://pkg.go.dev/github.com/masv3971/humantouch)

## Installation

```
go get -u github.com/masv3971/humantouch
```

## Example
```go
package main

import (
    "github.com/masv3971/humantouch"
)

func main() {
    person, _ := humantouch.New(nil)

    person.RandomHuman()
    // *Person
}
```