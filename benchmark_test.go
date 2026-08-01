package debugger_test

import (
	"testing"

	"github.com/aileron-projects/go-debugger"
)

func BenchmarkCallerFrame(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		debugger.CallerFrame(0)
	}
}

func BenchmarkCallerFrames(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		debugger.CallerFrames(0)
	}
}

func BenchmarkConvertFrame(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		debugger.FrameToLocation(debugger.CallerFrame(0))
	}
}

func BenchmarkConvertFrames(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		debugger.FramesToLocations(debugger.CallerFrames(0))
	}
}
