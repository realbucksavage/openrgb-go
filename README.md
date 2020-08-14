# Go OpenRGB Client

A Go client library for [OpenRGB](https://gitlab.com/CalcProgrammer1/OpenRGB) SDK Server.
This library is Go a port of [vlakreeh/openrgb](https://github.com/vlakreeh/openrgb), which is written in JavaScript.

## Usage

Here's how this library can be used to do some stuff. There's an example file in this repository 
([`examples/basic.go`](https://github.com/realbucksavage/openrgb-go/blob/master/examples/basic.go)) that 
sets the color of all connected devices to Cyan.

A connection to an OpenRGB Server can be opened like this:

```go
import "github.com/realbucksavage/openrgb-go"

func main() {
    c, err := openrgb.Connect("localhost", 6742)
    if err != nil {
        // ...
    }

    defer c.Close()
}
```

The `Connect(...)` method returns an instance of `openrgb.Client`. All available methods methods can be
found in the `client.go` file (I'll update the docs later, I promise :smile:).

## Open TODOs

- Documentation
- Save/Load profile methods
- I feel like `openrgb.Client.UpdateZoneLEDs` method can be written in a better way...

## Contributing

### What contributions are needed?

Anything that is an open to-do item, is a much welcome contribution.

### How to reach me?

These are some ways:

- Discord : bucksavage100x#0476 
- [Email](mailto:me@bsavage.xyz)
- Open a PR or an Issue on this project
