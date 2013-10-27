package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

var _ = Describe("gonico core test", func() {

	var (
		resp NicoVideoThumbResponse
		err  error
	)

	Describe("test GetVideoThumbResponse function", func() {

		Context("when the video is deleted", func() {
			BeforeEach(func() {
				resp, err = GetVideoThumbResponse("sm1")
				if err != nil {
					Fail("Network connection error!!!")
				}
			})

			It("should return ErrorInfo which tells the video is deleted", func() {
				errorInfo := resp.ErrorInfo
				Expect(errorInfo.Code).To(Equal("DELETED"))
				Expect(errorInfo.Description).To(Equal("deleted"))
			})
		})

		Context("when the video is alive", func() {
			BeforeEach(func() {
				resp, err = GetVideoThumbResponse("sm9")
			})

			It("should return VideoInfo", func() {
				videoInfo := resp.VideoInfo
				Expect(videoInfo.VideoId).To(Equal("sm9"))
				Expect(videoInfo.MovieType).To(Equal("flv"))
			})
		})
	})
})
