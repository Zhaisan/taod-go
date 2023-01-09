package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)
//Чтобы логирование шло в два место, это файл и аутпут
//Чтобы было несколько уровней логирования
type writerHook struct {
	Writer []io.Writer
	LogLevels []logrus.Level
}

//будет вызываться fire каждый раз когда будем что то куда то писать

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}
// чтобы можно было изменять логи, чтобы не был синглтоном
var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger{
	return &Logger{e}
}
//

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}



func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard) //ничего никуда не пиши; чтобы логи никуда не уходили

	l.AddHook(&writerHook{
		Writer: []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel) //чтобы всё видели

	e = logrus.NewEntry(l)
}

//write:
//kafka -- info, debug
//file -- error, trace
//stdout -- warning, critical