package log

import (
  "bytes"
  "fmt"
  "log"
  "os"
  "strings"
)

// Level to integer map
var levelMap = map[string]int{
  "FATAL": 4,
  "ERROR": 3,
  "WARN":  2,
  "INFO":  1,
  "DEBUG": 0,
}

type Logger struct {
  name  string
  log   *log.Logger
  level int
}

func New(name string, filename string, level string) *Logger {
  return &Logger{
    level: levelMap[level],
    log:   log.New(os.Stderr, name, log.Ldate|log.Ltime),
  }
}

func (log *Logger) Debug(format string, args ...interface{}) {
  log.Write("DEBUG", format, args...)
}

func (log *Logger) Info(format string, args ...interface{}) {
  log.Write("INFO", format, args...)
}

func (log *Logger) Warn(format string, args ...interface{}) {
  log.Write("WARN", format, args...)
}

func (log *Logger) Error(format string, args ...interface{}) {
  log.Write("ERROR", format, args...)
}

func (log *Logger) Fatal(format string, args ...interface{}) {
  log.Write("FATAL", format, args...)
}

func (l *Logger) Write(level string, format string, args ...interface{}) {
  if levelMap[level] < l.level {
    return
  }

  var buffer bytes.Buffer

  buffer.WriteString(level)
  buffer.WriteString(": ")
  buffer.WriteString(fmt.Sprintf(format, args...))
  l.log.Println(strings.TrimSpace(buffer.String()))
}