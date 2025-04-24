package trace

import (
	"runtime"
	"strings"
)

func CollectStack(pid int) string {
	// Placeholder: Just return current goroutine stack (not the real pid one)
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, true)
	return strings.TrimSpace(string(buf[:n]))
}
