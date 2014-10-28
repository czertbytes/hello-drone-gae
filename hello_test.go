package hello

import (
	"appengine/aetest"
	"testing"
)

func TestDummy(t *testing.T) {
	aeContext, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	if "a" != "a" {
		t.Fatal("something is wrong")
	}
	aeContext.Infof("Hello test!")
}
