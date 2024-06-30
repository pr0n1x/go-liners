package ctxlog

import (
	"fmt"
	"os"
)

type FileOrStdOutLogger struct {
	file *os.File
}

func NewFileLogger(path string) (logger *FileOrStdOutLogger, err error) {
	logger = &FileOrStdOutLogger{}
	if err = logger.Open(path); err != nil {
		logger = nil
	}
	return logger, err
}

func (l *FileOrStdOutLogger) Open(path string) error {
	if l.file != nil {
		return nil
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.file = file
	return nil
}

func (l *FileOrStdOutLogger) Close() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	if err == nil {
		l.file = nil
	}
	return err
}

func (l *FileOrStdOutLogger) Write(p []byte) (n int, err error) {
	if l.file == nil {
		return os.Stdout.Write(p)
	}
	return l.file.Write(p)
}

func (l *FileOrStdOutLogger) Print(a ...any) {
	if l.file == nil {
		fmt.Print(a...)
	}
	_, _ = fmt.Fprint(l.file, a...)
}

func (l *FileOrStdOutLogger) Println(a ...any) {
	if l.file == nil {
		fmt.Println(a...)
	}
	_, _ = fmt.Fprintln(l.file, a...)
}

func (l *FileOrStdOutLogger) Printf(format string, a ...any) {
	if l.file == nil {
		fmt.Printf(format, a...)
	}
	_, _ = fmt.Fprintf(l.file, format, a...)
}

func (l *FileOrStdOutLogger) Output(_ int, s string) error   { l.Print(s); return nil }
func (l *FileOrStdOutLogger) Fatal(a ...any)                 { l.Print(a...) }
func (l *FileOrStdOutLogger) Fatalln(a ...any)               { l.Println(a...) }
func (l *FileOrStdOutLogger) Fatalf(format string, a ...any) { l.Printf(format, a...) }
func (l *FileOrStdOutLogger) Panic(a ...any)                 { l.Print(a...) }
func (l *FileOrStdOutLogger) Panicln(a ...any)               { l.Println(a...) }
func (l *FileOrStdOutLogger) Panicf(format string, a ...any) { l.Printf(format, a...) }
