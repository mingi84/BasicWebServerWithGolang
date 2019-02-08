package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db Content Model
type CMS_Content struct {
	gorm.Model
	ContentID          string
	ContentType        string
	FilePath    	   string
	FileName    	   string
	State              int
	CMS_Video CMS_Video
	CMS_Audio CMS_Audio
	CMS_Image CMS_Image
}
