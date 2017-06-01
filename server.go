package iso8583

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// WebServer returns a http web server that interprets ISO8583 messages
func WebServer(listenAddr string) (*http.Server, error) {
	if _, err := os.Stat(filepath.Join("web", "index.html")); err != nil {
		return nil, fmt.Errorf("missing web/ folder with index.html in launch directory")
	}
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)
	svr := &http.Server{
		Addr:           listenAddr,
		ReadTimeout:    time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return svr, nil
}
