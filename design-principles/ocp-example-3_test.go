package design_principles

import "testing"

func TestApplicationContext(t *testing.T) {
	apiStatInfo := NewDefaultApiStatInfo()
	applicationContextInstance := GetApplicationContextInstance()
	applicationContextInstance.GetAlert().Check(apiStatInfo)
}
