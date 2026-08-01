package debugger_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/aileron-projects/go-debugger"
	"github.com/aileron-projects/go-tester"
)

func TestFrameToLocation(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		frame runtime.Frame
		want  debugger.Location
		str   string
	}{
		"empty frame": {
			frame: runtime.Frame{},
			want:  debugger.Location{},
			str:   "Pkg: File: Func: Line:0",
		},
		"non empty frame": {
			frame: runtime.Frame{PC: 123, Function: "foo/bar.testFunc", File: "test.go", Line: 100},
			want:  debugger.Location{Pkg: "foo/bar", File: "test.go", Func: "testFunc", Line: 100},
			str:   "Pkg:foo/bar File:test.go Func:testFunc Line:100",
		},
		"short func name": {
			frame: runtime.Frame{PC: 123, Function: "bar.testFunc", File: "test.go", Line: 100},
			want:  debugger.Location{Pkg: "bar", File: "test.go", Func: "testFunc", Line: 100},
			str:   "Pkg:bar File:test.go Func:testFunc Line:100",
		},
		"short file name": {
			frame: runtime.Frame{PC: 123, Function: "foo/bar.testFunc", File: "baz/test.go", Line: 100},
			want:  debugger.Location{Pkg: "foo/bar", File: "test.go", Func: "testFunc", Line: 100},
			str:   "Pkg:foo/bar File:test.go Func:testFunc Line:100",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			f := debugger.FrameToLocation(tc.frame)
			tester.AssertEqual(t, tc.want, f)
			tester.AssertEqual(t, tc.str, f.String())
		})
	}
}

func TestFramesToLocations(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		frames []runtime.Frame
		want   []debugger.Location
	}{
		"empty": {
			frames: nil,
			want:   nil,
		},
		"1 frame": {
			frames: []runtime.Frame{
				{PC: 1, Function: "foo/bar.testFunc1", File: "test1.go", Line: 101},
			},
			want: []debugger.Location{
				{Pkg: "foo/bar", File: "test1.go", Func: "testFunc1", Line: 101},
			},
		},
		"2 frames": {
			frames: []runtime.Frame{
				{PC: 1, Function: "foo/bar.testFunc1", File: "test1.go", Line: 101},
				{PC: 1, Function: "bar/foo.testFunc2", File: "test2.go", Line: 102},
			},
			want: []debugger.Location{
				{Pkg: "foo/bar", File: "test1.go", Func: "testFunc1", Line: 101},
				{Pkg: "bar/foo", File: "test2.go", Func: "testFunc2", Line: 102},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			fs := debugger.FramesToLocations(tc.frames)
			if !(len(fs) == len(tc.want)) {
				t.Errorf("length not match. want:%d got:%d", len(tc.want), len(fs))
			}
			for i, f := range fs {
				w := tc.want[i]
				tester.AssertEqual(t, w, f)
			}
		})
	}
}

func TestCallerFrame(t *testing.T) {
	t.Parallel()
	t.Run("skip=0", func(t *testing.T) {
		f := debugger.CallerFrame(0)
		tester.AssertEqual(t, true, strings.HasSuffix(f.File, "frame_test.go"))
		tester.AssertEqual(t, true, strings.HasSuffix(f.Function, "TestCallerFrame.func1"))
		tester.AssertEqual(t, true, f.Line > 0)
	})
	t.Run("skip=9999", func(t *testing.T) {
		f := debugger.CallerFrame(9999)
		tester.AssertEqual(t, "", f.File)
		tester.AssertEqual(t, "", f.Function)
		tester.AssertEqual(t, 0, f.Line)
	})

	t.Run("skip=-9999", func(t *testing.T) {
		f := debugger.CallerFrame(-9999)
		tester.AssertEqual(t, true, strings.HasSuffix(f.File, "runtime/extern.go"))
		tester.AssertEqual(t, true, strings.HasSuffix(f.Function, "runtime.Callers"))
		tester.AssertEqual(t, true, f.Line > 0)
	})
}

func TestCallerFrames(t *testing.T) {
	t.Parallel()

	t.Run("skip=0", func(t *testing.T) {
		fs := debugger.CallerFrames(0)
		tester.AssertEqual(t, true, len(fs) > 1)
		f := fs[0]
		tester.AssertEqual(t, true, strings.HasSuffix(f.File, "frame_test.go"))
		tester.AssertEqual(t, true, strings.HasSuffix(f.Function, "TestCallerFrames.func1"))
		tester.AssertEqual(t, true, f.Line > 0)
	})

	t.Run("skip=9999", func(t *testing.T) {
		fs := debugger.CallerFrames(9999)
		tester.AssertEqual(t, 0, len(fs))
	})

	t.Run("skip=-9999", func(t *testing.T) {
		fs := debugger.CallerFrames(-9999)
		tester.AssertEqual(t, true, len(fs) > 1)
		f := fs[0]
		tester.AssertEqual(t, true, strings.HasSuffix(f.File, "runtime/extern.go"))
		tester.AssertEqual(t, true, strings.HasSuffix(f.Function, "runtime.Callers"))
		tester.AssertEqual(t, true, f.Line > 0)
	})
}
