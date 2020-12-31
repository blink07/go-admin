package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize int
}

var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	LogDir string
	LogFile string
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DataBaseSettings = &Database{}

// 读取配置文件
func Setup() {
	Cfg,err := ini.Load("conf/config.ini")
	if err != nil {
		log.Fatalf("fail to parse 'conf/app.ini': :%v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil{
		log.Fatalf("Cfg.MapTo AppSetting err :%v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err :%v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DataBaseSettings)
	if err!= nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err:%v", err)
	}
	log.Println(DataBaseSettings)
}