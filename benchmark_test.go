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

func BenchmarkFrameLocation(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		debugger.FrameLocation(debugger.CallerFrame(0))
	}
}

func BenchmarkFramesLocations(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		debugger.FramesLocations(debugger.CallerFrames(0))
	}
}
