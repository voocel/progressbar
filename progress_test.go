package progressbar

import (
	"io"
	"net/http"
	"testing"
)

var url = "https://github.com/voocel/sv/releases/download/v1.1.2/sv-linux-amd-64"

func TestBar(t *testing.T) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bar := NewBar(resp.ContentLength)
	bar.SetText("test", "red")
	buf := make([]byte, 1024)
	io.CopyBuffer(bar, resp.Body, buf)
}
