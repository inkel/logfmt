package logfmt

import (
	"errors"
	"fmt"
	"io"
	"testing"
	"time"
)

func init() {
	now = func() time.Time {
		return time.Date(1978, time.July, 16, 2, 55, 00, 00, time.FixedZone("AR", -3*3600)).UTC()
	}
}

func BenchmarkLoggerLog(b *testing.B) {
	l := NewLogger(io.Discard)

	labels := make(Labels, len(formatTests))
	for k := range formatTests {
		labels[fmt.Sprintf("%T", k)] = k
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Log("lorem ipsum dolor sit amet", labels)
	}
}

func errval(err error) error { return err }

var formatTests = map[any]string{
	uint8(126):                   "126",
	uint16(65535):                "65535",
	uint32(4294967295):           "4294967295",
	uint64(18446744073709551615): "18446744073709551615",

	int8(-128):                  "-128",
	int8(127):                   "127",
	int16(-32768):               "-32768",
	int16(32767):                "32767",
	int32(-2147483648):          "-2147483648",
	int32(2147483647):           "2147483647",
	int64(-9223372036854775808): "-9223372036854775808",
	int64(9223372036854775807):  "9223372036854775807",

	float32(3.14159265359):           "3.14159",
	float32(-3.14159265359):          "-3.14159",
	float64(2.71828182845904523536):  "2.71828",
	float64(-2.71828182845904523536): "-2.71828",

	time.Date(1978, time.July, 16, 2, 55, 00, 00, time.FixedZone("AR", -3*3600)): "1978-07-16T05:55:00Z",

	errval(errors.New("something failed")): `"something failed"`,
	errval(nil):                            "<nil>",
}

func TestFormat(t *testing.T) {
	for v, exp := range formatTests {
		if got := Format(v); exp != got {
			t.Errorf("%T :: expecting %q, got %q", v, exp, got)
		}
	}
}

var S string

func BenchmarkFormat(b *testing.B) {
	for v := range formatTests {
		var s string

		b.Run(fmt.Sprintf("%T", v), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				s = Format(v)
			}
			S = s
		})
	}
}
