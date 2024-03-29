[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/rconfig)](https://goreportcard.com/report/github.com/Luzifer/rconfig)
[![Documentation](https://badges.fyi/static/godoc/reference/5272B4)](https://pkg.go.dev/github.com/Luzifer/rconfig/v2)
![](https://badges.fyi/github/license/Luzifer/rconfig)
![](https://badges.fyi/github/latest-tag/Luzifer/rconfig)

## Description

> Package rconfig implements a CLI configuration reader with struct-embedded defaults, environment variables and posix compatible flag parsing using the [pflag](https://github.com/spf13/pflag) library.

## Installation

Install by running:

```
go get -u github.com/Luzifer/rconfig/v2
```

Run tests by running in the checkout folder:

```
go test -v -race -cover ./...
```

## Usage

A very simple usecase is to just configure a struct inside the vars section of your `main.go` and to parse the commandline flags from the `main()` function:

```go
package main

import (
  "fmt"
  "github.com/Luzifer/rconfig/v2"
)

var (
  cfg = struct {
    Username string `default:"unknown" flag:"user" description:"Your name"`
    Details  struct {
      Age int `default:"25" flag:"age" env:"age" description:"Your age"`
    }
  }{}
)

func main() {
  rconfig.Parse(&cfg)

  fmt.Printf("Hello %s, happy birthday for your %dth birthday.",
    cfg.Username,
    cfg.Details.Age)
}
```

### Provide variable defaults by using a file

Given you have a file `~/.myapp.yml` containing some secrets or usernames (for the example below username is assumed to be "luzifer") as a default configuration for your application you can use this source code to load the defaults from that file using the `vardefault` tag in your configuration struct.

The order of the directives (lower number = higher precedence):

1. Flags provided in command line
1. Environment variables
1. Variable defaults (`vardefault` tag in the struct)
1. `default` tag in the struct

```go
var cfg = struct {
  Username string `vardefault:"username" flag:"username" description:"Your username"`
}

func main() {
  rconfig.SetVariableDefaults(rconfig.VarDefaultsFromYAMLFile("~/.myapp.yml"))
  rconfig.Parse(&cfg)

  fmt.Printf("Username = %s", cfg.Username)
  // Output: Username = luzifer
}
```

## More info

You can see the full reference documentation of the rconfig package [at pkg.go.dev](https://pkg.go.dev/github.com/Luzifer/rconfig/v2)
