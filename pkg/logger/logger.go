package logger

import (
	"log"
	"path/filepath"
	"time"

	"libra/pkg/util"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Setup(path string) {
	util.CreateDir(path)
	writer, err := rotatelogs.New(
		filepath.Join(path, "%Y%m%d.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Fatalf("config logger error. %+v", err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &Formatter{})

	logrus.SetReportCaller(true)
	logrus.AddHook(lfHook)
}
