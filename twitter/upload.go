package twitter

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"time"
)

var chunkSize int = 5 * 1024 * 1024

func (u *User) UploadMedia(text, filepath, mediaType, mediaCategory string) error {
	fi, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fi.Close()

	size, err := GetFileSize(fi)
	if err != nil {
		return err
	}
	byteMedia, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}

	chunkedMedia, err := u.API.UploadVideoInit(size, mediaType, mediaCategory)
	if err != nil {
		return err
	}
	index := 0
	for i := 0; i < size; i += chunkSize {
		var data string
		if (size-i)/chunkSize > 0 {
			data = base64.StdEncoding.EncodeToString(byteMedia[i : i+chunkSize])
		} else {
			data = base64.StdEncoding.EncodeToString(byteMedia[i:])
		}
		if err = u.API.UploadVideoAppend(chunkedMedia.MediaIDString, index, data); err != nil {
			return err
		}
		index++
	}
	video, err := u.API.UploadVideoFinalize(chunkedMedia.MediaIDString)
	if err != nil {
		return err
	}
	//status
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

	params := url.Values{}
	params.Add("media_ids", video.MediaIDString)
	if _, err := u.API.PostTweet(text, params); err != nil {
		return err
	}
	return nil
}

func (u *User) ImageTweet(text, filepath string) error {
	fi, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fi.Close()
	byteMedia, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}
	base64Image := base64.StdEncoding.EncodeToString(byteMedia)
	media, err := u.API.UploadMedia(base64Image)
	if err != nil {
		return err
	}
	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)
	if _, err = u.API.PostTweet(text, v); err != nil {
		return err
	}
	fmt.Printf("tweet :%s\n", text)
	return nil
}
