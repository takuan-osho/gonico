package main

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

func TestGetAllKindsOfVideoTags(t *testing.T) {
	allTags, lockedTags, unlockedTags := GetAllKindsOfVideoTags("sm1")
	assert.Nil(t, allTags)
	assert.Nil(t, lockedTags)
	assert.Nil(t, unlockedTags)

	allTags, lockedTags, unlockedTags = GetAllKindsOfVideoTags("sm9")
	assert.Equal(t, allTags[0], "陰陽師")
	assert.Equal(t, lockedTags[0], "陰陽師")
}
