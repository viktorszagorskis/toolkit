package toolkit

import "testing"

func TestToolsRandomString(t *testing.T) {
	var testTools Tools

	s:= testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}
}