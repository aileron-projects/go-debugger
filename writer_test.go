package debugger

import (
	"io"
	"os"
	"testing"

	"github.com/aileron-projects/go-tester"
)

func TestGetWriter(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		w := getWriter("", "")
		tester.AssertEqual(t, io.Writer(os.Stdout), w)
	})
	t.Run("stdout", func(t *testing.T) {
		w := getWriter("Stdout", "")
		tester.AssertEqual(t, io.Writer(os.Stdout), w)
	})
	t.Run("stderr", func(t *testing.T) {
		w := getWriter("Stderr", "")
		tester.AssertEqual(t, io.Writer(os.Stderr), w)
	})
	t.Run("discard", func(t *testing.T) {
		w := getWriter("Discard", "")
		tester.AssertEqual(t, io.Writer(io.Discard), w)
	})
	t.Run("discard", func(t *testing.T) {
		w := getWriter("Discard", "")
		tester.AssertEqual(t, io.Writer(io.Discard), w)
	})
	t.Run("file", func(t *testing.T) {
		var dir, pattern string
		f := &os.File{}
		createTemp = func(d, p string) (*os.File, error) {
			dir = d
			pattern = p
			return f, nil
		}
		w := getWriter("File", "test-*.log")
		tester.AssertEqual(t, io.Writer(f), w)
		tester.AssertEqual(t, "", dir)
		tester.AssertEqual(t, "test-*.log", pattern)
	})
	t.Run("file panic", func(t *testing.T) {
		defer func() {
			r := recover()
			tester.AssertEqual(t, true, r != nil)
		}()
		createTemp = func(_, _ string) (*os.File, error) {
			return nil, os.ErrPermission
		}
		w := getWriter("File", "test-*.log")
		tester.AssertEqual(t, nil, w) // This wont't be run.
	})
	t.Run("unknown panic", func(t *testing.T) {
		defer func() {
			r := recover()
			tester.AssertEqual(t, true, r != nil)
		}()
		w := getWriter("unknown", "")
		tester.AssertEqual(t, nil, w) // This wont't be run.
	})
}
