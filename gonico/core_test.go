package gonico

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVideoThumbResponse(t *testing.T) {
	result, err := GetVideoThumbResponse("sm1")
	if err != nil {
		t.Fail()
	}
	errorInfo := result.ErrorInfo
	assert.Equal(t, errorInfo.Code, "DELETED")
	assert.Equal(t, errorInfo.Description, "deleted")

	result, err = GetVideoThumbResponse("sm9")
	if err != nil {
		t.Fail()
	}
	videoInfo := result.VideoInfo
	assert.Equal(t, videoInfo.VideoId, "sm9")
	assert.Equal(t, videoInfo.MovieType, "flv")
}
