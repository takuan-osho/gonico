package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("gonico user test", func() {

	Describe("test Login function", func() {
		var (
			resp     *http.Response
			err      error
			authFlag string
		)

		BeforeEach(func() {
			resp, err = Login()
			authFlag = resp.Header["X-Niconico-Authflag"][0]
		})

		Context("when login error occurred", func() {
			if err != nil {
				It("should return 0", func() {
					Expect(authFlag).To(Equal("0"))
				})
			}
		})

		Context("when login success occurred", func() {
			if err == nil {
				It("should not return 0", func() {
					Expect(authFlag).NotTo(Equal("0"))
				})
			}
		})

	})
})
