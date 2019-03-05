package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*

VideoFramerate := fmt.Sprintf("%s", queryValues.Get("VideoFramerate"))
	VideoWidth := fmt.Sprintf("%s", queryValues.Get("VideoWidth"))
	VideoHeight := fmt.Sprintf("%s", queryValues.Get("VideoHeight"))
	VideoCodec := fmt.Sprintf("%s", queryValues.Get("VideoCodec"))
	VideoFormat := fmt.Sprintf("%s", queryValues.Get("VideoFormat"))
	VideoScantype := fmt.Sprintf("%s", queryValues.Get("VideoScantype"))
	VideoDuration := fmt.Sprintf("%s", queryValues.Get("VideoDuration"))
	VideoDisplayaspectratio := fmt.Sprintf("%s", queryValues.Get("VideoDisplayaspectratio"))
	VideoColorspace := fmt.Sprintf("%s", queryValues.Get("VideoColorspace"))
	VideoBitdepth := fmt.Sprintf("%s", queryValues.Get("VideoBitdepth"))


*/

// db Video model
type CMS_Video struct {
	gorm.Model
	//Map_Video_Extension Map_Video_Extension
	FileExtension      string
	Width              int
	Height             int
	Codec              string
	Format             string
	Scantype           string
	Duration           string
	Displayaspectratio string
	Colorspace         string
	Bitdepth           string
	Framerate          string
	CMS_ContentID      uint
}
