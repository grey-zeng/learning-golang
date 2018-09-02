package task

import "testing"

func TestTest(t *testing.T) {
	if "1" != "1" {
		t.Error("Error")
	}
}
