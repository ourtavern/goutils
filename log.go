package goutils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	InfoColor  = "\033[1;32m"
	ErrorColor = "\033[1;31m"
	WarnColor  = "\033[1;33m"
	FatalColor = "\033[1;35m"
	DebugColor = "\033[0;36m"
	TraceColor = "\033[0;34m"
	ResetColor = "\033[0m"
)

func colorize(color, message string) string {
	return color + message + ResetColor
}

func LogInfo(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(colorize(InfoColor, fmt.Sprintf("[INFO] %s", message)))
}

func LogError(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(colorize(ErrorColor, fmt.Sprintf("[ERROR] %s", message)))
}

func LogWarn(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(colorize(WarnColor, fmt.Sprintf("[WARN] %s", message)))
}

func LogFatal(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(colorize(FatalColor, fmt.Sprintf("[FATAL] %s", message)))
}

func LogDebug(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(colorize(DebugColor, fmt.Sprintf("[DEBUG] %s", message)))
}

func LogTrace(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(colorize(TraceColor, fmt.Sprintf("[TRACE] %s", message)))
}

func LogNone(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(message)
}

func LogClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
