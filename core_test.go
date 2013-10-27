package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVideoThumbResponse(t *testing.T) {
	resp, err := GetVideoThumbResponse("sm1")
	if err != nil {
		t.Fail()
	}
	errorInfo := resp.ErrorInfo
	assert.Equal(t, errorInfo.Code, "DELETED")
	assert.Equal(t, errorInfo.Description, "deleted")

	resp, err = GetVideoThumbResponse("sm9")
	if err != nil {
		t.Fail()
	}
	videoInfo := resp.VideoInfo
	assert.Equal(t, videoInfo.VideoId, "sm9")
	assert.Equal(t, videoInfo.MovieType, "flv")
}

func TestGetVideoTitle(t *testing.T) {
	title := GetVideoTitle("sm1")
	assert.Empty(t, title)

	title = GetVideoTitle("sm9")
	assert.Contains(t, title, "陰陽師")
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

var _ = Describe("gonico core", func() {
	var (
		resp NicoVideoThumbResponse
		err  error
	)

	BeforeEach(func() {
		resp, err = GetVideoThumbResponse("sm1")
	})

	It("error loading", func() {
		errorInfo := resp.ErrorInfo
		Expect(errorInfo.Code).To(Equal("DELETED"))
	})
})
