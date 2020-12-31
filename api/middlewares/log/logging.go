package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"  // 日志记录
//	"github.com/lestrrat/go-file-rotatelogs"  // 日志切割
//	"github.com/rifflock/lfshook"
//	"os"
//	"time"
//)
//
//func Logger() (gin gin.HandlerFunc){
//	logClient := logrus.New()
//
//	// 禁止logrus 的输出
//	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//
//	if err!= nil {
//		fmt.Println("err", err)
//	}
//
//	logClient.Out = src
//	logClient.SetLevel(logrus.DebugLevel)
//
//	apiLogPath := "api.log"
//
//	logWriter, err := rotatelogs.New(
//		apiLogPath + ".%Y-%m-%d-%H-%M.log",   // 日志文件格式
//		rotatelogs.WithLinkName(apiLogPath),  // 生成软连接，指向最新日志文件
//		rotatelogs.WithMaxAge(7*24*time.Hour),   // 文件最大保存时间
//		rotatelogs.WithRotationTime(24*time.Hour),  // 日志切割时间间隔
//		)
//
//	writeMap := lfshook.WriterMap{
//		logrus.InfoLevel: logWriter,
//		logrus.FatalLevel: logWriter,
//	}
//
//	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
//	logClient.AddHook(lfHook)
//
//
//	return func(context *gin.Context) {
//		// 开始时间
//		start := time.Now()
//
//		context.Next()
//	}
//
//}

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var logClient=logrus.New()

func Logger() gin.HandlerFunc {
	//logClient := logrus.New()

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_SYNC, os.ModeAppend)
	if err!= nil{
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	apiLogPath := "logs/api.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {

	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)


	return func (c *gin.Context) {
		//开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

//win10不能生成软连接,使用无分割日志
//func WinLoggerHandler() gin.HandlerFunc {
//	// get log file
//	//logFilePath := config.ServerConfig.LogDir
//	//logFileName := config.ServerConfig.LogFile
//	logFilePath := settings.ServerSetting.LogDir
//	logFileName := settings.ServerSetting.LogFile
//
//	return initWinLogger(logFilePath, logFileName, localLogger)
//}
//
//func initWinLogger(logFilePath string, logFileName string, logger *logrus.Logger) gin.HandlerFunc {
//	//日志文件
//	fileName := path.Join(logFilePath, logFileName)
//
//
//	//写入文件
//	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//	if err != nil {
//		fmt.Println("err", err)
//	}
//
//	//设置输出
//	logger.Out = src
//
//	//设置日志级别
//	logger.SetLevel(logrus.DebugLevel)
//
//	//设置日志格式
//	logger.SetFormatter(&logrus.TextFormatter{})
//
//	return formatWinLog
//}
//
//func formatWinLog(c *gin.Context) {
//	// 开始时间
//	startTime := time.Now()
//
//	// 处理请求
//	c.Next()
//
//	// 结束时间
//	endTime := time.Now()
//
//	// 执行时间
//	latencyTime := endTime.Sub(startTime)
//
//	// 请求方式
//	reqMethod := c.Request.Method
//
//	// 请求路由
//	reqUri := c.Request.RequestURI
//
//	// 状态码
//	statusCode := c.Writer.Status()
//
//	// 请求IP
//	clientIP := c.ClientIP()
//
//	// 日志格式
//	localLogger.Infof("| %3d | %13v | %15s | %s | %s |",
//		statusCode,
//		latencyTime,
//		clientIP,
//		reqMethod,
//		reqUri,
//	)
//}

