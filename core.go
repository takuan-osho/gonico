package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

var NicoAPIUrls = map[string]string{
	"login":        "https://secure.nicovideo.jp/secure/login?site=niconico",
	"getthumbinfo": "http://ext.nicovideo.jp/api/getthumbinfo/",
	"getflv":       "http://flapi.nicovideo.jp/api/getflv/",
}

type NicoVideoErrorInfo struct {
	XMLName     xml.Name `xml:"error"`
	Code        string   `xml:"code"`
	Description string   `xml:"description"`
}

type NicoVideoTag struct {
	XMLName xml.Name `xml:"tag"`
	Lock    string   `xml:"lock,attr"`
	Value   string   `xml:",innerxml"`
}

type NicoVideoTags struct {
	XMLName xml.Name       `xml:"tags"`
	Domain  string         `xml:"domain,attr"`
	Tag     []NicoVideoTag `xml:"tag"`
}

type NicoVideoInfo struct {
	XMLName       xml.Name        `xml:"thumb"`
	VideoId       string          `xml:"video_id"`
	Title         string          `xml:"title"`
	Description   string          `xml:"description"`
	ThumbnailUrl  string          `xml:"thumbnail_url"`
	FirstRetrieve string          `xml:"first_retrieve"`
	Length        string          `xml:"length"`
	MovieType     string          `xml:"movie_type"`
	SizeHigh      int             `xml:"size_high"`
	SizeLow       int             `xml:"size_low"`
	ViewCounter   int             `xml:"view_counter"`
	CommentNum    int             `xml:"comment_num"`
	MylistCounter int             `xml:"mylist_counter"`
	LastResBody   string          `xml:"last_res_body"`
	WatchUrl      string          `xml:"watch_url"`
	ThumbType     string          `xml:"thumb_type"`
	Embeddable    int             `xml:"embeddable"`
	NoLivePlay    int             `xml:"no_live_play"`
	Tags          []NicoVideoTags `xml:"tags"`
	UserId        int             `xml:"user_id"`
}

type NicoVideoThumbResponse struct {
	XMLName   xml.Name           `xml:"nicovideo_thumb_response"`
	Status    string             `xml:"status,attr"`
	ErrorInfo NicoVideoErrorInfo `xml:"error"`
	VideoInfo NicoVideoInfo      `xml:"thumb"`
}

func GetVideoThumbResponse(videoId string) (result NicoVideoThumbResponse, err error) {
	resp, _ := http.Get(NicoAPIUrls["getthumbinfo"] + videoId)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(body, &result)
	return
}

func GetVideoTitle(videoId string) (title string) {
	resp, err := GetVideoThumbResponse(videoId)

	if err != nil {
		return
	}

	if resp.Status != "ok" {
		return
	}

	title = resp.VideoInfo.Title
	return
}

func GetAllKindsOfVideoTags(videoId string) (allTags []string, lockedTags []string, unlockedTags []string) {
	resp, err := GetVideoThumbResponse(videoId)

	if err != nil {
		return
	}

	if resp.Status != "ok" {
		return
	}

	videoInfo := resp.VideoInfo
	for _, tag := range videoInfo.Tags[0].Tag {
		allTags = append(allTags, tag.Value)

		switch tag.Lock {
		case "1":
			lockedTags = append(lockedTags, tag.Value)
		default:
			unlockedTags = append(unlockedTags, tag.Value)
		}
	}
	return
}

func GetVideoTags(videoId string) (tags []string) {
	tags, _, _ = GetAllKindsOfVideoTags(videoId)
	return
}

func GetLockedVideoTags(videoId string) (lockedTags []string) {
	_, lockedTags, _ = GetAllKindsOfVideoTags(videoId)
	return
}

func GetUnlockedVideoTags(videoId string) (unlockedTags []string) {
	_, _, unlockedTags = GetAllKindsOfVideoTags(videoId)
	return
}
