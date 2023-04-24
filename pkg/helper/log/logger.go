package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()

	path := "log/"
	logFileName := "app.log"

	file, err := os.OpenFile(path+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log.Out = file

	fileInfo, err := os.Stat(logFileName)
	if err != nil {
		panic(err)
	}
	fileModTime := fileInfo.ModTime()
	now := time.Now()

	if fileModTime.Day() != now.Day() {
		logFileName = "app_" + now.Format("2006-01-02") + ".log"
		file.Close()

		file, err = os.OpenFile(path+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log.Out = file
	}
}

func Info(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)
}

func Debug(message string, fields map[string]interface{}) {
	log.WithFields(fields).Debug(message)
}

func Warn(message string, fields map[string]interface{}) {
	log.WithFields(fields).Warn(message)
}

func Error(message string, fields map[string]interface{}) {
	log.WithFields(fields).Error(message)
}

func Close() {

}
