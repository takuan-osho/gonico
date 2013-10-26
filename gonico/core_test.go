package gonico

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVideoThumbResponse(t *testing.T) {
	result, status := GetVideoThumbResponse("sm1")
	switch status {
	case "fail":
		errorInfo := result.ErrorInfo
		assert.Equal(t, errorInfo.Code, "DELETED")
		assert.Equal(t, errorInfo.Description, "deleted")
	default:
		t.Fail()
	}

	result, status = GetVideoThumbResponse("sm9")
	switch status {
	case "ok":
		videoInfo := result.VideoInfo
		assert.Equal(t, videoInfo.VideoId, "sm9")
		assert.Equal(t, videoInfo.MovieType, "flv")
	default:
		t.Fail()
	}
}
