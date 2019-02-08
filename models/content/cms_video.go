package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db Video model
type CMS_Video struct {
	gorm.Model
	//Map_Video_Extension Map_Video_Extension
	FileExtension string
	Width                                   int
	Height                                  int
	CMS_ContentID int
}

