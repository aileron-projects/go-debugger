package debugger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	// createTemp creates temporary file.
	// This variable is used for testing purpose.
	createTemp = os.CreateTemp
)

// initDebug initialize debug output.
// prefix specifies the file name prefix used for file output.
// The argument writeTo means
//   - "stdout": Standard output.
//   - "stderr": Standard error output.
//   - "discard": Discard all output.
//   - "file": File output. The second arg createFile must be specified.
//   - Others: Ignored.
func getWriter(writeTo string, prefix string) io.Writer {
	switch strings.ToLower(writeTo) {
	case "", "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	case "discard":
		return io.Discard
	case "file":
		f, err := createTemp("", prefix)
		if err != nil {
			err = fmt.Errorf("go-debugger/debugger: creating dump file failed [%w]", err)
			panic(err)
		}
		return f
	default:
		err := fmt.Errorf("go-debugger/debugger: unknown dump target `%s`", writeTo)
		panic(err)
	}
}

// prefixWriter writes lines with prefix.
type prefixWriter struct {
	w       io.Writer
	newline bool
	prefix  []byte
}

func (pw *prefixWriter) Write(p []byte) (n int, err error) {
	for len(p) > 0 {
		if i := bytes.IndexByte(p, '\n'); i >= 0 {
			if pw.newline {
				nn, _ := pw.w.Write(pw.prefix)
				n += nn
			}
			nn, _ := pw.w.Write(p[:i+1])
			n += nn
			p = p[i+1:]
			pw.newline = true
		} else {
			if pw.newline {
				nn, _ := pw.w.Write(pw.prefix)
				n += nn
			}
			nn, _ := pw.w.Write(p)
			n += nn
			p = nil
			pw.newline = false
		}
	}
	return n, nil
}
