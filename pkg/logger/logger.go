package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"io"
)

func Info(message string) {
    logFile := getLogFile("/tmp/info.log")
    logger, err := getLogger("INFO", logFile)

    if err != nil {
        fmt.Println(err)
    }
    logger.Print(message)
}

func Error(message string) {
    logFile := getLogFile("/tmp/error.log")
    logger, err := getLogger("ERROR", logFile)

    if err != nil {
        fmt.Println(err)
    }
    logger.Print(message)
}

func Fatal(message string) {
    logFile := getLogFile("/tmp/fatal.log")
    logger, err := getLogger("FATAL", logFile)

    if err != nil {
        fmt.Println(err)
    }
    logger.Print(message)
}

func getLogger(level string, logFile io.Writer) (*log.Logger, error) {
    if err := isValid(level); err != nil{
        return nil, err
    }

    return log.New(logFile, fmt.Sprintf("%s\t", level), log.Ldate|log.Ltime|log.Lshortfile), nil
}

func isValid(loggerType string) error {
    switch loggerType {
    case "INFO", "ERROR", "FATAL":
        return nil
    }
    return errors.New(
    "Invalid logger type")
}

func getLogFile(fileName string) io.Writer {
    logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    return logFile
}
