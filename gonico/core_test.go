package gonico

import (
	"testing"
)

const (
	videoId = "sm9"
)

func TestGetVideoThumbResponse(t *testing.T) {
	result, status := GetVideoThumbResponse(videoId)
	switch status {
	case "ok":
		videoInfo := result.VideoInfo
		if result.VideoInfo.VideoId != videoId {
			t.Errorf("Expected %v, got %v.", videoId, videoInfo.VideoId)
		}
	default:
		t.Fail()
	}
}
