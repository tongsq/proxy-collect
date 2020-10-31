package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var logger *log.Logger

var errorLogger *log.Logger

func init() {
	dir, _ := os.Getwd()
	errFile, err := os.OpenFile(dir+"/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open log file failed：", err)
	}
	log.SetFlags(log.Ltime | log.Lshortfile)
	//errorLogger = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errFile, "Error:", log.Ldate|log.Ltime)
	logger = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime)
}

func Info(v ...interface{}) {
	logger.Println(getLogContent("", v)...)
}

func Success(v ...interface{}) {
	logger.Println(getLogContent("32", v)...)
}

func Warning(v ...interface{}) {
	logger.Println(getLogContent("33", v)...)
}

func Error(v ...interface{}) {
	logger.Println(getLogContent("31", v)...)
	errorLogger.Println(getLogContent("", v)...)
}

type CronLogger struct {
}

func (l *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	v := []interface{}{"cron info", msg, keysAndValues}
	Info(v...)
}

func (l *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	v := []interface{}{"cron error", err, msg, keysAndValues}
	Error(v...)
}

func getLogContent(color string, v []interface{}) []interface{} {
	content := []interface{}{}
	if color != "" {
		c := "\x1b[0;" + color + "m"
		content = append(content, c)
	}
	content = append(content, getFileInfo())
	for _, value := range v {
		content = append(content, value)
	}
	if color != "" {
		content = append(content, "\x1b[0m")
	}
	return content
}

func getFileInfo() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return fmt.Sprintf("%s:%d:", file, line)
}
