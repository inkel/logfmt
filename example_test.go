package logfmt_test

import (
	"os"

	"github.com/inkel/logfmt"
)

func Example() {
	l := logfmt.NewLogger(os.Stdout)

	l.Log("just a string", nil)
	l.Logf("%s %d", nil, "Hello", 2022)

	l.Log("a string with labels", logfmt.Labels{
		"lorem": "ipsum",
		"int":   1234,
	})

	// Output:
	// ts=1978-07-16T05:55:00Z msg="just a string"
	// ts=1978-07-16T05:55:00Z msg="Hello 2022"
	// ts=1978-07-16T05:55:00Z msg="a string with labels" int=1234 lorem="ipsum"
}
