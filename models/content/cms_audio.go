package content

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*

	AudioSamplingrate := fmt.Sprintf("%s", queryValues.Get("AudioSamplingrate"))
	AudioChannel := fmt.Sprintf("%s", queryValues.Get("AudioChannel"))
	AudioBitrate := fmt.Sprintf("%s", queryValues.Get("AudioBitrate"))
	AudioBitdepth := fmt.Sprintf("%s", queryValues.Get("AudioBitdepth"))
	AudioDuration := fmt.Sprintf("%s", queryValues.Get("AudioDuration"))
	AudioFormat := fmt.Sprintf("%s", queryValues.Get("AudioFormat"))

*/
// db Audio Model
type CMS_Audio struct {
	gorm.Model
	FileExtension string
	Channels      string
	Samplingrate  string
	Bitrate       string
	Bitdepth      string
	Duration      string
	Format        string
	Codec         string
	CMS_ContentID uint
}
