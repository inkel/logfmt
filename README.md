# logfmt - An opinionated logging library in logfmt for Go
[![Go Reference](https://pkg.go.dev/badge/github.com/inkel/logfmt.svg)](https://pkg.go.dev/github.com/inkel/logfmt)

How opinionated?

* The only possible output format is [logfmt](https://brandur.org/logfmt).
* All string values are quoted.
* Label keys are sorted except `ts` for timestamp and `msg` for the message are sorted.
* [`time.Time`](https://pkg.go.dev/time#Time) values are in UTC and formatted using [time.RFC3339](https://pkg.go.dev/time#pkg-constants).

Another strong opinion is that it simply writes to an [`io.Writer`](https://pkg.go.dev/io#Writer).
You get to choose if this writer is buffered, a file, a network connection, a terminal.
It's not something the package should worry about.

## Usage
```go
package main

import (
	"os"
	"time"

	"github.com/inkel/logfmt"
)

func main() {
	l := logfmt.NewLogger(os.Stdout)

	l.Log("a simple log message without any labels", nil)
	l.Logf("a formatted log %s without any labels", nil, "message")

	l.Log("message with a few labels of different types", logfmt.Labels{
		"foo": "bar",
		"err": os.ErrPermission,
		"now": time.Now(),
	})
}
```

This will produce an output like the following:

```
ts=2022-07-13T12:31:41Z msg="a simple log message without any labels"
ts=2022-07-13T12:31:41Z msg="a formatted log message without any labels"
ts=2022-07-13T12:31:41Z msg="message with a few labels of different types" err="permission denied" foo="bar" now=2022-07-13T12:31:41Z
```

## License
MIT. See [LICENSE](LICENSE).
