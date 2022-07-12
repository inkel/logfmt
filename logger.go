// logfmt is a small and opinionated logging package.
//
// The only possible output format is [logfmt], all string values are
// quoted, keys are sorted, [time.Time] values are in UTC and
// formatted using [time.RFC3339].
//
// [logfmt]: https://brandur.org/logfmt
package logfmt

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
	"time"
	"unicode"
)

// Logger is the basic and only type of logger.
type Logger struct {
	w io.Writer
}

// NewLogger creates a new [Logger].
func NewLogger(w io.Writer) *Logger {
	return &Logger{w}
}

var now func() time.Time = time.Now().UTC

// Log writes the given msg and labels to w.
func (l *Logger) Log(msg string, labels Labels) (int, error) {
	var buf bytes.Buffer

	buf.WriteString("ts=")
	buf.WriteString(now().Format(time.RFC3339))
	buf.WriteString(" msg=")
	buf.WriteString(strconv.Quote(msg))

	keys := make([]string, 0, len(labels))
	for k := range labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		buf.WriteRune(' ')
		buf.WriteString(k)
		buf.WriteRune('=')
		buf.WriteString(format(labels[k]))
	}
	buf.WriteByte('\n')

	return l.w.Write(buf.Bytes())
}

// Logf writes a formatted message and labels to w.  It is equivalent
// calling [Log] with a previously formatted string.
func (l *Logger) Logf(format string, labels Labels, v ...any) (int, error) {
	return l.Log(fmt.Sprintf(format, v...), labels)
}

func format(v any) string {
	switch x := v.(type) {
	case string:
		return strconv.Quote(x)
	case uint8:
		return strconv.FormatUint(uint64(x), 10)
	case uint16:
		return strconv.FormatUint(uint64(x), 10)
	case uint32:
		return strconv.FormatUint(uint64(x), 10)
	case uint64:
		return strconv.FormatUint(x, 10)

	case int8:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case int64:
		return strconv.FormatInt(x, 10)

	case float32:
		return strconv.FormatFloat(float64(x), 'g', 6, 32)
	case float64:
		return strconv.FormatFloat(x, 'g', 6, 64)

	case time.Time:
		return x.UTC().Format(time.RFC3339)

	case error:
		return strconv.Quote(x.Error())

	case fmt.Stringer:
		return strconv.Quote(x.String())

	default:
		s := fmt.Sprint(x)

		for _, r := range s {
			if unicode.IsSpace(r) {
				return strconv.Quote(s)
			}
		}

		return s
	}
}

// Labels is a custom type for mapping keys and values.
type Labels map[string]any
