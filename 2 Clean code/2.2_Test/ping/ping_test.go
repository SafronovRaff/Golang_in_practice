package ping

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	client := &http.Client{}
	pinger := Pinger{client}
	got := pinger.Ping("https://www.google.com")
	if !got {
		t.Errorf(" https://www.google.com/404")
	}
	got = pinger.Ping("https://www.google.com/404")
	if got {
		t.Errorf("Expected google.com/404 to be unvailable")
	}
}
