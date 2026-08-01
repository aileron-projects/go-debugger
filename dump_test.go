package debugger_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/aileron-projects/go-debugger"
	"github.com/aileron-projects/go-tester"
)

func TestDumpAlwaysTo(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		a     []any
		wants []string
	}{
		"int": {
			a: []any{int(1), int32(2), int64(3)},
			wants: []string{
				`| (int) 1`,
				`| (int32) 2`,
				`| (int64) 3`,
			},
		},
		"uint": {
			a: []any{uint(1), uint32(2), uint64(3)},
			wants: []string{
				`| (uint) 1`,
				`| (uint32) 2`,
				`| (uint64) 3`,
			},
		},
		"float": {
			a: []any{float32(1.23), float64(4.56)},
			wants: []string{
				`| (float32) 1.23`,
				`| (float64) 4.56`,
			},
		},
		"bool": {
			a: []any{true, false},
			wants: []string{
				`| (bool) true`,
				`| (bool) false`,
			},
		},
		"complex": {
			a: []any{complex64(1 + 2i), complex128(3 + 4i)},
			wants: []string{
				`| (complex64) (1+2i)`,
				`| (complex128) (3+4i)`,
			},
		},
		"slice": {
			a: []any{[]int{1, 2}},
			wants: []string{
				`| ([]int) (len=2) {`,
				`|  (int) 1,`,
				`|  (int) 2`,
			},
		},
		"map": {
			a: []any{map[int]string{1: "a", 2: "b"}},
			wants: []string{
				`| (map[int]string) (len=2) {`,
				`|  (int) 1: (string) (len=1) "a"`,
				`|  (int) 2: (string) (len=1) "b"`,
			},
		},
		"struct": {
			a: []any{struct{ x int }{}},
			wants: []string{
				`| (struct { x int }) {`,
				`|  x: (int) 0`,
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			debugger.DumpAlwaysTo(&buf, "", tc.a...)
			result := buf.String()
			tester.AssertEqual(t, true, strings.Contains(result, "[DUMP]"))
			for c, w := range tc.wants {
				if !tester.AssertEqual(t, true, strings.Contains(result, w)) {
					t.Log(c, ": ", "`"+w+"`", "is not contained")
				}
			}
		})
	}
}
