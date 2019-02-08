package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db Audio Model
type CMS_Audio struct {
	gorm.Model
	Channels                                int
	SamplingRate                            string
	CMS_ContentID int
}
