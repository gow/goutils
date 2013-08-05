package debug

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func LogMsg(v ...interface{}) {
	const colorGreen = "\x1b[34m"
	logHelper(colorGreen, 3, v...)
}
func LogWarn(v ...interface{}) {
	const colorYellow = "\x1b[33m"
	logHelper(colorYellow, 3, v...)
}
func LogError(v ...interface{}) {
	const colorRed = "\x1b[31m"
	logHelper(colorRed, 3, v...)
}

func logHelper(color string, callStackDepth int, v ...interface{}) {
	const colorEnd = "\x1b[0m"
	currentLogFlags := log.Flags()
	log.SetFlags(0)
	log.SetPrefix(color + getLogPrefix(callStackDepth) + colorEnd)
	var op string
	for _, i := range v {
		if str, ok := i.(string); ok {
			op += string(str)
		} else {
			op += PP(i)
		}
	}
	log.Print(op)
	log.SetPrefix("")
	log.SetFlags(currentLogFlags)
}

func getLogPrefix(callDepth int) string {
	var timePrefix string = getLogTimeString()
	_, path, line, ok := runtime.Caller(callDepth)
	fileName := filepath.Base(path)
	if !ok {
		return timePrefix + " ??? "
	}
	return timePrefix + " " + fileName + ":" + strconv.Itoa(line) + ": "
}

func getLogTimeString() string {
	t := time.Now()
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
}
