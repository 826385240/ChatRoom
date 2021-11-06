package logger

import (
	"fmt"
	"ChatRoom/src/lib/util"
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	filename string
	*log.Logger
}

func NewLogger(filename string, flag int, prefix string) *Logger {
	//filename不能用~表示用户目录
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("创建日志文件失败!", err.Error())
	}
	newLogger := &Logger{filename, log.New(file, prefix, flag)}
	return newLogger
}

func (this *Logger) SetFlags(flag int) {
	this.Logger.SetFlags(flag)
}

func (this *Logger) SetPrefix(prefix string) {
	this.Logger.SetPrefix(prefix)
}

func (this *Logger) Printf(format string, v ...interface{}) {
	this.Logger.Printf(format, v)
}

func (this *Logger) Print(v ...interface{}) {
	this.Logger.Print(v)
}

func (this *Logger) Println(v ...interface{}) {
	this.Logger.Println(v)
}

func (this *Logger) Fatalf(format string, v ...interface{}) {
	this.Logger.Fatalf(format, v)
}

func (this *Logger) Fatal(v ...interface{}) {
	this.Logger.Fatal(v)
}

func (this *Logger) Fatalln(v ...interface{}) {
	this.Logger.Fatalln(v)
}

func (this *Logger) Panic(v ...interface{}) {
	this.Logger.Panic(v)
}

func (this *Logger) Panicf(format string, v ...interface{}) {
	this.Logger.Panicf(format, v)
}

func (this *Logger) Panicln(v ...interface{}) {
	this.Logger.Panicln(v)
}

var g_logger *Logger

func InitLog(filename string) *Logger {
	if e, _ := util.PathExists(filename); !e {
		if e, _ := util.PathExists(filepath.Dir(filename)); !e {
			util.CreateDirByPath(filename)
		}
	}

	g_logger = NewLogger(filename, log.Ldate|log.Ltime|log.Llongfile|log.Lshortfile|log.LstdFlags, "")
	return g_logger
}

func GetGlobalLogger() *Logger {
	return g_logger
}

func DEBUG(format string, v ...interface{}) {
	g_logger.Printf("DEBUG %s", fmt.Sprintf(format, v...))
}

func DEBUGLN(v ...interface{}) {
	g_logger.Printf("DEBUG %s", fmt.Sprintln(v...))
}

func ERROR(format string, v ...interface{}) {
	g_logger.Printf("ERROR %s", fmt.Sprintf(format, v...))
}

func ERRORLN(v ...interface{}) {
	g_logger.Printf("ERROR %s", fmt.Sprintln(v...))
}

func INFO(format string, v ...interface{}) {
	g_logger.Printf("INFO %s", fmt.Sprintf(format, v...))
}

func INFOLN(v ...interface{}) {
	g_logger.Printf("INFO %s", fmt.Sprintln(v...))
}

func WARN(format string, v ...interface{}) {
	g_logger.Printf("WARN %s", fmt.Sprintf(format, v...))
}

func WARNLN(v ...interface{}) {
	g_logger.Printf("WARN %s", fmt.Sprintln(v...))
}

func FATAL(format string, v ...interface{}) {
	g_logger.Printf("FATAL %s", fmt.Sprintf(format, v...))
	os.Exit(1)
}

func FATALLN(v ...interface{}) {
	g_logger.Printf("FATAL %s", fmt.Sprintln(v...))
	os.Exit(1)
}

func PANIC(format string, v ...interface{}) {
	str := fmt.Sprintf("%s", fmt.Sprintf(format, v...))
	g_logger.Printf("PANIC %s", str)
	panic(str)
}

func PANICLN(v ...interface{}) {
	str := fmt.Sprintln(v...)
	g_logger.Printf("PANIC %s", str)
	panic(str)
}
