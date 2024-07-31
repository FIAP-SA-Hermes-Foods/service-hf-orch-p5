package logger

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	logLevels = iota * 2
	infoLevel
	debugLevel
	diebugLevel
	warningLevel
	errorLevel
)

var logLevelsMap = map[int]string{
	infoLevel:    "INFO",
	debugLevel:   "DEBUG",
	diebugLevel:  "DIEBUG",
	warningLevel: "WARNING",
	errorLevel:   "ERROR",
}

func Info(message string) {
	generateLog(infoLevel, message, "", nil)
}

func Infof(message string, separator string, data ...interface{}) {
	generateLog(infoLevel, message, separator, data...)
}

func Debug(message string) {
	generateLog(debugLevel, message, "", nil)
}

func Debugf(message string, separator string, data ...interface{}) {
	generateLog(debugLevel, message, separator, data...)
}

func Diebug(message string) {
	generateLog(diebugLevel, message, "", nil)
}

func Diebugf(message string, separator string, data ...interface{}) {
	generateLog(diebugLevel, message, separator, data...)
}

func Warning(message string) {
	generateLog(warningLevel, message, "", nil)
}

func Warningf(message string, separator string, data ...interface{}) {
	generateLog(warningLevel, message, separator, data...)
}

func Error(message string) {
	generateLog(errorLevel, message, "", nil)
}

func Errorf(message string, separator string, data ...interface{}) {
	generateLog(errorLevel, message, separator, data...)
}

func generateLog(level int, message string, separator string, data ...interface{}) {
	var (
		msgStrWithData string
		logPrint       = fmt.Sprintf(`\e%s[%s]\e[0m %s`, colorByLabel(level), logLevelsMap[level], message)
	)

	for index, v := range data {
		if v != nil {

			if index == 0 {
				msgStrWithData = fmt.Sprintf("%s%v", msgStrWithData, v)
			} else {
				msgStrWithData = fmt.Sprintf("%s%s%v", msgStrWithData, separator, v)
			}
			if index == len(data)-1 {
				msgStrWithData = msgStrWithData + separator
			}
		}
	}

	logPrint = logPrint + msgStrWithData

	out, err := exec.Command("echo", "-e", logPrint).Output()

	if err != nil {
		Error(err.Error())
	}

	if level == diebugLevel {
		log.Fatal(string(out))
	}

	log.Print(string(out))
}

func colorByLabel(level int) string {
	switch level {
	case infoLevel:
		return "[32m"
	case debugLevel:
		return "[36m"
	case diebugLevel:
		return "[36m"
	case warningLevel:
		return "[33m"
	case errorLevel:
		return "[31m"
	default:
		return "[37m"
	}
}
