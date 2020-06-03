# Logre Go Client
> [Logre][logre] client written in GoLang

This *logre* package provides a fast way to push your logs into [logre.io][logre]

> Are you a Java developer? Check out the [Java driver](https://github.com/Rawnly/logre-client-java) for Logre.IO

## Installation
```shell script
  go get -u github.com/rawnly/go-logre
``` 

## Getting Started
The easiest way to log is with the `Message` or `Log` method.

```go
package main
    
import "github.com/rawnly/go-logre"

func main() {
    type Payload struct {
        ID uint `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
    }

    // initialize the client with your box_id
    logre.Init("<Your box_id>")

    // Explicity ignoring the error
    _, _ = logre.Message("Hello World!")
    
    // You can also log a JSON
    _, _ = logre.Log(Payload{ ID: 1, Name: "Rawnly" })
}
```

Since this library is just a wrapper for an API call, we'll need to handle http-errors. 

## Log with Severity
You can also log with `severity` via `Info`, `Debug`, `Warning`, `Error` and `Fatal` methods 
 
```go
package main
    
import "github.com/rawnly/go-logre"

func main() {
    type Payload struct {
        ID uint `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
    }

    // initialize the client with your box_id
    logre.Init("<Your box_id>")

    // Explicity ignoring the error
    _, _ = logre.DebugMessage("Hello at Debug level")
    
    // You can also log a JSON
    _, _ = logre.Debug(Payload{ ID: 1, Name: "Rawnly" })
}
```
 [logre]: https://logre.io
