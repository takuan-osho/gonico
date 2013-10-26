package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	resp, err := Login()
	authFlag := resp.Header["X-Niconico-Authflag"][0]
	if err != nil {
		assert.Equal(t, authFlag, 0)
	} else {
		assert.NotEqual(t, authFlag, 0)
	}
	fmt.Println(resp.Header)
}
