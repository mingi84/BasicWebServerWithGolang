package API

import (
	"fmt"
	"net/http"
	"CMS/models/content"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func ContentCreate
func ContentCreate(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	ContentID := fmt.Sprintf("%s", queryValues.Get("ContentID"))
	
	video := content.CMS_Video{FileExtension:"mp4", Width:1920, Height:1080}
	audio :=content.CMS_Audio{}
	image :=content.CMS_Image{}
	content := content.CMS_Content{ContentID:ContentID, ContentType:"Video", FilePath:"c:\\test\\test\\", FileName:"SEQ_MERGE.mp4", State:1, CMS_Video:video, CMS_Audio:audio, CMS_Image:image}

	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", "ID", "Password", "Ipaddress", "Port", "DBName"))
	defer db.Close()
	if err == nil {
		} else {
			panic(err)
		}
	
	w.Write([]byte(("Insert Complete")))	

	db.Create(&content)
}

//func ContentRead
func ContentRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Call ContentRead"))
}

//func ContentUpdate
func ContentUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Call ContentUpdate"))
}

//func ContentDelete
func ContentDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Call ContentDelete"))
}