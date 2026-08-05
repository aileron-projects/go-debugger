package debugger

import (
	"io"
	"os"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var (
	// muDumpErr protects dumpErrTo
	muDumpErr = sync.Mutex{}
	// dumpErrTo is the error dump output destination.
	// Use "DUMPERR_OUTPUT" environment variable.
	dumpErrTo io.Writer = os.Stdout
	// dumpErrPkg is the package names to dump.
	// If empty all dumps will be output.
	// Use "DUMPERR_PACKAGES" environment variable.
	dumpErrPkg []string = nil
	// DumpErrConfig is the dump format config.
	// See [spew.ConfigState].
	DumpErrConfig = spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}
)

func init() {
	dumpErrTo = getWriter(os.Getenv("DUMPERR_OUTPUT"), "go-dumperr-*.log")
	for s := range strings.SplitSeq(os.Getenv("DUMPERR_PACKAGES"), ",") {
		if dp := strings.TrimSpace(s); dp != "" {
			dumpPkg = append(dumpPkg, dp)
		}
	}
}

// DumpErr prints error dumps.
// DumpErr is effective when dump is enabled by build tag.
func DumpErr(msg string, errs ...error) {
	if dumpErrEnabled {
		muDumpErr.Lock()
		defer muDumpErr.Unlock()
		dumpErr(dumpErrTo, msg, errs...)
	}
}

// DumpErrTo prints error dumps to the given writer.
// DumpErrTo is effective when dump is enabled by build tag.
func DumpErrTo(w io.Writer, msg string, errs ...error) {
	if dumpErrEnabled {
		dumpErr(w, msg, errs...)
	}
}

// DumpErrAlways prints object dumps.
// Unlike [DumpErr] and [DumpErrTo], it does not requires build tags.
func DumpErrAlways(msg string, errs ...error) {
	muDumpErr.Lock()
	defer muDumpErr.Unlock()
	dumpErr(dumpErrTo, msg, errs...)
}

// DumpErrAlwaysTo prints error dumps to the given writer.
// Unlike [DumpErr] and [DumpErrTo], it does not requires build tags.
func DumpErrAlwaysTo(w io.Writer, msg string, errs ...error) {
	dumpErr(w, msg, errs...)
}

func dumpErr(w io.Writer, msg string, errs ...error) {
	if len(errs) == 0 {
		return
	}

	loc := FrameLocation(CallerFrame(2))
	if len(dumpErrPkg) > 0 && !slices.Contains(dumpErrPkg, loc.Pkg) {
		return
	}

	stack := make([]byte, 1<<13) // Read max 8kiB.
	n := runtime.Stack(stack, false)

	_, _ = w.Write([]byte(time.Now().Format(time.DateTime) + " [DEBUGGER][DUMPERR] " + msg + "\n"))
	_, _ = w.Write([]byte("  | Caller: " + loc.String() + "\n"))

	pw := &prefixWriter{
		w:       w,
		newline: true,
		prefix:  []byte("  | "),
	}
	dc := &DumpErrConfig
	for _, err := range errs {
		_, _ = pw.Write([]byte("┌── Error: " + err.Error() + "\n"))
		dc.Fdump(pw, err)
	}
	_, _ = pw.Write([]byte("┌── Stack Trace:\n"))
	_, _ = pw.Write(stack[:n])
}
