# Logger

Logger is an opinionated middleware for logging incoming HTTP requests.

It is expected to be used in conjuction with [Connector](https://github.com/attheapplab/connector-go).

## Example

```go
package main

import (
	"github.com/attheapplab/connector-go"
	"github.com/attheapplab/logger-go"
	"net/http"
)

func main() {
	// Create a new instance of Connector handler.
	handler := connector.New()

	// Create your logger middleware instance and place it on any route.
	logger := logger.New()
	handler.Handle(http.MethodGet, "home", logger)
	
	// Start the server.
	handler.ListenAndServe()
}

```

## Contributing
For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)