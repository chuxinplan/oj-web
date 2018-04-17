package g

import (
	"fmt"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

func InitLog() {
	conf := Conf()
	if !conf.Log.Enable {
		fmt.Println("log to std err")
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
		return
	}

	err := os.MkdirAll(conf.Log.Path, 0777)
	if err != nil {
		log.Fatalf("create directory %s failure\n", conf.Log.Path)
	}

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotationTime := time.Duration(conf.Log.RotatTime)
	lfhook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: getWriter(logPath, "debug", maxAge, rotationTime),
		log.InfoLevel:  getWriter(logPath, "info", maxAge, rotationTime),
		log.WarnLevel:  getWriter(logPath, "warn", maxAge, rotationTime),
		log.ErrorLevel: getWriter(logPath, "error", maxAge, rotationTime),
		log.FatalLevel: getWriter(logPath, "fatal", maxAge, rotationTime),
		log.PanicLevel: getWriter(logPath, "panic", maxAge, rotationTime),
	}, &log.TextFormatter{ForceColors: true})
	log.AddHook(lfhook)
}

// TODO 更改日志切割，为每月分割，并以不同级别设置不同输出目的
func getLogRotatFileHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) log.Hook {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d",
		rotatelogs.WithLinkName(logFileName),                // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge*time.Hour),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. ", err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: getWriter(), // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{ForceColors: true})
	return lfHook
}

func getWriter(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(logFileName),                // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge*time.Hour),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime*time.Hour), // 日志切割时间间隔
	)
	return writer
}
