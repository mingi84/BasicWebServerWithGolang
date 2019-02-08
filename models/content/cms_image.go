package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db Image Model
type CMS_Image struct {
	gorm.Model
	Width                                   int
	Height                                  int
	CMS_ContentID int
}
