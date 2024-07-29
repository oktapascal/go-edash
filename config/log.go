package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type LoggerFileHook struct {
	file      *os.File
	flag      int
	chmod     os.FileMode
	formatter *logrus.TextFormatter
}

// NewLoggerFileHook creates a new LoggerFileHook instance.
//
// file: The path to the log file.
// flag: The flag to open the log file with.
// chmod: The file mode to set for the log file.
//
// Returns a pointer to a LoggerFileHook instance and an error if any.
func NewLoggerFileHook(file string, flag int, chmod os.FileMode) (*LoggerFileHook, error) {

	// Create a new TextFormatter instance with custom settings
	textFormatter := &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	// Open the log file with the specified flag and chmod
	logFile, err := os.OpenFile(file, flag, chmod)
	if err != nil {
		// Print an error message and return nil and the error
		fmt.Printf("Failed to open log file: %s\n", err)
		return nil, err
	}

	// Return a new LoggerFileHook instance with the specified parameters
	return &LoggerFileHook{
		file:      logFile,
		flag:      flag,
		chmod:     chmod,
		formatter: textFormatter,
	}, nil
}

// Levels returns the levels of logrus.Level that this LoggerFileHook is configured to handle.
//
// It returns a slice of logrus.Level values representing the levels of log messages that this hook will handle.
// The levels are in the order of most severe to least severe: PanicLevel, FatalLevel, ErrorLevel, WarnLevel, InfoLevel, DebugLevel.
func (hook *LoggerFileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// Fire is a method of the LoggerFileHook struct that is called when a log entry needs to be written to the file.
//
// It takes a logrus.Entry as a parameter and returns an error.
// The function formats the log entry using the formatter specified in the LoggerFileHook struct,
// converts the formatted log entry to a string, and writes it to the file specified in the LoggerFileHook struct.
// If there is an error writing to the file, the function prints an error message and returns the error.
// Otherwise, it returns nil.
func (hook *LoggerFileHook) Fire(entry *logrus.Entry) error {
	var format, err = hook.formatter.Format(entry)
	var str = string(format)

	_, err = hook.file.WriteString(str)
	if err != nil {
		fmt.Printf("Failed to write log file: %s\n", err)
		return err
	}

	return nil
}

// createLogsDir creates the logs directory if it does not already exist.
// It returns an error if the directory creation fails.
func createLogsDir() error {
	// Define the path of the logs directory
	logsDir := "storage/logs"

	// Check if the logs directory exists
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		// If the logs directory does not exist, create it with the specified permissions (0755)
		return os.MkdirAll(logsDir, 0755)
	}

	// If the logs directory already exists, return nil (no error)
	return nil
}

// getLogFileName returns the name of the log file for the current date.
// The log file is stored in the "storage/logs" directory.
// The file name follows the format "YYYY-MM-DD.log".
func getLogFileName() string {
	// Generate the current date in the format "YYYY-MM-DD"
	currentDate := time.Now().Format("2006-01-02")

	// Construct the full path of the log file
	logFilePath := filepath.Join("storage/logs", currentDate+".log")

	return logFilePath
}

// CreateLoggers creates a logger with a custom output, formatter, and level.
// It also adds a file hook to log to a file.
// If a request is provided, additional fields are added to the logger.
// Returns the logger entry or nil if there was an error creating the logs directory.
func CreateLoggers(request *http.Request) *logrus.Entry {
	// Create a new logger
	logger := logrus.New()

	// Set the output to stdout and configure the formatter
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(logrus.DebugLevel)

	// Create the logs directory if it doesn't exist
	var errDirectory = createLogsDir()
	if errDirectory != nil {
		fmt.Println(errDirectory)
		return nil
	}

	// Get the name of the log file for the current date
	logFileName := getLogFileName()

	// Create a file hook to log to a file
	var loggerFileHook, err = NewLoggerFileHook(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Hooks.Add(loggerFileHook)
	}

	// If no request is provided, return the logger with empty fields
	if request == nil {
		return logger.WithFields(logrus.Fields{})
	}

	// If a request is provided, add additional fields to the logger
	return logger.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": request.Method,
		"uri":    request.RequestURI,
		"ip":     request.RemoteAddr,
	})
}
