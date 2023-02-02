package Config

import (
	"fmt"

	"github.com/caarlos0/env"
	"gopkg.in/ini.v1"
)

type Configuration struct {

	//검사<->관리서버 이미지 전달 FTP 설정
	IsConnectProjectManager string `env:"bIsConnectProjectManager" envDefault:"false"`
	PMIPAddress             string `env:"PMIPAddress" envDefault:""`
	WatchCallbackURL        string `env:"WatchCallbackURL" envDefault:""`
	WatcherURL              string `env:"WatcherURL" envDefault:""`
	PMPort                  string `env:"PMPort" envDefault:""`
	PMURL                   string `env:"PMURL" envDefault:""`
	PushType                string `env:"PushType" envDefault:"None"`
	AnalyzerURL             string `env:"AnalyzerURL" envDefault:""`
	DBInfo                  string `env:"DBInfo" envDefault:""`
}

//GetConfiguration
func GetConfiguration() Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg, err := ini.Load("setting.ini")
	if err != nil {
		fmt.Printf("%+v\n", err)
		fmt.Printf("setting.ini is not found\n")
	}
	configuration.IsConnectProjectManager = cfg.Section("CommunitySetting").Key("bIsConnectProjectManager").String()

	configuration.PushType = cfg.Section("CommunitySetting").Key("PushType").String()

	if configuration.IsConnectProjectManager == "true" {
		fmt.Printf("Config bIsConnect is true\n")

		configuration.PMIPAddress = cfg.Section("CommunitySetting").Key("PMIPAddress").String()
		configuration.PMPort = cfg.Section("CommunitySetting").Key("PMPort").String()
		configuration.PMURL = cfg.Section("CommunitySetting").Key("PMURL").String()
		fmt.Printf("Config PMIPAddress : " + configuration.PMIPAddress + "\n")
		fmt.Printf("Config PMPort : " + configuration.PMPort + "\n")
		fmt.Printf("Config PMURL : " + configuration.PMURL + "\n")

	} else {
		fmt.Printf("Config bIsConnect is false\n")
		configuration.PMIPAddress = ""
		configuration.PMPort = ""
		configuration.PMURL = ""
	}
	configuration.WatchCallbackURL = cfg.Section("CommunitySetting").Key("WatchCallbackURL").String()
	configuration.WatcherURL = cfg.Section("CommunitySetting").Key("WatcherURL").String()
	configuration.AnalyzerURL = cfg.Section("CommunitySetting").Key("AnalyzerURL").String()
	configuration.DBInfo = cfg.Section("CommunitySetting").Key("DBInfo").String()
	return configuration
}
