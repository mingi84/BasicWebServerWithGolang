package API

import (
	"CMS/config/db"
	"CMS/models/content"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//http://doc.gorm.io/crud.html#query
// gorm 쿼리 참조

//func ContentCreate
func AddContent(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	ContentID := fmt.Sprintf("%s", queryValues.Get("ContentID"))
	ContentType := fmt.Sprintf("%s", queryValues.Get("ContentType"))
	FilePath := fmt.Sprintf("%s", queryValues.Get("FilePath"))

	content := content.CMS_Content{ContentID: ContentID, ContentType: ContentType, FilePath: FilePath, State: 1}
	db.DB.Create(&content)

	w.Write([]byte(("AddContent Complete")))
}

//func ContentCreate
func CompleteTranscoding(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	ContentID := fmt.Sprintf("%s", queryValues.Get("ContentID"))
	ContentType := fmt.Sprintf("%s", queryValues.Get("ContentType"))
	FileName := fmt.Sprintf("%s", queryValues.Get("FileName"))

	var thisvideo content.CMS_Video
	var thisaudio content.CMS_Audio
	var thisimage content.CMS_Image

	var thiscontent content.CMS_Content

	db.DB.Where("content_id=?", ContentID).First(&thiscontent)

	switch ContentType {
	case "video":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisvideo).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisvideo.FileExtension = filepath.Ext(FileName)
				thisvideo.CMS_ContentID = thiscontent.ID
				db.DB.Create(&thisvideo)
			}
		} else {
			db.DB.Model(thisvideo).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName)})
		}

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisaudio).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisaudio.FileExtension = filepath.Ext(FileName)
				thisaudio.CMS_ContentID = thiscontent.ID
				db.DB.Create(&thisaudio)
			}
		} else {
			db.DB.Model(thisaudio).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName)})
		}
	case "audio":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisaudio).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisaudio.FileExtension = filepath.Ext(FileName)
				thisaudio.CMS_ContentID = thiscontent.ID
				db.DB.Create(&thisaudio)
			}
		} else {
			db.DB.Model(thisaudio).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName)})
		}
	case "image":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisimage).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisimage.FileExtension = filepath.Ext(FileName)
				thisimage.CMS_ContentID = thiscontent.ID
				db.DB.Create(&thisimage)
			}
		} else {
			db.DB.Model(thisimage).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName)})
		}
	default:

	}

	db.DB.Model(thiscontent).Updates(map[string]interface{}{"State": 2, "file_name": FileName, "is_trans": 1})
	w.Write([]byte(("Update Complete")))
}

//func StartHawkeye
func StartHawkeye(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	ContentID := fmt.Sprintf("%s", queryValues.Get("ContentID"))

	var thiscontent content.CMS_Content

	db.DB.Where("content_id=?", ContentID).First(&thiscontent)

	db.DB.Model(thiscontent).Updates(map[string]interface{}{"State": 3})
	w.Write([]byte(("Update Complete")))
}

