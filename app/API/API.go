package API

import (
	"CMS/Config"
	_ "CMS/app/DB"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conf Config.Configuration
)

//검사서버 등록
func Init() {
	Conf = Config.GetConfiguration()
}
