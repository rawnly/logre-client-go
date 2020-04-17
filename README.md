# Logre Go Client
> [Logre][logre] client written in GoLang

This *logre* package provides a fast way to push your logs into [logre.io][logre]

## Installation
```shell script
  go get -u github.com/rs/zerolog/log
``` 

## Getting Started
The easiest way to log is with the `Message` or `Log` method.

```go
package main
    
import "github.com/rawnly/logre"

func main() {
    var err error

    type Payload struct {
        ID uint `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
    }

    // initialize the client with your box_id
    logre.Init("<Your box_id>")

    // Explicity ignore the error
    _, err = logre.Message("Hello World!")
    // do something with error...
    
    // to log a JSON just
    _, err = logre.Log(Payload{ ID: 1, Name: "Rawnly" })
    // do something with error...
}
```

Since this library is just a wrapper for an API call, we'll need to handle http-errors. 

## Log with Severity
You can also log with `severity` via `Info`, `Debug`, `Warning`, `Error` and `Fatal` methods 
 
```go
package main
    
import "github.com/rawnly/logre"

func main() {
    var err error

    type Payload struct {
        ID uint `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
    }

    // initialize the client with your box_id
    logre.Init("<Your box_id>")
    
    _, err = logre.DebugMessage("Hello at Debug level")
    
    _, err = logre.Debug(Payload{ ID: 1, Name: "Rawnly" })
    // do something with error...
}
```
 [logre]: https://logre.io