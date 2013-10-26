package gonico

import (
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, videoInfo.VideoId, videoId)
	default:
		t.Fail()
	}
}
