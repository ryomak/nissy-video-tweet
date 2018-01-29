#nissy-tweet
- 動画(86000秒)
- 画像(7200秒)
ごとにツイート

### anaconda 
140秒動画に対応してない為，
自作で追加
- Init
- Append
- Finalize
- Status(NEW!)

#### UploadVideoInit()
```Go:UploadVideoInit.go
v.Set("media_category", mediaCategory)
```

#### UploadVideoStatus()
```Go:UploadVideoStatus.go
type StatusMedia struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	MediaKey         string `json:"media_key"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	ProcessingInfo   struct {
		State           string `json:"state"`
		CheckAfterSecs  int    `json:"check_after_secs"`
		ProgressPercent int    `json:"progress_percent"`
	} `json:"processing_info"`
}

func (a TwitterApi) UploadVideoStatus(mediaIdString string) (videoMedia StatusMedia, err error) {
	v := url.Values{}
	v.Set("command", "STATUS")
	v.Set("media_id", mediaIdString)

	var mediaResponse StatusMedia

	response_ch := make(chan response)
	a.queryQueue <- query{UploadBaseUrl + "/media/upload.json", v, &mediaResponse, _GET, response_ch}
	return mediaResponse, (<-response_ch).err
}
```

#### upload
```Go
//tweet する際,status部分
for {
		videos, err := u.API.UploadVideoStatus(chunkedMedia.MediaIDString)
		if err != nil {
			return err
		}
		if videos.ProcessingInfo.State == "succeeded" {
			break
		}
		time.Sleep(time.Duration(videos.ProcessingInfo.CheckAfterSecs))
	}
```
