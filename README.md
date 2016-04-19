[![Build Status](https://travis-ci.org/petrkotek/go-vector.svg?branch=master)](https://travis-ci.org/petrkotek/go-vector)
[![Coverage Status](https://coveralls.io/repos/petrkotek/go-vector/badge.svg?branch=master&service=github)](https://coveralls.io/github/petrkotek/go-vector?branch=master)

# go-vector
Go (golang) library implementing vector math.

## Usage

### 1. Fetch the package 

```
go get github.com/petrkotek/go-vector
```

### 2. Import the package
Import the package into your `.go` file:

```go
package main

import (
        "fmt"

        "github.com/petrkotek/go-vector/2d/32bit/vec"
)

func main() {    
        v1 := vec.New(1, 2)
        v2 := vec.New(3, 4)
        fmt.Println(v1.Add(v2))
}
```
