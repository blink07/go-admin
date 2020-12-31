package log

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
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logClient := logrus.New()

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	println("LOG>>>>>>>>>>>>>>>>>>")
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

