package debugger

import (
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var (
	// muDump protects dumpTo and dumpPkg.
	muDump = sync.Mutex{}
	// dumpTo is the dump output destination.
	// Use "DUMP_OUTPUT" environment variable.
	dumpTo io.Writer = os.Stdout
	// dumpPkg is the package names to dump.
	// If empty all dumps will be output.
	// Use "DUMP_PACKAGES" environment variable.
	dumpPkg []string = nil
	// DumpConfig is the dump format config.
	// See [spew.ConfigState].
	DumpConfig = spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}
)

func init() {
	dumpTo = getWriter(os.Getenv("DUMP_OUTPUT"), "go-dump-*.log")
	for s := range strings.SplitSeq(os.Getenv("DUMP_PACKAGES"), ",") {
		if dp := strings.TrimSpace(s); dp != "" {
			dumpPkg = append(dumpPkg, dp)
		}
	}
}

// Dump prints object dumps.
// Dump is effective when dump is enabled by build tag.
func Dump(msg string, a ...any) {
	if dumpEnabled {
		muDump.Lock()
		defer muDump.Unlock()
		dump(dumpTo, msg, a...)
	}
}

// DumpTo prints object dumps to the given writer.
// DumpTo is effective when dump is enabled by build tag.
func DumpTo(w io.Writer, msg string, a ...any) {
	if dumpEnabled {
		dump(w, msg, a...)
	}
}

// DumpAlways prints object dumps.
// Unlike [Dump] and [DumpTo], it does not requires build tags.
func DumpAlways(msg string, a ...any) {
	muDump.Lock()
	defer muDump.Unlock()
	dump(dumpTo, msg, a...)
}

// DumpAlwaysTo prints object dumps to the given writer.
// Unlike [Dump] and [DumpTo], it does not requires build tags.
func DumpAlwaysTo(w io.Writer, msg string, a ...any) {
	dump(w, msg, a...)
}

func dump(w io.Writer, msg string, a ...any) {
	loc := FrameLocation(CallerFrame(2))
	if len(dumpPkg) > 0 && !slices.Contains(dumpPkg, loc.Pkg) {
		return
	}

	_, _ = w.Write([]byte(time.Now().Format(time.DateTime) + " [DUMP] " + msg + "\n"))
	_, _ = w.Write([]byte("  | Caller: " + loc.String() + "\n"))

	pw := &prefixWriter{
		w:       w,
		newline: true,
		prefix:  []byte("  | "),
	}
	dc := &DumpConfig
	for i := range a {
		_, _ = pw.Write([]byte("┌── args[" + strconv.Itoa(i) + "]\n"))
		dc.Fdump(pw, a[i])
	}
}
