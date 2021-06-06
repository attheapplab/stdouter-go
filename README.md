# Stdouter

Stdouter is an opinionated middleware for logging incoming HTTP requests.

It may be used in conjuction with [Connector](https://github.com/attheapplab/connector-go).

## Example

```golang
package main

import (
	"github.com/attheapplab/connector-go"
	"github.com/attheapplab/stdouter-go"
)

func main() {
	// Create a new instance of Connector handler.
	handler := connector.New()

	// Create your stdouter middleware instance and place it on any route.
	stdouter := stdouter.New()
	handler.Get("home", stdouter)
	
	// Start the server.
	handler.ListenAndServe()
}

```

## License
[MIT](https://choosealicense.com/licenses/mit/)