package debugger

import (
	"runtime"
	"strconv"
	"strings"
)

// Location is the frame location.
// See also [runtime.Frame].
type Location struct {
	// Pkg is package name.
	Pkg string `json:"pkg" msgpack:"pkg" xml:"pkg" yaml:"pkg"`
	// File is the file name.
	File string `json:"file" msgpack:"file" xml:"file" yaml:"file"`
	// Func is the function name.
	Func string `json:"func" msgpack:"func" xml:"func" yaml:"func"`
	// Line is the line number.
	Line int `json:"line" msgpack:"line" xml:"line" yaml:"line"`
}

func (l *Location) String() string {
	var builder strings.Builder
	builder.Grow(80)
	builder.WriteString("Pkg:")
	builder.WriteString(l.Pkg)
	builder.WriteString(" File:")
	builder.WriteString(l.File)
	builder.WriteString(" Func:")
	builder.WriteString(l.Func)
	builder.WriteString(" Line:")
	builder.WriteString(strconv.Itoa(l.Line))
	return builder.String()
}

// FrameLocation converts [runtime.Frame] to [Location].
// It returns zero-value Location when f is empty.
func FrameLocation(f runtime.Frame) Location {
	if f.PC == 0 {
		return Location{}
	}
	file := f.File
	pkg, fn, pkgfn := "", f.Function, f.Function // pkgfn is "<Pkg>.<Func>"
	j := max(0, strings.LastIndexByte(pkgfn, '/'))
	if i := strings.IndexByte(pkgfn[j:], '.'); i > 0 {
		pkg, fn = pkgfn[:j+i], pkgfn[j+i+1:]
		file = strings.TrimPrefix(strings.TrimPrefix(file, pkg), "/")
	}
	return Location{
		Pkg:  pkg,
		File: file,
		Line: f.Line,
		Func: fn,
	}
}

// FramesLocations returns a slice of [Location] from frames.
func FramesLocations(fs []runtime.Frame) []Location {
	frames := make([]Location, len(fs))
	for i := range fs {
		frames[i] = FrameLocation(fs[i])
	}
	return frames
}

// CallerFrame returns single caller frame.
// skip is the number of stack frames to skip before recording frames.
// skip=0 means the caller frame and skip=1 means the caller of the caller.
// CallerFrame returns zero-value of [runtime.Frame]
// when there is no frames to report.
// See also [runtime.Caller].
func CallerFrame(skip int) runtime.Frame {
	pc := make([]uintptr, 1)
	n := runtime.Callers(skip+2, pc)
	if n < 1 {
		return runtime.Frame{} // No frame to report.
	}
	frame, _ := runtime.CallersFrames(pc).Next()
	return frame
}

// CallerFrames returns a slice of caller frames.
// skip is the number of stack frames to skip before recording frames.
// skip=0 means the caller frame and skip=1 means the caller of the caller.
// CallerFrames returns nil slice of [runtime.Frame]
// when there is no frames to report.
// CallerFrames returns 64 frames at maximum.
// See also [runtime.CallersFrames].
func CallerFrames(skip int) []runtime.Frame {
	pcs := make([]uintptr, 64) // Max 64 frames.
	n := runtime.Callers(skip+2, pcs)
	if n < 1 {
		return nil // No frames to report.
	}
	frames := runtime.CallersFrames(pcs[:n])
	fs := make([]runtime.Frame, n)
	for i := range n {
		fs[i], _ = frames.Next()
	}
	return fs
}
