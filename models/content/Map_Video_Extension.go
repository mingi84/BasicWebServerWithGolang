package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db Audio Model
type Map_Video_Extension struct {
	gorm.Model
	cms_videoID int
	FileExtension string
}
