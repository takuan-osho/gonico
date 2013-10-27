package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gonico core test", func() {

	Describe("test GetVideoThumbResponse function", func() {
		var (
			resp NicoVideoThumbResponse
			err  error
		)

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

	Describe("test GetVideoTitle function", func() {
		var title string

		Context("when the video is deleted", func() {
			BeforeEach(func() {
				title = GetVideoTitle("sm1")
			})

			It("should return empty value", func() {
				Expect(title).To(BeEmpty())
			})
		})

		Context("when the video is alive", func() {
			BeforeEach(func() {
				title = GetVideoTitle("sm9")
			})

			It("should return correct value", func() {
				Expect(title).To(ContainSubstring("陰陽師"))
			})
		})
	})

	Describe("test GetAllKindsOfVideoTags function", func() {
		var allTags, lockedTags, unlockedTags []string

		Context("when the video is deleted", func() {
			BeforeEach(func() {
				allTags, lockedTags, unlockedTags = GetAllKindsOfVideoTags("sm1")
			})

			It("should return empty value", func() {
				Expect(allTags).To(BeEmpty())
				Expect(lockedTags).To(BeEmpty())
				Expect(unlockedTags).To(BeEmpty())
			})
		})

		Context("when the video is alive", func() {
			BeforeEach(func() {
				allTags, lockedTags, unlockedTags = GetAllKindsOfVideoTags("sm9")
			})

			It("should return correct kinds of video tags", func() {
				Expect(allTags[0]).To(Equal("陰陽師"))
				Expect(lockedTags[0]).To(Equal("陰陽師"))
			})
		})
	})
})
