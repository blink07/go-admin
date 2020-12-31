package log

//var logClient = logrus.New()

//func Info(msg ...interface{}) {
//	localLogger.Info(msg)
//}
//
//func Debug(msg ...interface{}) {
//	localLogger.Debug(msg)
//}
//
//func Error(err ...interface{}) {
//	localLogger.Error(err)
//}


func Info(msg ...interface{}) {
	logClient.Info(msg)
}

func Debug(msg ...interface{}) {
	logClient.Debug(msg)
}

func Error(err ...interface{}) {
	logClient.Error(err)
}