//func CompleteHawkeye
func CompleteHawkeye(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	ContentID := fmt.Sprintf("%s", queryValues.Get("ContentID"))
	ContentType := fmt.Sprintf("%s", queryValues.Get("ContentType"))
	FileName := fmt.Sprintf("%s", queryValues.Get("FileName"))

	//Video info
	VideoWidth := fmt.Sprintf("%s", queryValues.Get("VideoWidth"))
	VideoHeight := fmt.Sprintf("%s", queryValues.Get("VideoHeight"))
	VideoFramerate := fmt.Sprintf("%s", queryValues.Get("VideoFramerate"))
	VideoCodec := fmt.Sprintf("%s", queryValues.Get("VideoCodec"))
	VideoFormat := fmt.Sprintf("%s", queryValues.Get("VideoFormat"))
	VideoScantype := fmt.Sprintf("%s", queryValues.Get("VideoScantype"))
	VideoDuration := fmt.Sprintf("%s", queryValues.Get("VideoDuration"))
	VideoDisplayaspectratio := fmt.Sprintf("%s", queryValues.Get("VideoDisplayaspectratio"))
	VideoColorspace := fmt.Sprintf("%s", queryValues.Get("VideoColorspace"))
	VideoBitdepth := fmt.Sprintf("%s", queryValues.Get("videoBitdepth"))

	//Audio info
	AudioSamplingrate := fmt.Sprintf("%s", queryValues.Get("AudioSamplingrate"))
	AudioChannel := fmt.Sprintf("%s", queryValues.Get("AudioChannel"))
	AudioBitrate := fmt.Sprintf("%s", queryValues.Get("AudioBitrate"))
	AudioBitdepth := fmt.Sprintf("%s", queryValues.Get("AudioBitdepth"))
	AudioDuration := fmt.Sprintf("%s", queryValues.Get("AudioDuration"))
	AudioFormat := fmt.Sprintf("%s", queryValues.Get("AudioFormat"))
	AudioCodec := fmt.Sprintf("%s", queryValues.Get("AudioCodec"))

	var thisvideo content.CMS_Video
	var thisaudio content.CMS_Audio
	var thisimage content.CMS_Image

	var thiscontent content.CMS_Content

	db.DB.Where("content_id=?", ContentID).First(&thiscontent)

	switch ContentType {
	case "video":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisvideo).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisvideo.FileExtension = filepath.Ext(FileName)
				thisvideo.CMS_ContentID = thiscontent.ID

				thisvideo.Codec = VideoCodec
				thisvideo.Format = VideoFormat
				thisvideo.Displayaspectratio = VideoDisplayaspectratio
				thisvideo.Duration = VideoDuration
				thisvideo.Framerate = VideoFramerate
				thisvideo.Scantype = VideoScantype
				thisvideo.Colorspace = VideoCodec
				thisvideo.Bitdepth = VideoBitdepth

				db.DB.Create(&thisvideo)
			}
		} else {
			db.DB.Model(thisvideo).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName), "width": VideoWidth, "height": VideoHeight, "codec": VideoCodec, "format": VideoFormat, "scantype": VideoScantype, "displayaspectratio": VideoDisplayaspectratio, "colorspace": VideoColorspace, "bitdepth": VideoBitdepth, "framerate": VideoFramerate})
		}

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisaudio).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisaudio.FileExtension = filepath.Ext(FileName)
				thisaudio.CMS_ContentID = thiscontent.ID

				thisaudio.Channels = AudioChannel
				thisaudio.Samplingrate = AudioSamplingrate
				thisaudio.Bitrate = AudioBitrate
				thisaudio.Bitdepth = AudioBitdepth
				thisaudio.Duration = AudioDuration
				thisaudio.Codec = AudioCodec

				db.DB.Create(&thisaudio)
			}
		} else {
			db.DB.Model(thisaudio).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName), "channels": AudioChannel, "samplingrate": AudioSamplingrate, "bitrate": AudioBitrate, "bitdepth": AudioBitdepth, "duration": AudioDuration, "format": AudioFormat, "codec": AudioCodec})
		}
	case "audio":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisaudio).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisaudio.FileExtension = filepath.Ext(FileName)
				thisaudio.CMS_ContentID = thiscontent.ID

				thisaudio.Channels = AudioChannel
				thisaudio.Samplingrate = AudioSamplingrate
				thisaudio.Bitrate = AudioBitrate
				thisaudio.Bitdepth = AudioBitdepth
				thisaudio.Duration = AudioDuration
				thisaudio.Codec = AudioCodec

				db.DB.Create(&thisaudio)
			}
		} else {
			db.DB.Model(thisaudio).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName), "channels": AudioChannel, "samplingrate": AudioSamplingrate, "bitrate": AudioBitrate, "bitdepth": AudioBitdepth, "duration": AudioDuration, "format": AudioFormat, "codec": AudioCodec})
		}

	case "image":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisimage).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisimage.FileExtension = filepath.Ext(FileName)
				thisimage.CMS_ContentID = thiscontent.ID
				db.DB.Create(&thisimage)
			}
		} else {
			db.DB.Model(thisimage).Updates(map[string]interface{}{"file_extension": filepath.Ext(FileName)})
		}
	default:

	}

	db.DB.Model(thiscontent).Updates(map[string]interface{}{"State": 2, "file_name": FileName, "is_hawkeye": 1})
	w.Write([]byte(("Update Complete")))
}

