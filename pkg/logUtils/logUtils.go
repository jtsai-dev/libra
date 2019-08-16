/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:22:56
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-15 10:41:06
 */
package logUtils

import (
	"time"

	"libra/pkg"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// config logrus log to local filesystem, with file rotation
func Setup() {
	writer, err := rotatelogs.New(
		pkg.Configs.Log.Path+"%Y%m%d.log",
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{
		DisableColors:   true,
		TimestampFormat: pkg.Configs.Log.TimestampFormat,
		FullTimestamp:   true,
	})
	// &log.JSONFormatter{
	// 	PrettyPrint:     true,
	// 	TimestampFormat: pkg.Configs.Log.TimestampFormat,
	// })

	log.SetReportCaller(true)
	log.AddHook(lfHook)
}
