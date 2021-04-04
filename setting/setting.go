package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type Server struct {
	HttpPort     int
	DriverName  string
	Database   string
}


var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'app.ini': %v", err)
	}
	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", "SERVER", err)
	}
}