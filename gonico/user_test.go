package gonico

import (
	"testing"
)

func TestLogin(t *testing.T) {
	if Login() != true {
		t.Errorf("error!")
	}
}