func parseInt(s string, dest *int) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*dest = n
	return nil
}

//func UpdateMediainfo
func UpdateMediainfo(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	ContentID := fmt.Sprintf("%s", queryValues.Get("ContentID"))
	ContentType := fmt.Sprintf("%s", queryValues.Get("ContentType"))

	//Video info\

	strVideoWidth := fmt.Sprintf("%s", queryValues.Get("VideoWidth"))

	strVideoHeight := fmt.Sprintf("%s", queryValues.Get("VideoHeight"))

	var VideoWidth, VideoHeight int
	if err := parseInt(strVideoWidth, &VideoWidth); err != nil {
		log.Printf("Error parsing VideoWidth: %s", err)
	}

	if err := parseInt(strVideoHeight, &VideoHeight); err != nil {
		log.Printf("Error parsing VideoHeight: %s", err)
	}

	VideoFramerate := fmt.Sprintf("%s", queryValues.Get("VideoFramerate"))
	VideoCodec := fmt.Sprintf("%s", queryValues.Get("VideoCodec"))
	VideoFormat := fmt.Sprintf("%s", queryValues.Get("VideoFormat"))
	VideoScantype := fmt.Sprintf("%s", queryValues.Get("VideoScantype"))
	VideoDuration := fmt.Sprintf("%s", queryValues.Get("VideoDuration"))
	VideoDisplayaspectratio := fmt.Sprintf("%s", queryValues.Get("VideoDisplayaspectratio"))
	VideoColorspace := fmt.Sprintf("%s", queryValues.Get("VideoColorspace"))
	VideoBitdepth := fmt.Sprintf("%s", queryValues.Get("videoBitdepth"))

	//Audio info
	AudioSamplingrate := fmt.Sprintf("%s", queryValues.Get("AudioSamplingrate"))
	AudioChannel := fmt.Sprintf("%s", queryValues.Get("AudioChannel"))
	AudioBitrate := fmt.Sprintf("%s", queryValues.Get("AudioBitrate"))
	AudioBitdepth := fmt.Sprintf("%s", queryValues.Get("AudioBitdepth"))
	AudioDuration := fmt.Sprintf("%s", queryValues.Get("AudioDuration"))
	AudioFormat := fmt.Sprintf("%s", queryValues.Get("AudioFormat"))
	AudioCodec := fmt.Sprintf("%s", queryValues.Get("AudioCodec"))

	var thisvideo content.CMS_Video
	var thisaudio content.CMS_Audio
	var thisimage content.CMS_Image

	var thiscontent content.CMS_Content

	db.DB.Where("content_id=?", ContentID).First(&thiscontent)

	switch ContentType {
	case "video":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisvideo).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisvideo.CMS_ContentID = thiscontent.ID
				thisvideo.Codec = VideoCodec
				thisvideo.Format = VideoFormat
				thisvideo.Displayaspectratio = VideoDisplayaspectratio
				thisvideo.Duration = VideoDuration
				thisvideo.Framerate = VideoFramerate
				thisvideo.Scantype = VideoScantype
				thisvideo.Colorspace = VideoColorspace
				thisvideo.Bitdepth = VideoBitdepth
				db.DB.Create(&thisvideo)
			}
		} else {
			if VideoWidth != 0 {
				thisvideo.Width = VideoWidth
				db.DB.Save(&thisvideo)
			}
			if VideoHeight != 0 {
				thisvideo.Height = VideoHeight
				db.DB.Save(&thisvideo)
			}
			if VideoCodec != "" {
				thisvideo.Codec = VideoCodec
				db.DB.Save(&thisvideo)
			}
			if VideoFormat != "" {
				thisvideo.Format = VideoFormat
				db.DB.Save(&thisvideo)
			}
			if VideoScantype != "" {
				thisvideo.Scantype = VideoScantype
				db.DB.Save(&thisvideo)
			}
			if VideoDisplayaspectratio != "" {
				thisvideo.Displayaspectratio = VideoDisplayaspectratio
				db.DB.Save(&thisvideo)
			}
			if VideoDuration != "" {
				thisvideo.Duration = VideoDuration
				db.DB.Save(&thisvideo)
			}
			if VideoColorspace != "" {
				thisvideo.Colorspace = VideoColorspace
				db.DB.Save(&thisvideo)
			}
			if VideoBitdepth != "" {
				thisvideo.Bitdepth = VideoBitdepth
				db.DB.Save(&thisvideo)
			}
			if VideoFramerate != "" {
				thisvideo.Framerate = VideoFramerate
				db.DB.Save(&thisvideo)
			}
		}

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisaudio).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisaudio.CMS_ContentID = thiscontent.ID
				thisaudio.Channels = AudioChannel
				thisaudio.Samplingrate = AudioSamplingrate
				thisaudio.Bitrate = AudioBitrate
				thisaudio.Bitdepth = AudioBitdepth
				thisaudio.Duration = AudioDuration
				thisaudio.Codec = AudioCodec
				db.DB.Create(&thisaudio)
			}
		} else {

			if AudioChannel != "" {
				thisaudio.Channels = AudioChannel
				db.DB.Save(&thisaudio)
			}
			if AudioSamplingrate != "" {
				thisaudio.Samplingrate = AudioSamplingrate
				db.DB.Save(&thisaudio)
			}
			if AudioBitrate != "" {
				thisaudio.Bitrate = AudioBitrate
				db.DB.Save(&thisaudio)
			}
			if AudioBitdepth != "" {
				thisaudio.Bitdepth = AudioBitdepth
				db.DB.Save(&thisaudio)
			}
			if AudioDuration != "" {
				thisaudio.Duration = AudioDuration
				db.DB.Save(&thisaudio)
			}
			if AudioFormat != "" {
				thisaudio.Format = AudioFormat
				db.DB.Save(&thisaudio)
			}
			if AudioCodec != "" {
				thisaudio.Codec = AudioCodec
				db.DB.Save(&thisaudio)
			}

		}

	case "audio":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisaudio).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisaudio.CMS_ContentID = thiscontent.ID

				thisaudio.Channels = AudioChannel
				thisaudio.Samplingrate = AudioSamplingrate
				thisaudio.Bitrate = AudioBitrate
				thisaudio.Bitdepth = AudioBitdepth
				thisaudio.Duration = AudioDuration
				thisaudio.Codec = AudioCodec

				db.DB.Create(&thisaudio)
			}
		} else {
			if AudioChannel != "" {
				thisaudio.Channels = AudioChannel
				db.DB.Save(&thisaudio)
			}
			if AudioSamplingrate != "" {
				thisaudio.Samplingrate = AudioSamplingrate
				db.DB.Save(&thisaudio)
			}
			if AudioBitrate != "" {
				thisaudio.Bitrate = AudioBitrate
				db.DB.Save(&thisaudio)
			}
			if AudioBitdepth != "" {
				thisaudio.Bitdepth = AudioBitdepth
				db.DB.Save(&thisaudio)
			}
			if AudioDuration != "" {
				thisaudio.Duration = AudioDuration
				db.DB.Save(&thisaudio)
			}
			if AudioFormat != "" {
				thisaudio.Format = AudioFormat
				db.DB.Save(&thisaudio)
			}
			if AudioCodec != "" {
				thisaudio.Codec = AudioCodec
				db.DB.Save(&thisaudio)
			}
		}

	case "image":

		if err := db.DB.Where("cms_content_id=?", thiscontent.ID).First(&thisimage).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				thisimage.CMS_ContentID = thiscontent.ID
				db.DB.Create(&thisimage)
			}
		} else {
			//db.DB.Model(&thisimage).Updates(map[string]interface{}{})
		}
	default:

	}

	w.Write([]byte(("Update Complete")))
}
