package DB

import (
	"CMS/Config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//외부에서 사용할 수 있도록 패키지Func의 첫글자는 대문자
//패키지 관련 정보
//http://pyrasis.com/book/GoForTheReallyImpatient/Unit39

var (
	DBCon *sql.DB // Note the sql package provides the namespace
	Conf  Config.Configuration
)

func Init() {
	Conf = Config.GetConfiguration()
	var err error
	DBCon, err = sql.Open("mysql", Conf.DBInfo)
	if err != nil {
		log.Fatal(err)
	}

}

func CloseDB() {
	DBCon.Close()
}
func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
