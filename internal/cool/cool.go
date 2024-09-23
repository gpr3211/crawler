package cool

import (
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"path"
	"sync"
)

var Red = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("9"))

var Blue = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("86"))

var (
	FAIL = Red.Render("WARNING")
	PASS = Blue.Render("INFO")
)

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	Lmsgprefix
)

type FlagList []int

var (
	mu     sync.Mutex
	logger *log.Logger
	cfg    *Coolfig
	info   *log.Logger
	warn   *log.Logger
	fatal  *log.Logger
)

type Coolfig struct {
	LogFileName string
	LogDir      string
	Flags       int
	FlagList    FlagList
}

func Initialize(c *Coolfig) error {
	mu.Lock()
	defer mu.Unlock()

	defaultConfig := Coolfig{
		LogFileName: "cool.log",
		LogDir:      os.TempDir(),
		Flags:       log.Ldate | log.Ltime | log.Lshortfile,
	}

	if c.Flags == 0 {
		c = &defaultConfig
	}

	if c.LogFileName == "" {
		c.LogFileName = defaultConfig.LogFileName
	}
	if c.LogDir == "" {
		c.LogDir = defaultConfig.LogDir
	}
	if c.Flags == 0 {
		c.Flags = defaultConfig.Flags
	}

	logFilePath := path.Join(c.LogDir, c.LogFileName)

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logger = log.New(f, "server : ", c.Flags)
	cfg = c
	info = log.New(os.Stdout, "", c.Flags)
	warn = log.New(os.Stdout, "", c.Flags)

	fatal = log.New(os.Stdout, "", c.Flags)

	return nil
}

func AddFlag(flag int) {
	mu.Lock()
	defer mu.Unlock()
	if logger != nil {
		cfg.Flags |= flag
		logger.SetFlags(cfg.Flags)
	}
}

func RemoveFlag(flag int) {
	mu.Lock()
	defer mu.Unlock()
	if logger != nil {
		cfg.Flags &^= flag
		logger.SetFlags(cfg.Flags)
	}
}

// Global wrapper functions
func Print(v ...interface{}) {
	if logger != nil {
		logger.Print(v...)
	}
}

func Println(v ...interface{}) {
	if logger != nil {
		logger.Println(v...)
	}
}

func Printf(format string, v ...interface{}) {
	if logger != nil {
		logger.Printf(format, v...)
	}
}
func Warn(format string, v ...interface{}) {
	yellow := "\033[33m"
	reset := "\033[0m" // ANSI escape code to reset color
	full := " " + yellow + "WARNIN :" + reset + " " + format
	if warn != nil {
		warn.Println(full)
	}
	// Printing the message in blue
	if logger != nil {
		logger.Println("WARNING: " + format)
	}

}

func Info(format string, v ...interface{}) {
	blue := "\033[34m"
	reset := "\033[0m" // ANSI escape code to reset color
	full := " " + blue + "INFO:" + reset + " " + format

	if info != nil {
		info.Println(full)
	}

	// Printing the message in blue
	if logger != nil {
		logger.Println("INFO: " + format)
	}

}

func Fatal(format string, v ...interface{}) {
	red := "\033[31m"
	reset := "\033[0m" // ANSI escape code to reset color
	full := " " + red + "FATAL:" + reset + " " + format
	if fatal != nil {
		fatal.Println(full)
	}

	// Printing the message in blue
	if logger != nil {
		logger.Fatal(format)
	}

}

func Fatalln(v ...interface{}) {
	if logger != nil {
		logger.Fatalln(v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	if logger != nil {
		logger.Fatalf(format, v...)
	}
}

func Panic(v ...interface{}) {
	if logger != nil {
		logger.Panic(v...)
	}
}

func Panicln(v ...interface{}) {
	if logger != nil {
		logger.Panicln(v...)
	}
}

func Panicf(format string, v ...interface{}) {
	if logger != nil {
		logger.Panicf(format, v...)
	}
}
