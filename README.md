# Logger

Logger is an opinionated middleware for logging incoming HTTP requests.

It may be used in conjuction with [Connector](https://github.com/attheapplab/connector-go).

## Example

```golang
package main

import (
	"github.com/attheapplab/connector-go"
	"github.com/attheapplab/logger-go"
)

func main() {
	// Create a new instance of Connector handler.
	handler := connector.New()

	// Create your logger middleware instance and place it on any route.
	logger := logger.New()
	handler.Get("home", logger)
	
	// Start the server.
	handler.ListenAndServe()
}

```

## License
[MIT](https://choosealicense.com/licenses/mit/)