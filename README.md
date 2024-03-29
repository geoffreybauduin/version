# Version

[![doc](https://godoc.org/github.com/geoffreybauduin/version?status.svg)](https://godoc.org/github.com/geoffreybauduin/version)
[![Build Status](https://travis-ci.org/geoffreybauduin/version.svg?branch=master)](https://travis-ci.org/geoffreybauduin/version)
[![Go Report Card](https://goreportcard.com/badge/github.com/geoffreybauduin/version)](https://goreportcard.com/report/github.com/geoffreybauduin/version)
[![Coverage Status](https://coveralls.io/repos/github/geoffreybauduin/version/badge.svg?branch=master)](https://coveralls.io/github/geoffreybauduin/version?branch=master)

Check your versions easily, and act accordingly

```golang
package main

import (
    "fmt"

    "github.com/geoffreybauduin/version"
)

func main() {
    v, errVersion := version.New("1.0.2")
    if errVersion != nil {
        panic(errVersion)
    }
    if v.Is("=1.0.2") {
        fmt.Printf("ok")
    }
    if v.Is(">1.0.0") {
        fmt.Printf("ok")
    }
}
```

## Comparators

- equality ; `=1.0`, `1.0`
- superior ; `>1.0`
- inferior ; `<1.0`
- equal or superior ; `>=1.0`
- equal of inferior ; `<=1.0`

## License

MIT License

Copyright (c) 2019 Geoffrey Bauduin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
