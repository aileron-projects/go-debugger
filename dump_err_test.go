package debugger_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/aileron-projects/go-debugger"
	"github.com/aileron-projects/go-tester"
)

func TestDumpErrAlwaysTo(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		errs  []error
		wants []string
	}{
		"1 error": {
			errs: []error{io.EOF},
			wants: []string{
				"Error: EOF",
				"(*errors.errorString)(EOF)",
			},
		},
		"2 errors": {
			errs: []error{io.EOF, io.ErrClosedPipe},
			wants: []string{
				"Error: EOF",
				"(*errors.errorString)(EOF)",
				"Error: io: read/write on closed pipe",
				"(*errors.errorString)(io: read/write on closed pipe)",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			debugger.DumpErrAlwaysTo(&buf, "", tc.errs...)
			result := buf.String()
			tester.AssertEqual(t, true, strings.Contains(result, "[DEBUGGER][DUMPERR]"))
			for c, w := range tc.wants {
				if !tester.AssertEqual(t, true, strings.Contains(result, w)) {
					t.Log(c, ": ", "`"+w+"`", "is not contained")
				}
			}
		})
	}
}
